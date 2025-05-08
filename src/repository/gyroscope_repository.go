package repository

import (
	"gorm.io/gorm"
	"v3-backend-challenge/model"
)

type GyroscopeRepository struct {
	DB *gorm.DB
}

func (r *GyroscopeRepository) Save(gyroscope *model.Gyroscope) error {
	return r.DB.Create(gyroscope).Error
}
