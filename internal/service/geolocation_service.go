package service

import (
	"fmt"

	"github.com/wellmtx/challenge/internal/dtos"
	"github.com/wellmtx/challenge/internal/infra/entities"
	"github.com/wellmtx/challenge/internal/infra/repositories"
)

type GeolocationService struct {
	geolocationRepository repositories.GeolocationRepository
}

func NewGeolocationService(geolocationRepository repositories.GeolocationRepository) *GeolocationService {
	return &GeolocationService{
		geolocationRepository: geolocationRepository,
	}
}

func (g *GeolocationService) SaveData(data dtos.GeolocationDataDto) (entities.GeolocationEntity, error) {
	var geolocationEntity entities.GeolocationEntity

	geolocationEntity.Latitude = data.Latitude
	geolocationEntity.Longitude = data.Longitude

	result, err := g.geolocationRepository.Create(geolocationEntity)
	if err != nil {
		return entities.GeolocationEntity{}, fmt.Errorf("failed to save data: %w", err)
	}

	return result, nil
}
