package service

import (
	"challenge-cloud/internal/models"
	repository "challenge-cloud/internal/repositories"
)

type GyroscopeService struct {
	Repo *repository.GyroscopeRepository
}

func NewGyroscopeService(r *repository.GyroscopeRepository) *GyroscopeService {
	return &GyroscopeService{Repo: r}
}

func (s *GyroscopeService) Save(data *models.Gyroscope) error {
	return s.Repo.Create(data)
}
func (s *GyroscopeService) GetAll(page, pageSize int) ([]models.Gyroscope, error) {
	return s.Repo.GetAll(page, pageSize)
}
