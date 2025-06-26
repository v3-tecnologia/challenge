package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"

	"telemetry-api/internal/config"
	"telemetry-api/internal/database"
	"telemetry-api/internal/dtos/requests"
	"telemetry-api/internal/services"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "error to sync logger: %v\n", err)
		}
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("cannot load config for worker", zap.Error(err))
	}

	db, err := database.Connect(cfg)
	if err != nil {
		logger.Fatal("worker cannot connect to db", zap.Error(err))
	}

	nc, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		logger.Fatal("worker cannot connect to nats", zap.Error(err))
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		logger.Fatal("worker cannot get jetstream context", zap.Error(err))
	}

	// Garante que os streams existam
	createStreams(js, logger)

	telemetryService := services.NewTelemetryService(db, logger)

	subscribe(js, "telemetry.gyroscope", handleGyroscope(telemetryService, logger))
	subscribe(js, "telemetry.gps", handleGPS(telemetryService, logger))
	subscribe(js, "telemetry.photo", handlePhoto(telemetryService, logger))

	logger.Info("Worker started, waiting for messages...")

	// Aguarda o sinal de interrupção para encerrar de formagraceful
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	logger.Info("Worker shutting down...")
	nc.Close()
}

func createStreams(js nats.JetStreamContext, logger *zap.Logger) {
	streamConfigs := []struct {
		name     string
		subjects []string
	}{
		{"TELEMETRY_GYROSCOPE", []string{"telemetry.gyroscope"}},
		{"TELEMETRY_GPS", []string{"telemetry.gps"}},
		{"TELEMETRY_PHOTO", []string{"telemetry.photo"}},
	}

	for _, config := range streamConfigs {
		_, err := js.StreamInfo(config.name)
		if err != nil {
			logger.Info("Creating stream", zap.String("stream", config.name))
			_, err = js.AddStream(&nats.StreamConfig{
				Name:     config.name,
				Subjects: config.subjects,
			})
			if err != nil {
				logger.Fatal("Could not create stream", zap.String("stream", config.name), zap.Error(err))
			}
		}
	}
}

func subscribe(js nats.JetStreamContext, subject string, handler nats.MsgHandler) {
	// Substitui pontos por underscores para um nome durável válido
	durableName := strings.ReplaceAll(subject, ".", "_") + "_worker"
	_, err := js.Subscribe(subject, handler, nats.Durable(durableName), nats.ManualAck())
	if err != nil {
		log.Fatalf("failed to subscribe to subject %s: %v", subject, err)
	}
}

func handleGyroscope(service *services.TelemetryService, logger *zap.Logger) nats.MsgHandler {
	return func(m *nats.Msg) {
		var req requests.CreateGyroscopeRequest
		if err := json.Unmarshal(m.Data, &req); err != nil {
			logger.Error("Failed to unmarshal gyroscope data", zap.Error(err))
			if err := m.Nak(); err != nil {
				logger.Error("Failed to NAK message", zap.Error(err))
			}
			return
		}
		if err := service.CreateGyroscope(&req); err != nil {
			logger.Error("Failed to process gyroscope data", zap.Error(err))
			if err := m.Nak(); err != nil {
				logger.Error("Failed to NAK message", zap.Error(err))
			}
			return
		}
		logger.Info("Processed gyroscope message", zap.String("device_id", req.DeviceID))
		if err := m.Ack(); err != nil {
			logger.Error("Failed to ACK message", zap.Error(err))
		}
	}
}

func handleGPS(service *services.TelemetryService, logger *zap.Logger) nats.MsgHandler {
	return func(m *nats.Msg) {
		var req requests.CreateGPSRequest
		if err := json.Unmarshal(m.Data, &req); err != nil {
			logger.Error("Failed to unmarshal GPS data", zap.Error(err))
			if err := m.Nak(); err != nil {
				logger.Error("Failed to NAK message", zap.Error(err))
			}
			return
		}
		if err := service.CreateGPS(&req); err != nil {
			logger.Error("Failed to process GPS data", zap.Error(err))
			if err := m.Nak(); err != nil {
				logger.Error("Failed to NAK message", zap.Error(err))
			}
			return
		}
		logger.Info("Processed GPS message", zap.String("device_id", req.DeviceID))
		if err := m.Ack(); err != nil {
			logger.Error("Failed to ACK message", zap.Error(err))
		}
	}
}

func handlePhoto(service *services.TelemetryService, logger *zap.Logger) nats.MsgHandler {
	return func(m *nats.Msg) {
		var req requests.CreateTelemetryPhotoRequest
		if err := json.Unmarshal(m.Data, &req); err != nil {
			logger.Error("Failed to unmarshal photo data", zap.Error(err))
			if err := m.Nak(); err != nil {
				logger.Error("Failed to NAK message", zap.Error(err))
			}
			return
		}
		if err := service.CreateTelemetryPhoto(&req); err != nil {
			logger.Error("Failed to process photo data", zap.Error(err))
			if err := m.Nak(); err != nil {
				logger.Error("Failed to NAK message", zap.Error(err))
			}
			return
		}
		logger.Info("Processed photo message", zap.String("device_id", req.DeviceID))
		if err := m.Ack(); err != nil {
			logger.Error("Failed to ACK message", zap.Error(err))
		}
	}
}
