package repository

import (
	"gorm.io/gorm"
	"v3-backend-challenge/src/model"
)

type GpsRepository struct {
	DB *gorm.DB
}

func (r *GpsRepository) Save(photo *model.GPS) error {
	return r.DB.Create(photo).Error
}
