package repository

import "challenge-cloud/internal/models"

type PhotoRepositoryInterface interface {
	Create(data *models.Photo) error
	GetAll(page, size int) ([]models.Photo, error)
}
