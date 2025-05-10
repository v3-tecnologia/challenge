package usecases

import (
	"context"
	"fmt"
	"io"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ricardoraposo/challenge/internal/interfaces"
	"github.com/ricardoraposo/challenge/internal/repository"
)

type CreatePhotoParams struct {
	DeviceID    string
	File        io.Reader
	Key         string
	CollectedAt pgtype.Timestamp
}

type PhotosUseCase interface {
	CreatePhoto(ctx context.Context, params CreatePhotoParams) (repository.Photo, error)
}

type PhotosQueries interface {
	GetDeviceByID(ctx context.Context, deviceID string) (repository.Device, error)
	InsertDevice(ctx context.Context, deviceID string) (repository.Device, error)
	InsertPhoto(ctx context.Context, arg repository.InsertPhotoParams) (repository.Photo, error)
}

type photosUseCase struct {
	queries      PhotosQueries
	uploader     interfaces.BucketUploader
	faceDetector interfaces.FaceDetector
}

func NewPhotosUseCase(
	queries PhotosQueries,
	uploader interfaces.BucketUploader,
	faceDetector interfaces.FaceDetector,
) PhotosUseCase {
	return &photosUseCase{
		queries:      queries,
		uploader:     uploader,
		faceDetector: faceDetector,
	}
}

func (uc *photosUseCase) CreatePhoto(ctx context.Context, params CreatePhotoParams) (repository.Photo, error) {
	cancelCtx, cancelCtxFunc := context.WithCancel(ctx)
	defer cancelCtxFunc()

	uploadCh := make(chan string)
	uploadErrCh := make(chan error, 1)

	go uc.uploader.UploadAsync(cancelCtx, params.File, params.Key, uploadCh, uploadErrCh)

	deviceUC := NewDeviceUseCase(uc.queries)
	device, err := deviceUC.CreateDevice(ctx, params.DeviceID)
	if err != nil {
		return repository.Photo{}, err
	}

	var imageURL string
	select {
	case url := <-uploadCh:
		imageURL = url
	case uploadErr := <-uploadErrCh:
		return repository.Photo{}, fmt.Errorf("upload failed: %w", uploadErr)
	case <-ctx.Done():
		return repository.Photo{}, ctx.Err()
	}

	faceMatches, err := uc.faceDetector.HandleFaceRecognition(cancelCtx, params.Key)
	if err != nil {
		return repository.Photo{}, err
	}

	recurrentFace := false
	if len(*faceMatches) > 0 {
		recurrentFace = true
	}

	insertPhotoParams := repository.InsertPhotoParams{
		DeviceID:      device.DeviceID,
		ImageUrl:      imageURL,
		RecurrentUser: pgtype.Bool{Bool: recurrentFace, Valid: true},
		CollectedAt:   params.CollectedAt,
	}

	return uc.queries.InsertPhoto(ctx, insertPhotoParams)
}
