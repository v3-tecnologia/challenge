package service

import (
	"challenge-cloud/internal/models"
	repository "challenge-cloud/internal/repositories"
)

type GPSService struct {
	Repo *repository.GPSRepository
}

func NewGPSService(r *repository.GPSRepository) *GPSService {
	return &GPSService{Repo: r}
}

func (s *GPSService) Save(data *models.GPS) error {
	return s.Repo.Create(data)
}
func (s *GPSService) GetAll(page, pageSize int) ([]models.GPS, error) {
	return s.Repo.GetAll(page, pageSize)
}
