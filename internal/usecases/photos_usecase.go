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
	queries  PhotosQueries
	uploader interfaces.BucketUploader
}

func NewPhotosUseCase(queries PhotosQueries, uploader interfaces.BucketUploader) PhotosUseCase {
	return &photosUseCase{
		queries:  queries,
		uploader: uploader,
	}
}

func (uc *photosUseCase) CreatePhoto(ctx context.Context, params CreatePhotoParams) (repository.Photo, error) {
	uploadCtx, cancelUpload := context.WithCancel(ctx)
	defer cancelUpload()

	ch := make(chan string)
	errCh := make(chan error, 1)

	go uc.uploader.UploadAsync(uploadCtx, params.File, params.Key, ch, errCh)

	deviceUC := NewDeviceUseCase(uc.queries)
	device, err := deviceUC.CreateDevice(ctx, params.DeviceID)
	if err != nil {
		return repository.Photo{}, err
	}

	var imageURL string
	select {
	case url := <-ch:
		imageURL = url
	case uploadErr := <-errCh:
		return repository.Photo{}, fmt.Errorf("upload failed: %w", uploadErr)
	case <-ctx.Done():
		return repository.Photo{}, ctx.Err()
	}

	insertPhotoParams := repository.InsertPhotoParams{
		DeviceID:    device.DeviceID,
		ImageUrl:    imageURL,
		CollectedAt: params.CollectedAt,
	}

	return uc.queries.InsertPhoto(ctx, insertPhotoParams)
}
