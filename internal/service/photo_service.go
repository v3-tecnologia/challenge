package service

import (
	"fmt"
	"os"

	"github.com/wellmtx/challenge/internal/dtos"
	"github.com/wellmtx/challenge/internal/infra/entities"
	"github.com/wellmtx/challenge/internal/infra/providers"
	"github.com/wellmtx/challenge/internal/infra/repositories"
)

type PhotoService struct {
	photoRepository     repositories.PhotoRepository
	recognitionProvider providers.RecognitionProvider
}

func NewPhotoService(photoRepository repositories.PhotoRepository, recognitionProvider providers.RecognitionProvider) *PhotoService {
	return &PhotoService{
		photoRepository:     photoRepository,
		recognitionProvider: recognitionProvider,
	}
}

func (s *PhotoService) RecognizePhoto(dto dtos.SavePhotoDTO) (dtos.SavePhotoResponseDTO, error) {
	photoEntity := entities.PhotoEntity{
		FilePath: dto.FilePath,
	}
	photoEntity.MacAddress = dto.MacAddress
	photoEntity.Timestamp = dto.GetTimestamp()

	photoEntity, err := s.photoRepository.Create(photoEntity)
	if err != nil {
		return dtos.SavePhotoResponseDTO{}, err
	}

	photos, err := s.photoRepository.ListByMacAddress(dto.MacAddress)
	if err != nil {
		return dtos.SavePhotoResponseDTO{}, fmt.Errorf("unable to get photos, %v", err)
	}

	matched := false
	for _, photo := range photos {
		if photo.ID == photoEntity.ID {
			continue
		}

		photoBytes, err := os.ReadFile(photo.FilePath)
		if err != nil {
			return dtos.SavePhotoResponseDTO{}, fmt.Errorf("unable to read photo file, %v", err)
		}

		matched, err = s.recognitionProvider.CompareFaces(photoBytes, dto.Image)
		if err != nil {
			return dtos.SavePhotoResponseDTO{}, fmt.Errorf("unable to compare faces, %v", err)
		}

		if matched {
			break
		}
	}

	return dtos.SavePhotoResponseDTO{
		PhotoEntity: photoEntity,
		Matched:     matched,
	}, nil
}
