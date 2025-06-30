// Exemplo de repository
package repository

import (
	"challenge-cloud/internal/models"

	"gorm.io/gorm"
)

type PhotoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{DB: db}
}

func (r *PhotoRepository) Create(data *models.Photo) error {
	return r.DB.Create(data).Error
}

func (r *PhotoRepository) GetAll(page int, pageSize int) ([]models.Photo, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	var results []models.Photo
	offset := (page - 1) * pageSize

	err := r.DB.
		Limit(pageSize).
		Offset(offset).
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
