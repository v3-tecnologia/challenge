package usecase

import (
	"context"

	"github.com/mkafonso/go-cloud-challenge/entity"
	"github.com/mkafonso/go-cloud-challenge/repository"
	"github.com/mkafonso/go-cloud-challenge/utils"
)

type SaveGPSDataRequest struct {
	DeviceID  string
	Latitude  float64
	Longitude float64
	Timestamp string
}

type SaveGPSDataResponse struct{}

type SaveGPSData struct {
	repo repository.GPSRepositoryInterface
}

func NewSaveGPSData(repo repository.GPSRepositoryInterface) *SaveGPSData {
	return &SaveGPSData{repo: repo}
}

func (uc *SaveGPSData) Execute(ctx context.Context, data *SaveGPSDataRequest) (*SaveGPSDataResponse, error) {
	timestamp, err := utils.ParseRFC3339(data.Timestamp)
	if err != nil {
		return nil, err
	}

	gps, err := entity.NewGPS(data.DeviceID, data.Latitude, data.Longitude, timestamp)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.SaveGPS(ctx, gps); err != nil {
		return nil, err
	}

	return &SaveGPSDataResponse{}, nil
}
