package memory_repository

import (
	"context"
	"sync"

	"github.com/mkafonso/go-cloud-challenge/entity"
)

type InMemoryPhotoRepository struct {
	sync.Mutex
	data map[string][]*entity.Photo
}

func NewInMemoryPhotoRepository() *InMemoryPhotoRepository {
	return &InMemoryPhotoRepository{
		data: make(map[string][]*entity.Photo),
	}
}

func (r *InMemoryPhotoRepository) SavePhoto(ctx context.Context, data *entity.Photo) error {
	r.Lock()
	defer r.Unlock()

	r.data[data.DeviceID] = append(r.data[data.DeviceID], data)
	return nil
}
