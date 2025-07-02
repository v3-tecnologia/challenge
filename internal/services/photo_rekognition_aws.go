package services

import (
	"context"
	"crypto/sha256"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

type PhotoRecognitionResult struct {
	Recognized bool
	Details    string
	Timestamp  time.Time // Adiciona timestamp para rastreabilidade
}

type PhotoRecognitionService interface {
	AnalyzePhoto(photo []byte) (PhotoRecognitionResult, error)
	GetPreviousPhotos() []PhotoRecognitionResult // Permite acesso às fotos anteriores
}

type PhotoCache struct {
	mu    sync.Mutex
	cache map[[32]byte]PhotoRecognitionResult
	order [][32]byte // Para manter a ordem de chegada das fotos
}

func NewPhotoCache() *PhotoCache {
	return &PhotoCache{cache: make(map[[32]byte]PhotoRecognitionResult), order: make([][32]byte, 0)}
}

func (c *PhotoCache) Get(photo []byte) (PhotoRecognitionResult, bool) {
	hash := sha256.Sum256(photo)
	c.mu.Lock()
	defer c.mu.Unlock()
	result, ok := c.cache[hash]
	return result, ok
}

func (c *PhotoCache) Set(photo []byte, result PhotoRecognitionResult) {
	hash := sha256.Sum256(photo)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[hash] = result
	c.order = append(c.order, hash)
}

func (c *PhotoCache) All() []PhotoRecognitionResult {
	c.mu.Lock()
	defer c.mu.Unlock()
	results := make([]PhotoRecognitionResult, 0, len(c.order))
	for _, hash := range c.order {
		results = append(results, c.cache[hash])
	}
	return results
}

type AWSPhotoRecognition struct {
	client *rekognition.Client
	cache  *PhotoCache
}

func NewAWSPhotoRecognition() (*AWSPhotoRecognition, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	client := rekognition.NewFromConfig(cfg)
	return &AWSPhotoRecognition{
		client: client,
		cache:  NewPhotoCache(),
	}, nil
}

func (a *AWSPhotoRecognition) AnalyzePhoto(photo []byte) (PhotoRecognitionResult, error) {
	if result, ok := a.cache.Get(photo); ok {
		log.Printf("[CACHE HIT] Foto reconhecida: %v, Detalhes: %s", result.Recognized, result.Details)
		return result, nil
	}

	input := &rekognition.DetectLabelsInput{
		Image: &types.Image{
			Bytes: photo,
		},
		MaxLabels:     aws.Int32(5),
		MinConfidence: aws.Float32(80),
	}
	output, err := a.client.DetectLabels(context.TODO(), input)
	if err != nil {
		log.Printf("[AWS ERROR] Erro Rekognition: %v", err)
		return PhotoRecognitionResult{Recognized: false, Details: "AWS error", Timestamp: time.Now()}, err
	}
	recognized := len(output.Labels) > 0
	details := ""
	for _, label := range output.Labels {
		details += *label.Name + " "
	}
	result := PhotoRecognitionResult{Recognized: recognized, Details: details, Timestamp: time.Now()}
	a.cache.Set(photo, result)
	log.Printf("[AWS REKOGNITION] Foto analisada. Reconhecida: %v, Detalhes: %s", recognized, details)
	// Log de comparação com fotos anteriores
	for _, prev := range a.cache.All() {
		if prev.Details == details && prev.Timestamp != result.Timestamp {
			log.Printf("[COMPARISON] Foto atual reconhecida igual a uma anterior (timestamp: %v)", prev.Timestamp)
		}
	}
	return result, nil
}

func (a *AWSPhotoRecognition) GetPreviousPhotos() []PhotoRecognitionResult {
	return a.cache.All()
}

// Para funcionar na AWS, use as credenciais padrão do ambiente (env vars AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)
// Para forçar uma região, adicione:
// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))

// MOCK
type MockPhotoRecognition struct{}

func (m *MockPhotoRecognition) AnalyzePhoto(photo []byte) (PhotoRecognitionResult, error) {
	return PhotoRecognitionResult{
		Recognized: true,
		Details:    "Mocked label",
		Timestamp:  time.Now(),
	}, nil
}

func (m *MockPhotoRecognition) GetPreviousPhotos() []PhotoRecognitionResult {
	return nil
}
