package usecase

import (
	"challenge-v3-backend/internal/domain/entity"
	"challenge-v3-backend/internal/interface/dto"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockPictureGateway struct {
	mock.Mock
}

func (m *MockPictureGateway) CreatePictures(ctx context.Context, picture *entity.Picture) (*entity.Picture, error) {
	args := m.Called(ctx, picture)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Picture), args.Error(1)
}

type MockS3Client struct {
	mock.Mock
}

func (m *MockS3Client) UploadImage(ctx context.Context, imageData []byte, deviceID string) (string, error) {
	args := m.Called(ctx, imageData, deviceID)
	return args.String(0), args.Error(1)
}

func (m *MockS3Client) GetSignedURL(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

type MockRekognitionClient struct {
	mock.Mock
}

func (m *MockRekognitionClient) IndexFace(ctx context.Context, s3Key string, deviceID string) (string, error) {
	args := m.Called(ctx, s3Key, deviceID)
	return args.String(0), args.Error(1)
}

func (m *MockRekognitionClient) CompareFaces(ctx context.Context, s3Key string, deviceID string) (bool, float64, error) {
	args := m.Called(ctx, s3Key, deviceID)
	return args.Bool(0), args.Get(1).(float64), args.Error(2)
}

func TestPicturesUseCase_Create_Success(t *testing.T) {
	mockGateway := new(MockPictureGateway)
	mockS3Client := new(MockS3Client)
	mockRekognition := new(MockRekognitionClient)

	useCase := NewPicturesUseCase(mockGateway, mockS3Client, mockRekognition)

	ctx := context.Background()
	now := time.Now()
	deviceID := "test-device-123"
	pictureData := []byte("test-image-data")
	s3Key := "images/test-device-123/image.jpg"
	signedURL := "https://example.com/image.jpg"
	faceID := "face-123"

	input := dto.CreatePictureRequestDTO{
		DeviceId:    deviceID,
		CreatedAt:   now,
		PictureData: pictureData,
		PictureType: "SELFIE",
	}

	mockS3Client.On("UploadImage", ctx, pictureData, deviceID).Return(s3Key, nil)
	mockS3Client.On("GetSignedURL", ctx, s3Key).Return(signedURL, nil)
	mockRekognition.On("IndexFace", ctx, s3Key, deviceID).Return(faceID, nil)
	mockRekognition.On("CompareFaces", ctx, s3Key, deviceID).Return(true, float64(90.00), nil)

	expectedID := uuid.New()
	expectedPicture := &entity.Picture{
		BaseEntity: entity.BaseEntity{
			ID:         expectedID,
			DeviceID:   deviceID,
			CreatedAt:  now,
			ReceivedAt: time.Now(),
		},
		PictureURL:     signedURL,
		RecognizedFace: true,
	}

	mockGateway.On("CreatePictures", ctx, mock.MatchedBy(func(p *entity.Picture) bool {
		return p.DeviceID == deviceID &&
			p.PictureURL == signedURL &&
			p.RecognizedFace == true
	})).Return(expectedPicture, nil)

	result, err := useCase.Create(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedPicture.ID.String(), result.Id)
	assert.Equal(t, deviceID, result.DeviceId)
	assert.Equal(t, signedURL, result.PictureURL)
	assert.Equal(t, true, result.IsRecognized)

	mockS3Client.AssertExpectations(t)
	mockRekognition.AssertExpectations(t)
	mockGateway.AssertExpectations(t)
}

func TestPicturesUseCase_Create_UploadImageError(t *testing.T) {
	mockGateway := new(MockPictureGateway)
	mockS3Client := new(MockS3Client)
	mockRekognition := new(MockRekognitionClient)

	useCase := NewPicturesUseCase(mockGateway, mockS3Client, mockRekognition)

	ctx := context.Background()
	now := time.Now()
	deviceID := "test-device-123"
	pictureData := []byte("test-image-data")

	input := dto.CreatePictureRequestDTO{
		DeviceId:    deviceID,
		CreatedAt:   now,
		PictureData: pictureData,
		PictureType: "SELFIE",
	}

	expectedError := errors.New("failed to upload image")
	mockS3Client.On("UploadImage", ctx, pictureData, deviceID).Return("", expectedError)

	result, err := useCase.Create(ctx, input)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, result)

	mockS3Client.AssertExpectations(t)
	mockRekognition.AssertNotCalled(t, "IndexFace")
	mockGateway.AssertNotCalled(t, "CreatePictures")
}

func TestPicturesUseCase_Create_GetSignedURLError(t *testing.T) {
	mockGateway := new(MockPictureGateway)
	mockS3Client := new(MockS3Client)
	mockRekognition := new(MockRekognitionClient)

	useCase := NewPicturesUseCase(mockGateway, mockS3Client, mockRekognition)

	ctx := context.Background()
	now := time.Now()
	deviceID := "test-device-123"
	pictureData := []byte("test-image-data")
	s3Key := "images/test-device-123/image.jpg"

	input := dto.CreatePictureRequestDTO{
		DeviceId:    deviceID,
		CreatedAt:   now,
		PictureData: pictureData,
		PictureType: "SELFIE",
	}

	mockS3Client.On("UploadImage", ctx, pictureData, deviceID).Return(s3Key, nil)

	expectedError := errors.New("failed to get signed URL")
	mockS3Client.On("GetSignedURL", ctx, s3Key).Return("", expectedError)

	result, err := useCase.Create(ctx, input)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, result)

	mockS3Client.AssertExpectations(t)
	mockRekognition.AssertNotCalled(t, "IndexFace")
	mockGateway.AssertNotCalled(t, "CreatePictures")
}

func TestPicturesUseCase_Create_IndexFaceError(t *testing.T) {
	mockGateway := new(MockPictureGateway)
	mockS3Client := new(MockS3Client)
	mockRekognition := new(MockRekognitionClient)

	useCase := NewPicturesUseCase(mockGateway, mockS3Client, mockRekognition)

	ctx := context.Background()
	now := time.Now()
	deviceID := "test-device-123"
	pictureData := []byte("test-image-data")
	s3Key := "images/test-device-123/image.jpg"
	signedURL := "https://example.com/image.jpg"

	input := dto.CreatePictureRequestDTO{
		DeviceId:    deviceID,
		CreatedAt:   now,
		PictureData: pictureData,
		PictureType: "SELFIE",
	}

	mockS3Client.On("UploadImage", ctx, pictureData, deviceID).Return(s3Key, nil)
	mockS3Client.On("GetSignedURL", ctx, s3Key).Return(signedURL, nil)

	expectedError := errors.New("failed to index face")
	mockRekognition.On("IndexFace", ctx, s3Key, deviceID).Return("", expectedError)

	result, err := useCase.Create(ctx, input)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, result)

	mockS3Client.AssertExpectations(t)
	mockRekognition.AssertExpectations(t)
	mockGateway.AssertNotCalled(t, "CreatePictures")
}

func TestPicturesUseCase_Create_CreatePicturesError(t *testing.T) {
	mockGateway := new(MockPictureGateway)
	mockS3Client := new(MockS3Client)
	mockRekognition := new(MockRekognitionClient)

	useCase := NewPicturesUseCase(mockGateway, mockS3Client, mockRekognition)

	ctx := context.Background()
	now := time.Now()
	deviceID := "test-device-123"
	pictureData := []byte("test-image-data")
	s3Key := "images/test-device-123/image.jpg"
	signedURL := "https://example.com/image.jpg"
	faceID := "face-123"

	input := dto.CreatePictureRequestDTO{
		DeviceId:    deviceID,
		CreatedAt:   now,
		PictureData: pictureData,
		PictureType: "SELFIE",
	}

	mockS3Client.On("UploadImage", ctx, pictureData, deviceID).Return(s3Key, nil)
	mockS3Client.On("GetSignedURL", ctx, s3Key).Return(signedURL, nil)
	mockRekognition.On("IndexFace", ctx, s3Key, deviceID).Return(faceID, nil)
	mockRekognition.On("CompareFaces", ctx, s3Key, deviceID).Return(true, float64(90.00), nil)

	expectedError := errors.New("failed to create picture")
	mockGateway.On("CreatePictures", ctx, mock.Anything).Return(nil, expectedError)

	result, err := useCase.Create(ctx, input)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, result)

	mockS3Client.AssertExpectations(t)
	mockRekognition.AssertExpectations(t)
	mockGateway.AssertExpectations(t)
}
