package service

import (
	"fmt"

	"github.com/wellmtx/challenge/internal/dtos"
	"github.com/wellmtx/challenge/internal/infra/entities"
	"github.com/wellmtx/challenge/internal/infra/repositories"
)

type GyroscopeService struct {
	gyroscopeRepository repositories.GyroscopeRepository
}

func NewGyroscopeService(gyroscopeRepository repositories.GyroscopeRepository) *GyroscopeService {
	return &GyroscopeService{
		gyroscopeRepository: gyroscopeRepository,
	}
}

func (g *GyroscopeService) SaveData(data dtos.GyroscopeDataDto) (entities.GyroscopeEntity, error) {
	var gyroscopeEntity entities.GyroscopeEntity

	gyroscopeEntity.XValue = data.X
	gyroscopeEntity.YValue = data.Y
	gyroscopeEntity.ZValue = data.Z
	gyroscopeEntity.MacAddress = data.MacAddress

	result, err := g.gyroscopeRepository.Create(gyroscopeEntity)
	if err != nil {
		return entities.GyroscopeEntity{}, fmt.Errorf("failed to save data: %w", err)
	}

	return result, nil
}
