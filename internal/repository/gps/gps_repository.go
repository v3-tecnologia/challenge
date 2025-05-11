package gps

import (
	"github.com/iamrosada0/v3/internal/domain"
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
