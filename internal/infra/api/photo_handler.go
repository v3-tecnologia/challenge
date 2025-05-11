package api

import (
	usecase "github.com/iamrosada0/v3/internal/usecase/photos"
)

type PhotoHandlers struct {
	CreatePhotoUseCase *usecase.CreatePhotoUseCase
}
