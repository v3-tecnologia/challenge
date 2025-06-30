package service_test

import (
	"challenge-cloud/internal/models"
	service "challenge-cloud/internal/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGyroscopeRepo struct {
	mock.Mock
}

func (m *MockGyroscopeRepo) Create(data *models.Gyroscope) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockGyroscopeRepo) GetAll(page, size int) ([]models.Gyroscope, error) {
	args := m.Called(page, size)
	return args.Get(0).([]models.Gyroscope), args.Error(1)
}

func TestGyroscopeService_Save(t *testing.T) {
	mockRepo := new(MockGyroscopeRepo)
	service := service.NewGyroscopeService(mockRepo)

	// timestamp, err := time.Parse(time.RFC3339, "2025-06-30T12:00:00Z")
	// assert.Nil(t, err)
	data := &models.Gyroscope{
		X:         1.1,
		Y:         2.2,
		Z:         3.3,
		Timestamp: time.Now(),
		MAC:       "abc123",
	}

	mockRepo.On("Create", data).Return(nil)

	err := service.Save(data)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
