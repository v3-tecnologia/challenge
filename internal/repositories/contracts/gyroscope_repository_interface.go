package repository

import "challenge-cloud/internal/models"

type GyroscopeRepositoryInterface interface {
	Create(data *models.Gyroscope) error
	GetAll(page, size int) ([]models.Gyroscope, error)
}
