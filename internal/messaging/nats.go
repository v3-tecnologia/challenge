package messaging

import (
	"encoding/json"
	"log"
	"time"

	"go-challenge/internal/services"

	"github.com/nats-io/nats.go"
)

type NATSProducer struct {
	conn *nats.Conn
}

func NewNATSProducer(url string) (*NATSProducer, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NATSProducer{conn: nc}, nil
}

func (p *NATSProducer) Publish(subject string, msg interface{}) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return p.conn.Publish(subject, data)
}

func (p *NATSProducer) Subscribe(subject string, handler func([]byte)) (*nats.Subscription, error) {
	return p.conn.Subscribe(subject, func(m *nats.Msg) {
		handler(m.Data)
	})
}

// Retry and DLQ configuration
const (
	MaxRetries = 3
	DLQSubject = "photo.telemetry.dlq"
)

// Consumer with retry and DLQ
func (p *NATSProducer) SubscribeWithRetry(subject string, handler func([]byte) error) (*nats.Subscription, error) {
	return p.conn.Subscribe(subject, func(m *nats.Msg) {
		var attempt int
		var err error
		for attempt = 0; attempt < MaxRetries; attempt++ {
			err = handler(m.Data)
			if err == nil {
				log.Printf("[NATS] Message processed successfully on attempt %d", attempt+1)
				return
			}
			log.Printf("[NATS] Handler error (attempt %d): %v", attempt+1, err)
			time.Sleep(500 * time.Millisecond)
		}
		// Send to DLQ after retries
		log.Printf("[NATS] Sending message to DLQ after %d attempts", MaxRetries)
		_ = p.conn.Publish(DLQSubject, m.Data)
	})
}

// Example: publish to different telemetry topics
func (p *NATSProducer) PublishTelemetry(telemetryType string, msg interface{}) error {
	subject := telemetryType + ".telemetry"
	return p.Publish(subject, msg)
}

// Exemplo de integração completa: Producer → NATS → Consumer → Rekognition → Resposta/DB/DLQ
// Supondo que PhotoMessage já exista em outro local do projeto
// type PhotoMessage struct { ... }

type PhotoResultMessage struct {
	ID         string    `json:"id"`
	Recognized bool      `json:"recognized"`
	Details    string    `json:"details"`
	Timestamp  time.Time `json:"timestamp"`
}

// Interface esperada do serviço de reconhecimento
// (import "internal/services" e use services.PhotoRecognitionService se necessário)
// type PhotoRecognitionService interface { AnalyzePhoto([]byte) (PhotoRecognitionResult, error) }

// Producer: publica foto recebida
func (p *NATSProducer) PublishPhoto(photoMsg PhotoMessage) error {
	return p.PublishTelemetry("photo", photoMsg)
}

// Consumer: consome, processa com Rekognition e publica resultado
func (p *NATSProducer) StartPhotoConsumer(photoService interface {
	AnalyzePhoto([]byte) (services.PhotoRecognitionResult, error)
}) error {
	_, err := p.SubscribeWithRetry("photo.telemetry", func(data []byte) error {
		var msg PhotoMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			return err
		}
		result, err := photoService.AnalyzePhoto(msg.Photo)
		if err != nil {
			return err
		}
		resultMsg := PhotoResultMessage{
			ID:         msg.DeviceID, // Usando DeviceID como identificador
			Recognized: result.Recognized,
			Details:    result.Details,
			Timestamp:  result.Timestamp,
		}
		err = p.PublishTelemetry("photo_result", resultMsg)
		if err != nil {
			log.Printf("[NATS] Falha ao publicar resultado: %v", err)
			p.PublishTelemetry("photo_result.dlq", resultMsg) // Publica na DLQ em caso de falha
			return err
		}
		log.Printf("[RESULT] Foto %s reconhecida: %v", msg.DeviceID, result.Recognized)
		return nil // sucesso
	})
	return err
}
