package repository

import (
	"go-challenge/internal/models"

	"gorm.io/gorm"
)

type TelemetryRepository struct {
	DB *gorm.DB
}

func (r *TelemetryRepository) SaveGyroscopeData(data models.Gyroscope) error {
	return r.DB.Create(&data).Error
}

func (r *TelemetryRepository) SaveGPSData(data models.GPS) error {
	return r.DB.Create(&data).Error
}

func (r *TelemetryRepository) SavePhotoData(data models.Photo) error {
	return r.DB.Create(&data).Error
}
