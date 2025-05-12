package usecase

import (
	"challenge-v3-backend/internal/domain/entity"
	"challenge-v3-backend/internal/domain/gateway"
	"challenge-v3-backend/internal/interface/dto"
	"context"
)

type GPSUseCase interface {
	Create(ctx context.Context, input dto.CreateGPSRequestDTO) error
}

type GPSUseCaseImpl struct {
	gateway gateway.GPSTelemetryGateway
}

func NewGPSUseCase(gateway gateway.GPSTelemetryGateway) *GPSUseCaseImpl {
	return &GPSUseCaseImpl{
		gateway: gateway,
	}
}

func (uc *GPSUseCaseImpl) Create(ctx context.Context, input dto.CreateGPSRequestDTO) error {
	entityCreated := entity.BuildGPSTelemetry(input.DeviceId, input.CreatedAt, input.Latitude, input.Longitude)

	return uc.gateway.CreateGPSTelemetry(ctx, entityCreated)
}
