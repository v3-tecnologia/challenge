package usecases

import (
	"time"
	dtos "v3-test/internal/dtos/telemetries"
	models "v3-test/internal/models/telemetries"
)

type GpsUsecase struct{}

func NewGpsUsecase() GpsUsecase {
	return GpsUsecase{}
}

func (gpsUsecase *GpsUsecase) CreateGps(gpsDto dtos.GpsDto) (models.GpsModel, error) {
	gpsModel := models.GpsModel{
		Latitude:  gpsDto.Latitude,
		Longitude: gpsDto.Longitude,
		Timestamp: time.Now(),
	}
	return gpsModel, nil
}
