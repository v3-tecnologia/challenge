package usecases

import (
	"context"

	"github.com/ricardoraposo/challenge/internal/repository"
)

type GyroscopeUseCase interface {
	CreateGyroscopeReading(ctx context.Context, params repository.InsertGyroscopeReadingParams) (repository.GyroscopeReading, error)
}

type GyroscopeQuerier interface {
	GetDeviceByID(ctx context.Context, deviceID string) (repository.Device, error)
	InsertDevice(ctx context.Context, deviceID string) (repository.Device, error)
	InsertGyroscopeReading(ctx context.Context, params repository.InsertGyroscopeReadingParams) (repository.GyroscopeReading, error)
}

type gyroscopeUseCase struct {
	queries GyroscopeQuerier
}

func NewGyroscopeUseCase(queries GyroscopeQuerier) GyroscopeUseCase {
	return &gyroscopeUseCase{
		queries: queries,
	}
}

func (uc *gyroscopeUseCase) CreateGyroscopeReading(ctx context.Context, params repository.InsertGyroscopeReadingParams) (repository.GyroscopeReading, error) {
	deviceUC := NewDeviceUseCase(uc.queries)
	_, err := deviceUC.CreateDevice(ctx, params.DeviceID)

	if err != nil {
		return repository.GyroscopeReading{}, err
	}

	return uc.queries.InsertGyroscopeReading(ctx, params)
}
