package services

import (
	"challenge-v3/crypto"
	"challenge-v3/ierr"
	"challenge-v3/models"
	"challenge-v3/storage"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/patrickmn/go-cache"
)

type PhotoAnalyzer interface {
	AnalyzeAndSavePhoto(data *models.PhotoData) (bool, error)
}
type RekognitionClient interface {
	SearchFacesByImage(ctx context.Context, params *rekognition.SearchFacesByImageInput, optFns ...func(*rekognition.Options)) (*rekognition.SearchFacesByImageOutput, error)
	IndexFaces(ctx context.Context, params *rekognition.IndexFacesInput, optFns ...func(*rekognition.Options)) (*rekognition.IndexFacesOutput, error)
}
type PhotoAnalyzerService struct {
	rekognitionClient RekognitionClient
	collectionID      string
	cache             *cache.Cache
	db                storage.Storage
}

func NewPhotoAnalyzerService(rekClient RekognitionClient, collID string, db storage.Storage) *PhotoAnalyzerService {
	return &PhotoAnalyzerService{
		rekognitionClient: rekClient,
		collectionID:      collID,
		cache:             cache.New(5*time.Minute, 10*time.Minute),
		db:                db,
	}
}

func (s *PhotoAnalyzerService) AnalyzeAndSavePhoto(data *models.PhotoData) (bool, error) {
	if err := data.Validate(); err != nil {
		return false, ierr.NewValidationError("dados da foto inválidos: %w", err)
	}
	slog.Info("iniciando análise da foto", "device_id", data.DeviceID)

	imageBytes, err := base64.StdEncoding.DecodeString(data.Photo)
	if err != nil {
		slog.Error("falha ao decodificar imagem base64", "error", err)
		return false, fmt.Errorf("imagem base64 inválida")
	}

	cacheKey := fmt.Sprintf("%x", sha256.Sum256(imageBytes))
	slog.Debug("chave de cache gerada para a imagem", "key", cacheKey)

	var recognized bool

	if rec, found := s.cache.Get(cacheKey); found {
		slog.Info("cache hit para imagem", "key", cacheKey)
		recognized = rec.(bool)
	} else {
		slog.Info("cache miss para imagem", "key", cacheKey)

		searchResult, searchErr := s.rekognitionClient.SearchFacesByImage(context.TODO(), &rekognition.SearchFacesByImageInput{
			CollectionId: aws.String(s.collectionID), Image: &types.Image{Bytes: imageBytes}, MaxFaces: aws.Int32(1), FaceMatchThreshold: aws.Float32(90.0),
		})
		if searchErr != nil {
			slog.Error("falha ao buscar face no rekognition", "error", searchErr)
			return false, fmt.Errorf("erro ao analisar a imagem")
		}

		if len(searchResult.FaceMatches) > 0 {
			recognized = true
			slog.Info("rosto reconhecido", "similarity", *searchResult.FaceMatches[0].Similarity, "face_id", *searchResult.FaceMatches[0].Face.FaceId)
		} else {
			recognized = false
			slog.Warn("rosto não reconhecido, tentando indexar", "device_id", data.DeviceID)
			indexResult, indexErr := s.rekognitionClient.IndexFaces(context.TODO(), &rekognition.IndexFacesInput{
				CollectionId: aws.String(s.collectionID), Image: &types.Image{Bytes: imageBytes}, MaxFaces: aws.Int32(1), DetectionAttributes: []types.Attribute{types.AttributeDefault},
			})
			if indexErr != nil || len(indexResult.FaceRecords) == 0 {
				slog.Error("falha ao indexar novo rosto", "error", indexErr)
			} else {
				slog.Info("novo rosto indexado com sucesso", "face_id", *indexResult.FaceRecords[0].Face.FaceId)
			}
		}

		if recognized {
			slog.Info("cache set para imagem", "key", cacheKey, "value", true)
			s.cache.Set(cacheKey, true, cache.DefaultExpiration)
		}
	}

	data.Recognized = recognized

	encryptionKey := []byte(os.Getenv("ENCRYPTION_KEY"))
	if len(encryptionKey) == 32 {
		slog.Info("criptografando dados da foto antes de salvar", "device_id", data.DeviceID)

		originalPhotoB64 := data.Photo
		encryptedPhotoBytes, err := crypto.Encrypt([]byte(originalPhotoB64), encryptionKey)
		if err != nil {
			slog.Error("falha ao criptografar foto", "error", err)
			return false, fmt.Errorf("erro interno de criptografia")
		}
		data.Photo = base64.StdEncoding.EncodeToString(encryptedPhotoBytes)
	}

	if err := s.db.SavePhoto(data); err != nil {
		slog.Error("falha ao salvar foto no banco de dados", "error", err)
		return false, err
	}

	slog.Info("análise e salvamento da foto concluídos", "device_id", data.DeviceID, "recognized", data.Recognized)
	return recognized, nil
}
