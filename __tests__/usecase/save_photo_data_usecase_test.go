package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/mkafonso/go-cloud-challenge/recognition/provider"
	repository "github.com/mkafonso/go-cloud-challenge/repository/memory"
	"github.com/mkafonso/go-cloud-challenge/usecase"
)

func TestSavePhoto_ShouldSaveAndRecognizeSuccessfully(t *testing.T) {
	repo := repository.NewInMemoryPhotoRepository()
	recognizer := provider.NewInMemoryFaceRecognition()

	uc := usecase.NewSavePhoto(repo, recognizer)

	request := &usecase.SavePhotoRequest{
		DeviceID:  "00:11:22:33:44:55",
		FilePath:  "/photos/image.jpg",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	response, err := uc.Execute(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.True(t, response.Recognized)
}

func TestSavePhoto_ShouldFailWithInvalidTimestamp(t *testing.T) {
	repo := repository.NewInMemoryPhotoRepository()
	recognizer := provider.NewInMemoryFaceRecognition()

	uc := usecase.NewSavePhoto(repo, recognizer)

	request := &usecase.SavePhotoRequest{
		DeviceID:  "00:11:22:33:44:55",
		FilePath:  "/photos/image.jpg",
		Timestamp: "invalid-timestamp",
	}

	response, err := uc.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.EqualError(t, err, "invalid timestamp format")
}
