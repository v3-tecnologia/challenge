package service_test

import (
	"challenge-cloud/internal/models"
	service "challenge-cloud/internal/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGPSRepo struct {
	mock.Mock
}

func (m *MockGPSRepo) Create(data *models.GPS) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockGPSRepo) GetAll(page, size int) ([]models.GPS, error) {
	args := m.Called(page, size)
	return args.Get(0).([]models.GPS), args.Error(1)
}

func TestGPSService_Save(t *testing.T) {
	mockRepo := new(MockGPSRepo)
	service := service.NewGPSService(mockRepo)

	// timestamp, err := time.Parse(time.RFC3339, "2025-06-30T12:00:00Z")
	// assert.Nil(t, err)
	data := &models.GPS{
		Longitude: 1.1,
		Latitude:  2.2,
		Timestamp: time.Now(),
		MAC:       "abc123",
	}

	mockRepo.On("Create", data).Return(nil)

	err := service.Save(data)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
