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

func TestSaveGyroscopeData_ShouldSaveSuccessfully(t *testing.T) {
	repo := repository.NewInMemoryGyroscopeRepository()
	uc := usecase.NewSaveGyroscopeData(repo)

	request := &usecase.SaveGyroscopeDataRequest{
		DeviceID:  "00:11:22:33:44:55",
		X:         utils.Ptr(1.0),
		Y:         utils.Ptr(2.0),
		Z:         utils.Ptr(3.0),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	response, err := uc.Execute(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestSaveGyroscopeData_ShouldFailWithInvalidTimestamp(t *testing.T) {
	repo := repository.NewInMemoryGyroscopeRepository()
	uc := usecase.NewSaveGyroscopeData(repo)

	request := &usecase.SaveGyroscopeDataRequest{
		DeviceID:  "00:11:22:33:44:55",
		X:         utils.Ptr(1.0),
		Y:         utils.Ptr(2.0),
		Z:         utils.Ptr(3.0),
		Timestamp: "invalid-timestamp",
	}

	response, err := uc.Execute(context.Background(), request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.EqualError(t, err, "invalid timestamp format")
}
