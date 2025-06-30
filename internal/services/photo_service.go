package service

import (
	"challenge-cloud/internal/models"
	repository "challenge-cloud/internal/repositories/contracts"
)

type PhotoService struct {
	Repo repository.PhotoRepositoryInterface
}
type PhotoServiceInterface interface {
	Save(data *models.Photo) error
	GetAll(page, size int) ([]models.Photo, error)
}

func NewPhotoService(r repository.PhotoRepositoryInterface) *PhotoService {
	return &PhotoService{Repo: r}
}

func (s *PhotoService) Save(data *models.Photo) error {
	return s.Repo.Create(data)
}
func (s *PhotoService) GetAll(page, pageSize int) ([]models.Photo, error) {
	return s.Repo.GetAll(page, pageSize)
}
