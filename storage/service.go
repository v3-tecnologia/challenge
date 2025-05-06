package storage

import "context"

type PhotoStorageService interface {
	Store(ctx context.Context, deviceID string, fileBytes []byte) (string, error)
}
