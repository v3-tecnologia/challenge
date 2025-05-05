package usecase

import (
	"context"

	"github.com/mkafonso/go-cloud-challenge/entity"
	"github.com/mkafonso/go-cloud-challenge/repository"
	"github.com/mkafonso/go-cloud-challenge/utils"
)

type SaveGyroscopeDataRequest struct {
	DeviceID  string
	X         float64
	Y         float64
	Z         float64
	Timestamp string
}

type SaveGyroscopeDataResponse struct{}

type SaveGyroscopeData struct {
	repo repository.GyroscopeRepositoryInterface
}

func NewSaveGyroscopeData(repo repository.GyroscopeRepositoryInterface) *SaveGyroscopeData {
	return &SaveGyroscopeData{repo: repo}
}

func (uc *SaveGyroscopeData) Execute(ctx context.Context, data *SaveGyroscopeDataRequest) (*SaveGyroscopeDataResponse, error) {
	timestamp, err := utils.ParseRFC3339(data.Timestamp)
	if err != nil {
		return nil, err
	}

	gyro, err := entity.NewGyroscope(data.DeviceID, data.X, data.Y, data.Z, timestamp)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.SaveGyroscope(ctx, gyro); err != nil {
		return nil, err
	}

	return &SaveGyroscopeDataResponse{}, nil
}
