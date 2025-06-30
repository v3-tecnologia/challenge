package service

import (
	"challenge-cloud/internal/models"
	repository "challenge-cloud/internal/repositories/contracts"
)

type GPSService struct {
	Repo repository.GPSRepositoryInterface
}
type GPSServiceInterface interface {
	Save(data *models.GPS) error
	GetAll(page, size int) ([]models.GPS, error)
}

func NewGPSService(r repository.GPSRepositoryInterface) *GPSService {
	return &GPSService{Repo: r}
}

func (s *GPSService) Save(data *models.GPS) error {
	return s.Repo.Create(data)
}
func (s *GPSService) GetAll(page, pageSize int) ([]models.GPS, error) {
	return s.Repo.GetAll(page, pageSize)
}
