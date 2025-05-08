package usecases

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ricardoraposo/challenge/internal/repository"
)

type DeviceUseCase interface {
	CreateDevice(ctx context.Context, deviceID string) (repository.Device, error)
}

type DeviceQuerier interface {
	GetDeviceByID(ctx context.Context, deviceID string) (repository.Device, error)
	InsertDevice(ctx context.Context, deviceID string) (repository.Device, error)
}

type deviceUseCaseImpl struct {
	queries DeviceQuerier
}

func NewDeviceUseCase(queries DeviceQuerier) DeviceUseCase {
	return &deviceUseCaseImpl{
		queries: queries,
	}
}

func (uc *deviceUseCaseImpl) CreateDevice(ctx context.Context, deviceID string) (repository.Device, error) {
	device, err := uc.queries.GetDeviceByID(ctx, deviceID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			device, err = uc.queries.InsertDevice(ctx, deviceID)

			if err != nil {
				return repository.Device{}, err
			}
		}
	}

	return device, nil
}
