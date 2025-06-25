package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/nats-io/nats.go"
	"github.com/yanvic/challenge/core/entity"
	"github.com/yanvic/challenge/infra/cache"
	"github.com/yanvic/challenge/infra/database/dynamo"
	rekognitionInfra "github.com/yanvic/challenge/infra/rekognition"
)

func StartImageConsumer() {
	_, err := conn.Subscribe("telemetry.image", func(m *nats.Msg) {
		log.Println("Mensagem recebida na fila telemetry.image")
		start := time.Now()

		var photo entity.Photo
		if err := json.Unmarshal(m.Data, &photo); err != nil {
			log.Printf("Erro ao fazer unmarshal da imagem: %v", err)
			return
		}

		if recognized, exists := cache.IsCached(photo.Image); exists {
			log.Printf("Imagem já processada. Reconhecida: %v", recognized)

			duration := time.Since(start).Milliseconds()
			savePhotoAnalysis(photo, entity.PhotoAnalysisResult{
				Recognized: recognized,
				Reason:     "Resultado obtido do cache",
			}, nil, "", 0, duration)

			log.Printf("Análise via cache concluída para device %s em %dms", photo.DeviceID, duration)
			return
		}

		client, err := rekognitionInfra.InitRekognitionClient(context.TODO())
		if err != nil {
			log.Printf("Erro ao criar cliente Rekognition: %v", err)
			return
		}

		// Detecta faces na foto atual
		result, err := detectFaces(client, photo.Image)
		if err != nil {
			log.Printf("Erro ao detectar faces: %v", err)
			return
		}

		log.Printf("Detectadas %d faces na imagem do device %s", len(result.FaceDetails), photo.DeviceID)

		isRecognized, matchedPhotoID, similarity := compareWithPreviousPhotos(client, photo, result.FaceDetails)

		analysisResult := entity.PhotoAnalysisResult{
			Recognized: isRecognized,
			Reason:     generateReason(isRecognized, len(result.FaceDetails), similarity),
		}

		duration := time.Since(start).Milliseconds()
		err = savePhotoAnalysis(photo, analysisResult, result, matchedPhotoID, similarity, duration)
		if err != nil {
			log.Printf("Erro ao salvar análise: %v", err)
			return
		}

		cache.SetCache(photo.Image, isRecognized)

		log.Printf("Análise concluída para device %s em %dms - Reconhecida: %v",
			photo.DeviceID, duration, isRecognized)
	})

	if err != nil {
		log.Fatalf("Erro ao subscrever no NATS: %v", err)
	}
}

func detectFaces(client *rekognition.Client, imageBytes []byte) (*rekognition.DetectFacesOutput, error) {
	input := &rekognition.DetectFacesInput{
		Image: &types.Image{
			Bytes: imageBytes,
		},
		Attributes: []types.Attribute{types.AttributeAll},
	}

	return client.DetectFaces(context.TODO(), input)
}

func compareWithPreviousPhotos(client *rekognition.Client, currentPhoto entity.Photo, currentFaces []types.FaceDetail) (bool, string, float64) {
	return false, "", 0.0
}

func generateReason(recognized bool, facesCount int, similarity float64) string {
	if facesCount == 0 {
		return "Nenhuma face detectada"
	}
	if recognized {
		return fmt.Sprintf("Face reconhecida com %.2f%% de similaridade", similarity)
	}
	return fmt.Sprintf("%d face(s) detectada(s), mas não reconhecida(s)", facesCount)
}

func savePhotoAnalysis(photo entity.Photo, result entity.PhotoAnalysisResult, rekognitionResult *rekognition.DetectFacesOutput, matchedID string, similarity float64, duration int64) error {
	return dynamo.SavePhotoAnalysis(photo, result, rekognitionResult, matchedID, similarity, duration)
}
