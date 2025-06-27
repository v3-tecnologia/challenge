package services

import (
	"challenge-v3/crypto"
	"challenge-v3/ierr"
	"challenge-v3/models"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockRekognitionClient struct{ mock.Mock }

func (m *MockRekognitionClient) SearchFacesByImage(ctx context.Context, params *rekognition.SearchFacesByImageInput, optFns ...func(*rekognition.Options)) (*rekognition.SearchFacesByImageOutput, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*rekognition.SearchFacesByImageOutput), args.Error(1)
}
func (m *MockRekognitionClient) IndexFaces(ctx context.Context, params *rekognition.IndexFacesInput, optFns ...func(*rekognition.Options)) (*rekognition.IndexFacesOutput, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*rekognition.IndexFacesOutput), args.Error(1)
}

type MockStorage struct{ mock.Mock }

func (m *MockStorage) SavePhoto(data *models.PhotoData) error         { return m.Called(data).Error(0) }
func (m *MockStorage) SaveGyroscope(data *models.GyroscopeData) error { return m.Called(data).Error(0) }
func (m *MockStorage) SaveGPS(data *models.GPSData) error             { return m.Called(data).Error(0) }
func (m *MockStorage) LogAuditEvent(event models.AuditEvent) error {
	args := m.Called(event)
	return args.Error(0)
}

func validTestPhoto() models.PhotoData {
	return models.PhotoData{
		DeviceID:  "test-device",
		Photo:     "aW1hZ2VtLWRhdGE=",
		Timestamp: time.Now(),
	}
}

func TestPhotoAnalyzer_FaceRecognized(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Log("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema/CI.")
	}
	encryptionKey := []byte(os.Getenv("ENCRYPTION_KEY"))
	require.Len(t, encryptionKey, 32, "A chave de criptografia deve ter 32 bytes")

	mockRek := new(MockRekognitionClient)
	mockDB := new(MockStorage)
	photoAnalyzer := NewPhotoAnalyzerService(mockRek, "test-collection", mockDB)
	testPhoto := validTestPhoto()
	originalPhotoB64 := testPhoto.Photo

	faceID, similarity := "test-face-id", float32(99.9)
	searchOutput := &rekognition.SearchFacesByImageOutput{
		FaceMatches: []types.FaceMatch{{Face: &types.Face{FaceId: &faceID}, Similarity: &similarity}},
	}
	mockRek.On("SearchFacesByImage", mock.Anything, mock.Anything).Return(searchOutput, nil)
	mockDB.On("SavePhoto", mock.MatchedBy(func(p *models.PhotoData) bool {
		if !p.Recognized {
			return false
		}
		if p.Photo == originalPhotoB64 {
			return false
		}
		encryptedBytes, _ := base64.StdEncoding.DecodeString(p.Photo)
		decryptedBytes, err := crypto.Decrypt(encryptedBytes, encryptionKey)
		if err != nil {
			return false
		}
		return string(decryptedBytes) == originalPhotoB64
	})).Return(nil)

	recognized, err := photoAnalyzer.AnalyzeAndSavePhoto(&testPhoto)

	assert.NoError(t, err)
	assert.True(t, recognized)
	mockRek.AssertExpectations(t)
	mockDB.AssertExpectations(t)

	imageBytes, _ := base64.StdEncoding.DecodeString(originalPhotoB64)
	cacheKey := fmt.Sprintf("%x", sha256.Sum256(imageBytes))
	cachedResult, found := photoAnalyzer.cache.Get(cacheKey)
	assert.True(t, found, "O resultado 'true' deveria ter sido salvo no cache")
	assert.Equal(t, true, cachedResult)
}

func TestPhotoAnalyzer_FaceNotRecognized_AndIndexed(t *testing.T) {
	mockRek := new(MockRekognitionClient)
	mockDB := new(MockStorage)
	photoAnalyzer := NewPhotoAnalyzerService(mockRek, "test-collection", mockDB)
	testPhoto := validTestPhoto()

	mockRek.On("SearchFacesByImage", mock.Anything, mock.Anything).Return(&rekognition.SearchFacesByImageOutput{}, nil)
	faceID := "new-face-id"
	indexOutput := &rekognition.IndexFacesOutput{FaceRecords: []types.FaceRecord{{Face: &types.Face{FaceId: &faceID}}}}
	mockRek.On("IndexFaces", mock.Anything, mock.Anything).Return(indexOutput, nil)
	mockDB.On("SavePhoto", mock.Anything).Return(nil)

	recognized, err := photoAnalyzer.AnalyzeAndSavePhoto(&testPhoto)

	assert.NoError(t, err)
	assert.False(t, recognized)
	mockRek.AssertExpectations(t)
	mockDB.AssertExpectations(t)

	imageBytes, _ := base64.StdEncoding.DecodeString(testPhoto.Photo)
	cacheKey := fmt.Sprintf("%x", sha256.Sum256(imageBytes))
	_, found := photoAnalyzer.cache.Get(cacheKey)
	assert.False(t, found, "Um resultado 'false' não deveria ser salvo no cache")
}

func TestPhotoAnalyzer_CacheHit(t *testing.T) {
	mockRek := new(MockRekognitionClient)
	mockDB := new(MockStorage)
	photoAnalyzer := NewPhotoAnalyzerService(mockRek, "test-collection", mockDB)
	testPhoto := validTestPhoto()

	imageBytes, _ := base64.StdEncoding.DecodeString(testPhoto.Photo)
	cacheKey := fmt.Sprintf("%x", sha256.Sum256(imageBytes))
	photoAnalyzer.cache.Set(cacheKey, true, cache.DefaultExpiration)

	mockDB.On("SavePhoto", mock.Anything).Return(nil)

	recognized, err := photoAnalyzer.AnalyzeAndSavePhoto(&testPhoto)

	assert.NoError(t, err)
	assert.True(t, recognized)
	mockDB.AssertExpectations(t)
	mockRek.AssertNotCalled(t, "SearchFacesByImage", mock.Anything, mock.Anything)
	mockRek.AssertNotCalled(t, "IndexFaces", mock.Anything, mock.Anything)
}

func TestPhotoAnalyzer_ValidationFail(t *testing.T) {
	mockRek := new(MockRekognitionClient)
	mockDB := new(MockStorage)
	photoAnalyzer := NewPhotoAnalyzerService(mockRek, "test-collection", mockDB)

	testPhoto := models.PhotoData{Photo: "dGVzdA=="}

	recognized, err := photoAnalyzer.AnalyzeAndSavePhoto(&testPhoto)

	assert.False(t, recognized)
	assert.Error(t, err)

	var validationErr *ierr.ValidationError
	assert.ErrorAs(t, err, &validationErr, "O erro deveria ser do tipo ValidationError")
}
