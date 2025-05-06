package dtos

import "github.com/wellmtx/challenge/internal/infra/entities"

type SavePhotoDTO struct {
	BaseDTO
	FilePath string `json:"file_path"`
	Image    []byte `json:"image"`
}

type SavePhotoResponseDTO struct {
	entities.PhotoEntity
	Matched bool `json:"matched"`
}
