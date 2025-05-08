package usecases

import (
	"context"

	"github.com/ricardoraposo/challenge/internal/repository"
)

type GyroscopeUseCase interface {
	CreateGyroscopeReading(ctx context.Context, params repository.InsertGyroscopeReadingParams) (repository.GyroscopeReading, error)
}

type gyroscopeUseCaseImpl struct {
	queries *repository.Queries
}

func NewGyroscopeUseCase(queries *repository.Queries) GyroscopeUseCase {
	return &gyroscopeUseCaseImpl{
		queries: queries,
	}
}

func (uc *gyroscopeUseCaseImpl) CreateGyroscopeReading(ctx context.Context, params repository.InsertGyroscopeReadingParams) (repository.GyroscopeReading, error) {
	return uc.queries.InsertGyroscopeReading(ctx, params)
}
