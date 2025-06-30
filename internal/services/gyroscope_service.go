package service

import (
	"challenge-cloud/internal/models"
	repository "challenge-cloud/internal/repositories/contracts"
)

type GyroscopeService struct {
	Repo repository.GyroscopeRepositoryInterface
}

type GyroscopeServiceInterface interface {
	Save(data *models.Gyroscope) error
	GetAll(page, size int) ([]models.Gyroscope, error)
}

func NewGyroscopeService(repo repository.GyroscopeRepositoryInterface) *GyroscopeService {
	return &GyroscopeService{Repo: repo}
}

func (s *GyroscopeService) Save(data *models.Gyroscope) error {
	return s.Repo.Create(data)
}
func (s *GyroscopeService) GetAll(page, pageSize int) ([]models.Gyroscope, error) {
	return s.Repo.GetAll(page, pageSize)
}
