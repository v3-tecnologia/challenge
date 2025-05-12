package gyroscope

import (
	"errors"
	"v3/internal/domain"

	"gorm.io/gorm"
)

type GyroscopeRepository interface {
	Create(d *domain.Gyroscope) (*domain.Gyroscope, error)
}

type gyroscopeRepository struct {
	DB *gorm.DB
}

func NewGyroscopeRepository(db *gorm.DB) GyroscopeRepository {
	return &gyroscopeRepository{DB: db}
}

func (r *gyroscopeRepository) Create(d *domain.Gyroscope) (*domain.Gyroscope, error) {
	// Validate required fields
	if d.DeviceID == "" {
		return nil, errors.New("NOT NULL constraint failed: gyroscopes.device_id")
	}
	if d.X == 0 {
		return nil, errors.New("NOT NULL constraint failed: gyroscopes.x")
	}
	if d.Y == 0 {
		return nil, errors.New("NOT NULL constraint failed: gyroscopes.y")
	}
	if d.Z == 0 {
		return nil, errors.New("NOT NULL constraint failed: gyroscopes.z")
	}
	if d.Timestamp.IsZero() {
		return nil, errors.New("NOT NULL constraint failed: gyroscopes.timestamp")
	}

	if err := r.DB.Create(d).Error; err != nil {
		return nil, err
	}
	return d, nil
}
