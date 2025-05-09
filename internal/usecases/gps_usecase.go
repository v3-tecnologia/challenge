package usecases

import (
	"context"

	"github.com/ricardoraposo/challenge/internal/repository"
)

type GpsUseCase interface {
	CreateGPSReading(ctx context.Context, params repository.InsertGPSReadingParams) (repository.GpsReading, error)
}

type GpsQuerier interface {
	GetDeviceByID(ctx context.Context, deviceID string) (repository.Device, error)
	InsertDevice(ctx context.Context, deviceID string) (repository.Device, error)
	InsertGPSReading(ctx context.Context, params repository.InsertGPSReadingParams) (repository.GpsReading, error)
}

type gpsUseCase struct {
	queries GpsQuerier
}

func NewGPSUseCase(queries GpsQuerier) GpsUseCase {
	return &gpsUseCase{
		queries: queries,
	}
}

func (uc *gpsUseCase) CreateGPSReading(ctx context.Context, params repository.InsertGPSReadingParams) (repository.GpsReading, error) {
	deviceUC := NewDeviceUseCase(uc.queries)
	_, err := deviceUC.CreateDevice(ctx, params.DeviceID)

	if err != nil {
		return repository.GpsReading{}, err
	}

	return uc.queries.InsertGPSReading(ctx, params)
}
