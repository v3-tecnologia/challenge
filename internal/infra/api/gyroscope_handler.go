package api

import (
	usecase "github.com/iamrosada0/v3/internal/usecase/gyroscope"
)

type GyroscopeHandlers struct {
	CreateGyroscopeUseCase *usecase.CreateGyroscopeUseCase
}

func NewGyroscopeHandlers(createGyroscopeUseCase *usecase.CreateGyroscopeUseCase) *GyroscopeHandlers {
	return &GyroscopeHandlers{
		CreateGyroscopeUseCase: createGyroscopeUseCase,
	}
}
