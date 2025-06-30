package repository

import "challenge-cloud/internal/models"

type GPSRepositoryInterface interface {
	Create(data *models.GPS) error
	GetAll(page, size int) ([]models.GPS, error)
}
