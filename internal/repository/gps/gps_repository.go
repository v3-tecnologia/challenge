package gps

import (
	"errors"
	"v3/internal/domain"

	"gorm.io/gorm"
)

type GPSRepository interface {
	Create(d *domain.GPS) (*domain.GPS, error)
}

type gpsRepository struct {
	DB *gorm.DB
}

func NewGPSRepository(db *gorm.DB) GPSRepository {
	return &gpsRepository{DB: db}
}

func (r *gpsRepository) Create(d *domain.GPS) (*domain.GPS, error) {
	// Validate required fields
	if d.DeviceID == "" {
		return nil, errors.New("NOT NULL constraint failed: gps.device_id")
	}
	if d.Timestamp.IsZero() {
		return nil, errors.New("NOT NULL constraint failed: gps.timestamp")
	}
	if d.Latitude == 0 {
		return nil, errors.New("NOT NULL constraint failed: gps.latitude")
	}
	if d.Longitude == 0 {
		return nil, errors.New("NOT NULL constraint failed: gps.longitude")
	}

	if err := r.DB.Create(d).Error; err != nil {
		return nil, err
	}
	return d, nil
}
