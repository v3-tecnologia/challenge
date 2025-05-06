package usecase_test

import (
	"context"
	"testing"
	"time"

	recognition "github.com/mkafonso/go-cloud-challenge/recognition/provider"
	repository "github.com/mkafonso/go-cloud-challenge/repository/memory"
	storage "github.com/mkafonso/go-cloud-challenge/storage/provider"
	"github.com/mkafonso/go-cloud-challenge/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSavePhoto_ShouldSaveAndRecognizeSuccessfully(t *testing.T) {
	repo := repository.NewInMemoryPhotoRepository()
	recognizer := recognition.NewInMemoryFaceRecognition()
	storage := storage.NewInMemoryPhotoStorage()

	uc := usecase.NewSavePhoto(repo, recognizer, storage)

	request := &usecase.SavePhotoRequest{
		DeviceID:  "00:11:22:33:44:55",
		FileBytes: []byte("random image bytes"),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	response, err := uc.Execute(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.True(t, response.Recognized)
}

func TestSavePhoto_ShouldFailWithInvalidTimestamp(t *testing.T) {
	repo := repository.NewInMemoryPhotoRepository()
	recognizer := recognition.NewInMemoryFaceRecognition()
	storage := storage.NewInMemoryPhotoStorage()

	uc := usecase.NewSavePhoto(repo, recognizer, storage)

	request := &usecase.SavePhotoRequest{
		DeviceID:  "00:11:22:33:44:55",
		FileBytes: []byte("random image bytes"),
		Timestamp: "not-a-valid-timestamp",
	}

	response, err := uc.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.EqualError(t, err, "invalid timestamp format")
}
