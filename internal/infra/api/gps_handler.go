package api

import (
	usecase "github.com/iamrosada0/v3/internal/usecase/gps"
)

type GPSHandlers struct {
	CreateGPSUseCase *usecase.CreateGPSUseCase
}
