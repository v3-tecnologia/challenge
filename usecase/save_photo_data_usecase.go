package usecase

import (
	"context"

	"github.com/mkafonso/go-cloud-challenge/entity"
	"github.com/mkafonso/go-cloud-challenge/recognition"
	"github.com/mkafonso/go-cloud-challenge/repository"
	"github.com/mkafonso/go-cloud-challenge/utils"
)

type SavePhotoRequest struct {
	DeviceID  string
	FilePath  string
	Timestamp string
}

type SavePhotoResponse struct {
	Recognized bool
}

type SavePhoto struct {
	repo       repository.PhotoRepositoryInterface
	recognizer recognition.FaceRecognitionService
}

func NewSavePhoto(
	repo repository.PhotoRepositoryInterface,
	recognizer recognition.FaceRecognitionService,
) *SavePhoto {
	return &SavePhoto{
		repo:       repo,
		recognizer: recognizer,
	}
}

func (uc *SavePhoto) Execute(ctx context.Context, data *SavePhotoRequest) (*SavePhotoResponse, error) {
	timestamp, err := utils.ParseRFC3339(data.Timestamp)
	if err != nil {
		return nil, err
	}

	// ⚠️ Blocking request until Rekognition finishes
	// ⚠️ Explanation: Required by challenge specification
	// ⚠️ More details can be found here: `docs/documentacao-tecnica.md`
	recognized, err := uc.recognizer.CompareWithHistory(ctx, data.FilePath, data.DeviceID)
	if err != nil {
		return nil, err
	}

	photo, err := entity.NewPhoto(data.DeviceID, data.FilePath, timestamp, recognized)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.SavePhoto(ctx, photo); err != nil {
		return nil, err
	}

	return &SavePhotoResponse{Recognized: recognized}, nil
}
