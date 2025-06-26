package telemetriesUsecases

import (
	"time"
	"v3-test/internal/dtos/telemetriesDtos"
	"v3-test/internal/models/telemetriesModels"
	"v3-test/internal/repositories/telemetriesRepositories"
)

type IGpsUsecase interface {
	CreateGps(gpsDto telemetriesDtos.CreateGpsDto) (telemetriesModels.GpsModel, error)
}

type GpsUsecase struct {
	repo telemetriesRepositories.GpsRepository
}

func NewGpsUsecase(repo telemetriesRepositories.GpsRepository) GpsUsecase {
	return GpsUsecase{repo: repo}
}

func (u *GpsUsecase) CreateGps(gpsDto telemetriesDtos.CreateGpsDto) (telemetriesModels.GpsModel, error) {
	gpsModel := telemetriesModels.GpsModel{
		Latitude:  gpsDto.Latitude,
		Longitude: gpsDto.Longitude,
		Timestamp: time.Now(),
	}

	newGps, err := u.repo.CreateGps(gpsModel)
	if err != nil {
		return telemetriesModels.GpsModel{}, err
	}

	return newGps, nil
}
