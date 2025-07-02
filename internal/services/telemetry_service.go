package services

import (
	"go-challenge/internal/models"
	"go-challenge/internal/repository"
)

type TelemetryService struct {
	Repo repository.TelemetryRepository
}

func NewTelemetryService(repo repository.TelemetryRepository) *TelemetryService {
	return &TelemetryService{Repo: repo}
}

func (s *TelemetryService) SaveGyroscopeData(data models.Gyroscope) error {
	// Validação adicional (se necessário)
	return s.Repo.SaveGyroscopeData(data)
}

func (s *TelemetryService) SaveGPSData(data models.GPS) error {
	// Validação adicional (se necessário)
	return s.Repo.SaveGPSData(data)
}

func (s *TelemetryService) SavePhotoData(data models.Photo) error {
	// Validação adicional (se necessário)
	return s.Repo.SavePhotoData(data)
}
