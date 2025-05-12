package usecase

import (
	"challenge-v3-backend/internal/domain/entity"
	"challenge-v3-backend/internal/domain/gateway"
	"challenge-v3-backend/internal/interface/dto"
	"context"
)

type GyroscopeUseCase interface {
	Create(ctx context.Context, input dto.CreateGyroscopeRequestDTO) error
}

type GyroscopeUseCaseImpl struct {
	gateway gateway.GyroscopeTelemetryGateway
}

func NewGyroscopeUseCase(gateway gateway.GyroscopeTelemetryGateway) *GyroscopeUseCaseImpl {
	return &GyroscopeUseCaseImpl{
		gateway: gateway,
	}
}

func (uc *GyroscopeUseCaseImpl) Create(ctx context.Context, input dto.CreateGyroscopeRequestDTO) error {
	entityCreated := entity.BuildGyroscopeTelemetry(input.DeviceId, input.CreatedAt, input.X, input.Y, input.Z)

	if err := entityCreated.Validate(); err != nil {
		return err
	}

	return uc.gateway.CreateGyroscopeTelemetry(ctx, entityCreated)
}
