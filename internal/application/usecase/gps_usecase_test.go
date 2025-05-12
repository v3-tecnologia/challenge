package usecase

import (
	"challenge-v3-backend/internal/domain/entity"
	"challenge-v3-backend/internal/interface/dto"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockGPSRepository struct {
	mock.Mock
}

func (m *MockGPSRepository) CreateGPSTelemetry(ctx context.Context, entity *entity.GPSTelemetry) error {
	args := m.Called(ctx, entity)
	return args.Error(0)
}

func TestCreateGPSUseCase(t *testing.T) {
	mockRepo := new(MockGPSRepository)
	gpsUseCase := NewGPSUseCase(mockRepo)

	gpsRequest := dto.CreateGPSRequestDTO{
		Latitude:  -23.550520,
		Longitude: -46.633308,
		DeviceId:  "00:11:22:33:44:55",
		CreatedAt: time.Now(),
	}

	mockRepo.On("CreateGPSTelemetry",
		mock.Anything,
		mock.MatchedBy(func(entity *entity.GPSTelemetry) bool {
			return entity.DeviceID == gpsRequest.DeviceId
		})).Return(nil)

	err := gpsUseCase.Create(context.Background(), gpsRequest)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
