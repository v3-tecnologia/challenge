package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/yanvic/challenge/core/entity"
	"github.com/yanvic/challenge/internal/rekognition"
)

func StartImageConsumer() {
	_, err := conn.Subscribe("telemetry.image", func(m *nats.Msg) {
		log.Println("Mensagem recebida na fila telemetry.image")

		var photo entity.Photo
		if err := json.Unmarshal(m.Data, &photo); err != nil {
			log.Printf("Erro ao fazer unmarshal da imagem: %v", err)
			return
		}

		client, err := rekognition.NewClient(context.TODO())
		if err != nil {
			log.Printf("Erro ao criar cliente Rekognition: %v", err)
			return
		}

		result, err := client.DetectFaces(photo.Image)
		if err != nil {
			log.Printf("Erro ao detectar faces: %v", err)
			return
		}

		log.Printf("Detectadas %d faces na imagem do device %s", len(result.FaceDetails), photo.DeviceID)
	})
	if err != nil {
		log.Fatalf("Erro ao subscrever no NATS: %v", err)
	}
}
