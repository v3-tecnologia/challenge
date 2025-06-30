// Exemplo de repository
package repository

import (
	"challenge-cloud/internal/models"

	"gorm.io/gorm"
)

type GyroscopeRepository struct {
	DB *gorm.DB
}

func NewGyroscopeRepository(db *gorm.DB) *GyroscopeRepository {
	return &GyroscopeRepository{DB: db}
}

func (r *GyroscopeRepository) Create(data *models.Gyroscope) error {
	return r.DB.Create(data).Error
}

// func (r *GyroscopeRepository) GetAll() ([]models.Gyroscope, error) {
// 	var gyroscopes []models.Gyroscope
// 	err := r.DB.Find(&gyroscopes).Error
// 	return gyroscopes, err
// }

func (r *GyroscopeRepository) GetAll(page int, pageSize int) ([]models.Gyroscope, error) {
	// Defaults
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	var results []models.Gyroscope
	offset := (page - 1) * pageSize

	err := r.DB.
		Limit(pageSize).
		Offset(offset).
		// Order("timestamp DESC"). // opcional: ordena por timestamp
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
