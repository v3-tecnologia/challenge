package usecase_test

import (
	"context"
	"testing"
	"time"

	repository "github.com/mkafonso/go-cloud-challenge/repository/memory"
	"github.com/mkafonso/go-cloud-challenge/usecase"
	"github.com/mkafonso/go-cloud-challenge/utils"
	"github.com/stretchr/testify/assert"
)

func TestSaveGPSData_ShouldSaveSuccessfully(t *testing.T) {
	repo := repository.NewInMemoryGPSRepository()
	uc := usecase.NewSaveGPSData(repo)

	request := &usecase.SaveGPSDataRequest{
		DeviceID:  "00:11:22:33:44:55",
		Latitude:  utils.Ptr(10.0),
		Longitude: utils.Ptr(10.0),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	response, err := uc.Execute(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestSaveGPSData_ShouldFailWithInvalidTimestamp(t *testing.T) {
	repo := repository.NewInMemoryGPSRepository()
	uc := usecase.NewSaveGPSData(repo)

	request := &usecase.SaveGPSDataRequest{
		DeviceID:  "00:11:22:33:44:55",
		Latitude:  utils.Ptr(10.0),
		Longitude: utils.Ptr(10.0),
		Timestamp: "invalid-timestamp",
	}

	response, err := uc.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.EqualError(t, err, "invalid timestamp format")
}
