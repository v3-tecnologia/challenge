package memory_repository

import (
	"context"
	"sync"

	"github.com/mkafonso/go-cloud-challenge/entity"
)

type InMemoryGyroscopeRepository struct {
	sync.Mutex
	data map[string][]*entity.Gyroscope
}

func NewInMemoryGyroscopeRepository() *InMemoryGyroscopeRepository {
	return &InMemoryGyroscopeRepository{
		data: make(map[string][]*entity.Gyroscope),
	}
}

func (r *InMemoryGyroscopeRepository) SaveGyroscope(ctx context.Context, data *entity.Gyroscope) error {
	r.Lock()
	defer r.Unlock()

	r.data[data.DeviceID] = append(r.data[data.DeviceID], data)
	return nil
}
