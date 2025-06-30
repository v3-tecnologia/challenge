// Exemplo de repository
package repository

import (
	"challenge-cloud/internal/models"

	"gorm.io/gorm"
)

type GPSRepository struct {
	DB *gorm.DB
}

func NewGPSRepository(db *gorm.DB) *GPSRepository {
	return &GPSRepository{DB: db}
}

func (r *GPSRepository) Create(data *models.GPS) error {
	return r.DB.Create(data).Error
}

func (r *GPSRepository) GetAll(page int, pageSize int) ([]models.GPS, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	var results []models.GPS
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
