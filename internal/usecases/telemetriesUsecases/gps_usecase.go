package telemetriesUsecases

import (
	"time"
	dtos "v3-test/internal/dtos/telemetriesDtos"
	models "v3-test/internal/models/telemetriesModels"
	repositories "v3-test/internal/repositories/telemetriesRepositories"
)

type GpsUsecase struct {
	repo repositories.GpsRepository
}

func NewGpsUsecase(repo repositories.GpsRepository) GpsUsecase {
	return GpsUsecase{repo: repo}
}

func (u *GpsUsecase) CreateGps(gpsDto dtos.CreateGpsDto) (models.GpsModel, error) {
	gpsModel := models.GpsModel{
		Latitude:  gpsDto.Latitude,
		Longitude: gpsDto.Longitude,
		Timestamp: time.Now(),
	}

	newGps, err := u.repo.CreateGps(gpsModel)
	if err != nil {
		return models.GpsModel{}, err
	}

	return newGps, nil
}
