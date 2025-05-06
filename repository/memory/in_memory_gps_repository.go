package memory_repository

import (
	"context"
	"sync"

	"github.com/mkafonso/go-cloud-challenge/entity"
)

type InMemoryGPSRepository struct {
	sync.Mutex
	data map[string][]*entity.GPS
}

func NewInMemoryGPSRepository() *InMemoryGPSRepository {
	return &InMemoryGPSRepository{
		data: make(map[string][]*entity.GPS),
	}
}

func (r *InMemoryGPSRepository) SaveGPS(ctx context.Context, data *entity.GPS) error {
	r.Lock()
	defer r.Unlock()

	r.data[data.DeviceID] = append(r.data[data.DeviceID], data)
	return nil
}
