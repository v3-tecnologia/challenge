package api

import (
	usecase "github.com/iamrosada0/v3/internal/usecase/photos"
)

type PhotoHandlers struct {
	CreatePhotoUseCase *usecase.CreatePhotoUseCase
}

func NewPhotoHandlers(createPhotoUseCase *usecase.CreatePhotoUseCase) *PhotoHandlers {
	return &PhotoHandlers{
		CreatePhotoUseCase: createPhotoUseCase,
	}
}
