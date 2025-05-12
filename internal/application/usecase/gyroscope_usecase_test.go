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

type MockGyroscopeRepository struct {
	mock.Mock
}

func (m *MockGyroscopeRepository) CreateGyroscopeTelemetry(ctx context.Context, entity *entity.GyroscopeTelemetry) error {
	args := m.Called(ctx, entity)
	return args.Error(0)
}

func TestSaveGyroscopeService(t *testing.T) {
	mockRepo := new(MockGyroscopeRepository)
	gyroscopeUseCase := NewGyroscopeUseCase(mockRepo)

	gyroscopeRequest := dto.CreateGyroscopeRequestDTO{
		X:         1,
		Y:         2,
		Z:         3,
		DeviceId:  "00:11:22:33:44:55",
		CreatedAt: time.Now(),
	}

	mockRepo.On("CreateGyroscopeTelemetry",
		mock.Anything,
		mock.MatchedBy(func(entity *entity.GyroscopeTelemetry) bool {
			return entity.DeviceID == gyroscopeRequest.DeviceId &&
				entity.X == gyroscopeRequest.X &&
				entity.Y == gyroscopeRequest.Y &&
				entity.Z == gyroscopeRequest.Z
		})).Return(nil)

	err := gyroscopeUseCase.Create(context.Background(), gyroscopeRequest)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSaveGyroscopeService_ValidationError(t *testing.T) {
	mockRepo := new(MockGyroscopeRepository)
	gyroscopeUseCase := NewGyroscopeUseCase(mockRepo)

	gyroscopeRequest := dto.CreateGyroscopeRequestDTO{
		DeviceId:  "",
		CreatedAt: time.Now(),
		X:         1.0,
		Y:         2.0,
		Z:         3.0,
	}

	mockRepo.On("CreateGyroscopeTelemetry", mock.Anything, mock.Anything).Return(nil)

	err := gyroscopeUseCase.Create(context.Background(), gyroscopeRequest)

	assert.Error(t, err)
	mockRepo.AssertNotCalled(t, "CreateGyroscopeTelemetry")

}
