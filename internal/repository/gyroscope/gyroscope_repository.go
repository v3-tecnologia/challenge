package gyroscope

import (
	"errors"
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
