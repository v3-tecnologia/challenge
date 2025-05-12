package gps

import (
	"errors"
	"fmt"
	"v3/internal/domain"

	"gorm.io/gorm"
)

var (
	ErrDeviceIDEmpty  = errors.New("NOT NULL constraint failed: gps.device_id")
	ErrTimestampEmpty = errors.New("NOT NULL constraint failed: gps.timestamp")
	ErrLatitudeEmpty  = errors.New("NOT NULL constraint failed: gps.latitude")
	ErrLongitudeEmpty = errors.New("NOT NULL constraint failed: gps.longitude")
	ErrCreateFailed   = errors.New("failed to create GPS record")
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

	if d.DeviceID == "" {
		return nil, ErrDeviceIDEmpty
	}
	if d.Timestamp.IsZero() {
		return nil, ErrTimestampEmpty
	}
	if d.Latitude == 0 {
		return nil, ErrLatitudeEmpty
	}
	if d.Longitude == 0 {
		return nil, ErrLongitudeEmpty
	}

	if err := r.DB.Create(d).Error; err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateFailed, err)
	}

	return d, nil
}
