package service_test

import (
	"challenge-cloud/internal/models"
	service "challenge-cloud/internal/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPhotoRepo struct {
	mock.Mock
}

func (m *MockPhotoRepo) Create(data *models.Photo) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockPhotoRepo) GetAll(page, size int) ([]models.Photo, error) {
	args := m.Called(page, size)
	return args.Get(0).([]models.Photo), args.Error(1)
}

func TestPhotoService_Save(t *testing.T) {
	mockRepo := new(MockPhotoRepo)
	service := service.NewPhotoService(mockRepo)

	data := &models.Photo{
		ImageURL:  "url",
		Timestamp: time.Now(),
		MAC:       "AA:BB:CC:DD:EE:FF",
	}

	mockRepo.On("Create", data).Return(nil)

	err := service.Save(data)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
