package repository

import (
	"gorm.io/gorm"
	"v3-backend-challenge/src/model"
)

type PhotoRepository struct {
	DB *gorm.DB
}

func (r *PhotoRepository) Save(photo *model.Photo) error {
	return r.DB.Create(photo).Error
}
