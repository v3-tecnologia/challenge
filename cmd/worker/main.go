package main

import (
	"challenge-v3/ierr"
	"challenge-v3/messaging"
	"challenge-v3/metrics"
	"challenge-v3/models"
	"challenge-v3/services"
	"challenge-v3/storage"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log/slog"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

type Worker struct {
	db            storage.Storage
	photoAnalyzer services.PhotoAnalyzer
}

func (w *Worker) handleGyroscopeMsg(msg *nats.Msg) {
	subject := "telemetry.gyroscope"
	var data models.GyroscopeData
	if err := json.Unmarshal(msg.Data, &data); err != nil {
		slog.Error("Falha ao decodificar mensagem de giroscópios", "error", err.Error(), "action", "retrying_message")

		msg.Term()
		metrics.NatsMessagesProcessed.WithLabelValues(subject, "terminated").Inc()
		return
	}
	if err := w.db.SaveGyroscope(&data); err != nil {
		slog.Error("falha ao salvar dados de giroscópio", "error", err, "device_id", data.DeviceID)
		msg.Nak()
		metrics.NatsMessagesProcessed.WithLabelValues(subject, "failed").Inc()
		return
	}
	slog.Info("mensagem de giroscópio processada", "device_id", data.DeviceID)

	auditEvent := models.AuditEvent{
		Actor:  data.DeviceID,
		Action: "GYROSCOPE_PROCESSED",
		Details: map[string]interface{}{
			"x": *data.X,
			"y": *data.Y,
			"z": *data.Z,
		},
	}
	if err := w.db.LogAuditEvent(auditEvent); err != nil {
		slog.Error("falha ao registrar evento de auditoria para giroscópio", "error", err, "device_id", data.DeviceID)
	}

	msg.Ack()
	metrics.NatsMessagesProcessed.WithLabelValues(subject, "success").Inc()
}
func (w *Worker) handleGpsMsg(msg *nats.Msg) {
	subject := "telemetry.gps"
	var data models.GPSData
	if err := json.Unmarshal(msg.Data, &data); err != nil {
		slog.Error("falha ao decodificar mensagem de gps", "error", err)
		msg.Term()
		metrics.NatsMessagesProcessed.WithLabelValues(subject, "terminated").Inc()
		return
	}
	if err := w.db.SaveGPS(&data); err != nil {
		slog.Error("falha ao salvar dados de gps", "error", err, "device_id", data.DeviceID)
		msg.Nak()
		metrics.NatsMessagesProcessed.WithLabelValues(subject, "failed").Inc()
		return
	}
	slog.Info("mensagem de gps processada", "device_id", data.DeviceID)
	auditEvent := models.AuditEvent{
		Actor:  data.DeviceID,
		Action: "GPS_DATA_PROCESSED",
		Details: map[string]interface{}{
			"latitude":  *data.Latitude,
			"longitude": *data.Longitude,
		},
	}
	if err := w.db.LogAuditEvent(auditEvent); err != nil {
		slog.Error("falha ao registrar evento de auditoria para gps", "error", err, "device_id", data.DeviceID)
	}
	msg.Ack()
	metrics.NatsMessagesProcessed.WithLabelValues(subject, "success").Inc()
}
func (w *Worker) handlePhotoMsg(msg *nats.Msg) {
	subject := "telemetry.photo"
	var data models.PhotoData
	if err := json.Unmarshal(msg.Data, &data); err != nil {
		slog.Error("falha ao decodificar mensagem de foto", "error", err)
		msg.Term()
		metrics.NatsMessagesProcessed.WithLabelValues(subject, "terminated").Inc()
		return
	}
	_, err := w.photoAnalyzer.AnalyzeAndSavePhoto(&data)
	if err != nil {
		var validationErr *ierr.ValidationError
		if errors.As(err, &validationErr) {
			slog.Warn("erro de validação ao processar foto, mensagem terminada", "error", err, "device_id", data.DeviceID)
			msg.Term()
			metrics.NatsMessagesProcessed.WithLabelValues(subject, "terminated").Inc()
		} else {
			slog.Error("falha ao processar foto, mensagem será reenviada", "error", err, "device_id", data.DeviceID)
			msg.Nak()
			metrics.NatsMessagesProcessed.WithLabelValues(subject, "failed").Inc()
		}
		return
	}
	slog.Info("mensagem de foto processada com sucesso", "device_id", data.DeviceID)
	auditEvent := models.AuditEvent{
		Actor:   data.DeviceID,
		Action:  "PHOTO_PROCESSED",
		Details: map[string]interface{}{"recognized": data.Recognized},
	}
	if err := w.db.LogAuditEvent(auditEvent); err != nil {
		slog.Error("falha ao registrar evento de auditoria no worker", "error", err)
	}
	msg.Ack()
	metrics.NatsMessagesProcessed.WithLabelValues(subject, "success").Inc()
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("iniciando o worker")
	godotenv.Load()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := storage.NewPostgresStorage(connStr)
	if err != nil {
		slog.Error("falha ao conectar ao banco de dados", "error", err)
		os.Exit(1)

	}

	if err := db.InitTables(); err != nil {
		slog.Error("Não foi possível inicializar as tabelas a partir do worker", "error", err)
		os.Exit(1)
	}

	awsRegion := os.Getenv("AWS_REGION")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		slog.Error("Falha ao carregar config da AWS", "error", err)
		os.Exit(1)
	}
	rekognitionClient := rekognition.NewFromConfig(cfg)
	collectionID := os.Getenv("REKOGNITION_COLLECTION_ID")
	_, err = rekognitionClient.CreateCollection(context.TODO(), &rekognition.CreateCollectionInput{CollectionId: &collectionID})
	var resourceExistsErr *types.ResourceAlreadyExistsException
	if err != nil && !errors.As(err, &resourceExistsErr) {
		slog.Error("Falha ao criar/verificar coleção no Rekognition", "error", err)

	}

	photoAnalyzer := services.NewPhotoAnalyzerService(rekognitionClient, collectionID, db)
	natsURL := os.Getenv("NATS_URL")
	nc, err := messaging.ConnectNATS(natsURL)
	if err != nil {
		slog.Error("Falha ao conectar ao NATS", "error", err)
		os.Exit(1)
	}
	defer nc.Close()
	js, err := messaging.SetupJetStream(nc)
	if err != nil {
		slog.Error("Falha ao configurar o JetStream", "error", err)
		os.Exit(1)
	}

	metricsRouter := http.NewServeMux()
	metricsRouter.Handle("/metrics", metrics.MetricsHandler())
	go func() {
		slog.Info("Servidor de Métricas do Worker iniciado", "porta", ":8082")
		if err := http.ListenAndServe(":8082", metricsRouter); err != nil {
			slog.Error("Servidor de Métricas do Worker falhou", "error", err)
		}
	}()

	worker := &Worker{
		db:            db,
		photoAnalyzer: photoAnalyzer,
	}

	ackWait := nats.AckWait(30 * time.Second)
	js.Subscribe("telemetry.gyroscope", worker.handleGyroscopeMsg, nats.Durable("GYROSCOPE_WORKER"))
	js.Subscribe("telemetry.gps", worker.handleGpsMsg, nats.Durable("GPS_WORKER"))
	js.Subscribe("telemetry.photo", worker.handlePhotoMsg, nats.Durable("PHOTO_WORKER"), ackWait)

	slog.Info("Worker está no ar, esperando por mensagens de telemetria...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	slog.Info("Desligando o Worker...")
}
