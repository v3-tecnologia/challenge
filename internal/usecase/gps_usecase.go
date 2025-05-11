package usecase

import (
	"v3/internal/domain"
	"v3/internal/repository/gps"
)

type CreateGPSUseCase struct {
	Repo gps.GPSRepository
}

func NewCreateGPSUseCase(repo gps.GPSRepository) *CreateGPSUseCase {
	return &CreateGPSUseCase{Repo: repo}
}

func (uc *CreateGPSUseCase) Execute(input domain.GPSDto) (*domain.GPS, error) {

	gpsData, err := domain.NewGPSData(&domain.GPSDto{
		DeviceID:  input.DeviceID,
		Timestamp: input.Timestamp,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
	})

	if err != nil {
		return nil, err
	}
	savedGPS, err := uc.Repo.Create(gpsData)
	if err != nil {
		return nil, domain.ErrSaveGPSData
	}

	return savedGPS, nil
}
