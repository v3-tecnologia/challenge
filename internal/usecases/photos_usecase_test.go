package usecases

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ricardoraposo/challenge/internal/interfaces"
	"github.com/ricardoraposo/challenge/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPhotosQueries struct {
	mock.Mock
}

func (m *MockPhotosQueries) GetDeviceByID(ctx context.Context, deviceID string) (repository.Device, error) {
	args := m.Called(ctx, deviceID)
	return args.Get(0).(repository.Device), args.Error(1)
}

func (m *MockPhotosQueries) InsertDevice(ctx context.Context, deviceID string) (repository.Device, error) {
	args := m.Called(ctx, deviceID)
	return args.Get(0).(repository.Device), args.Error(1)
}

func (m *MockPhotosQueries) InsertPhoto(ctx context.Context, arg repository.InsertPhotoParams) (repository.Photo, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repository.Photo), args.Error(1)
}

type MockBucketUploader struct {
	mock.Mock
}

func (m *MockBucketUploader) UploadAsync(ctx context.Context, file io.Reader, key string, resultCh chan<- string, errCh chan<- error) {
	m.Called(ctx, file, key, resultCh, errCh)
}

type MockFaceDetector struct {
	mock.Mock
}

func (m *MockFaceDetector) HandleFaceRecognition(ctx context.Context, imageKey string) (*[]interfaces.FaceMatch, error) {
	args := m.Called(ctx, imageKey)
	return args.Get(0).(*[]interfaces.FaceMatch), args.Error(1)
}

func Test_CreatePhoto_Success(t *testing.T) {
    t.Parallel()

	mockQueries := new(MockPhotosQueries)
	mockUploader := new(MockBucketUploader)
	mockFaceDetector := new(MockFaceDetector)
	uc := NewPhotosUseCase(mockQueries, mockUploader, mockFaceDetector)

	ctx := context.Background()
	testDeviceID := "test-device-123"
	testPhotoKey := "photos/test-image.png"
	testImageURL := "https://s3.example.com/" + testPhotoKey
	now := time.Now()
	pgNow := pgtype.Timestamp{Time: now, Valid: true}

	params := CreatePhotoParams{
		DeviceID:    testDeviceID,
		File:        strings.NewReader("fake image data"),
		Key:         testPhotoKey,
		CollectedAt: pgNow,
	}

	expectedDevice := repository.Device{
		DeviceID:     testDeviceID,
		RegisteredAt: pgNow,
	}

	expectedPhoto := repository.Photo{
		ID:            uuid.New(),
		DeviceID:      testDeviceID,
		ImageUrl:      testImageURL,
		RecurrentUser: pgtype.Bool{Bool: false, Valid: true},
		CollectedAt:   pgNow,
	}

	mockQueries.On("GetDeviceByID", ctx, testDeviceID).Return(repository.Device{}, sql.ErrNoRows)
	mockQueries.On("InsertDevice", ctx, testDeviceID).Return(expectedDevice, nil)

	mockUploader.On("UploadAsync", mock.AnythingOfType("*context.cancelCtx"), params.File, params.Key, mock.AnythingOfType("chan<- string"), mock.AnythingOfType("chan<- error")).
		Run(func(args mock.Arguments) {
			resultCh := args.Get(3).(chan<- string)
			go func() {
				resultCh <- testImageURL
				close(resultCh)
			}()
		}).Return()

	mockFaceDetector.On("HandleFaceRecognition", mock.AnythingOfType("*context.cancelCtx"), params.Key).Return(&[]interfaces.FaceMatch{}, nil)

	insertPhotoParamsWithFace := repository.InsertPhotoParams{
		DeviceID:      testDeviceID,
		ImageUrl:      testImageURL,
		CollectedAt:   pgNow,
		RecurrentUser: pgtype.Bool{Bool: false, Valid: true},
	}
	mockQueries.On("InsertPhoto", ctx, insertPhotoParamsWithFace).Return(expectedPhoto, nil)

	photo, err := uc.CreatePhoto(ctx, params)

	assert.NoError(t, err)
	assert.Equal(t, expectedPhoto, photo)

	mockQueries.AssertExpectations(t)
	mockUploader.AssertExpectations(t)
	mockFaceDetector.AssertExpectations(t)
}

func TestPhotosUseCase_CreatePhoto_UploadFails(t *testing.T) {
    t.Parallel()

	mockQueries := new(MockPhotosQueries)
	mockUploader := new(MockBucketUploader)
	mockFaceDetector := new(MockFaceDetector)
	uc := NewPhotosUseCase(mockQueries, mockUploader, mockFaceDetector)

	ctx := context.Background()
	testDeviceID := "test-device-upload-fail"
	testPhotoKey := "photos/test-image-upload-fail.png"
	now := time.Now()
	pgNow := pgtype.Timestamp{Time: now, Valid: true}

	params := CreatePhotoParams{
		DeviceID:    testDeviceID,
		File:        strings.NewReader("fake image data"),
		Key:         testPhotoKey,
		CollectedAt: pgNow,
	}

	expectedDevice := repository.Device{
		DeviceID:     testDeviceID,
		RegisteredAt: pgNow,
	}

	uploadError := errors.New("failed to upload to S3")

	mockQueries.On("GetDeviceByID", ctx, testDeviceID).Return(repository.Device{}, sql.ErrNoRows)
	mockQueries.On("InsertDevice", ctx, testDeviceID).Return(expectedDevice, nil)

	mockUploader.On("UploadAsync", mock.AnythingOfType("*context.cancelCtx"), params.File, params.Key, mock.AnythingOfType("chan<- string"), mock.AnythingOfType("chan<- error")).
		Run(func(args mock.Arguments) {
			errCh := args.Get(4).(chan<- error)
			go func() {
				errCh <- uploadError
				close(errCh)
			}()
		}).Return()

	photo, err := uc.CreatePhoto(ctx, params)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "upload failed")
	assert.Contains(t, err.Error(), uploadError.Error())
	assert.Equal(t, repository.Photo{}, photo)

	mockQueries.AssertExpectations(t)
	mockUploader.AssertExpectations(t)

	mockQueries.AssertNotCalled(t, "InsertPhoto", mock.Anything, mock.Anything)
}
