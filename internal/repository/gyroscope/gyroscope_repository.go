package gyroscope

import (
	"errors"
	"fmt"
	"v3/internal/domain"

	"gorm.io/gorm"
)

var (
	ErrDeviceIDEmpty  = errors.New("NOT NULL constraint failed: gyroscopes.device_id")
	ErrXEmpty         = errors.New("NOT NULL constraint failed: gyroscopes.x")
	ErrYEmpty         = errors.New("NOT NULL constraint failed: gyroscopes.y")
	ErrZEmpty         = errors.New("NOT NULL constraint failed: gyroscopes.z")
	ErrTimestampEmpty = errors.New("NOT NULL constraint failed: gyroscopes.timestamp")
	ErrCreateFailed   = errors.New("failed to create gyroscope record")
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
	if d.DeviceID == "" {
		return nil, ErrDeviceIDEmpty
	}
	if d.X == 0 {
		return nil, ErrXEmpty
	}
	if d.Y == 0 {
		return nil, ErrYEmpty
	}
	if d.Z == 0 {
		return nil, ErrZEmpty
	}
	if d.Timestamp.IsZero() {
		return nil, ErrTimestampEmpty
	}

	if err := r.DB.Create(d).Error; err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateFailed, err)
	}

	return d, nil
}
