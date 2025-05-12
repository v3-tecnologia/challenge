package usecase

import (
	"errors"
	"testing"
	"time"
	"v3/internal/domain"
	"v3/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do GPSRepository
type MockGPSRepository struct {
	mock.Mock
}

func (m *MockGPSRepository) Create(gpsData *domain.GPS) (*domain.GPS, error) {
	args := m.Called(gpsData)
	return args.Get(0).(*domain.GPS), args.Error(1)
}

func TestCreateGPSUseCase_Execute(t *testing.T) {
	t.Run("sucesso ao criar GPS", func(t *testing.T) {
		mockRepo := new(MockGPSRepository)
		useCase := usecase.NewCreateGPSUseCase(mockRepo)

		input := domain.GPSDto{
			DeviceID:  "device123",
			Timestamp: 1618481023,
			Latitude:  10.0,
			Longitude: 20.0,
		}

		expectedTime := time.Unix(input.Timestamp, 0)

		expectedGPS := &domain.GPS{
			ID:        "gps123",
			DeviceID:  input.DeviceID,
			Latitude:  input.Latitude,
			Longitude: input.Longitude,
			Timestamp: expectedTime,
			CreatedAt: time.Now(), // opcional para assert
		}

		mockRepo.On("Create", mock.MatchedBy(func(g *domain.GPS) bool {
			return g.DeviceID == input.DeviceID &&
				g.Latitude == input.Latitude &&
				g.Longitude == input.Longitude &&
				g.Timestamp.Equal(expectedTime)
		})).Return(expectedGPS, nil)

		result, err := useCase.Execute(input)

		assert.NoError(t, err)
		assert.Equal(t, expectedGPS, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("erro ao criar GPS", func(t *testing.T) {
		mockRepo := new(MockGPSRepository)
		useCase := usecase.NewCreateGPSUseCase(mockRepo)

		input := domain.GPSDto{
			DeviceID:  "device123",
			Timestamp: 1618481023,
			Latitude:  10.0,
			Longitude: 20.0,
		}

		mockRepo.On("Create", mock.Anything).Return(nil, errors.New("erro ao salvar GPS"))

		result, err := useCase.Execute(input)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}
