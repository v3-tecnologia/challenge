package usecase

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	d "github.com/iamrosada0/v3/internal/domain/gyroscope"
	"github.com/iamrosada0/v3/internal/usecase/gyroscope"
	uc "github.com/iamrosada0/v3/internal/usecase/gyroscope"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGyroscopeRepository struct {
	mock.Mock
}

func (m *MockGyroscopeRepository) Create(data *d.GyroscopeData) (*d.GyroscopeData, error) {
	args := m.Called(data)
	return args.Get(0).(*d.GyroscopeData), args.Error(1)
}

func TestGyroscopeUseCase_Create_Success(t *testing.T) {
	mockRepo := new(MockGyroscopeRepository)
	uc := uc.NewGyroscopeUseCase(mockRepo)

	deviceID := "00:1A:2B:3C:4D:5E"
	x := gofakeit.Float64()
	y := gofakeit.Float64()
	z := gofakeit.Float64()
	timestamp := time.Now()

	data, _ := d.NewGyroscopeData(deviceID, x, y, z, timestamp)
	mockRepo.On("Create", data).Return(data, nil)

	result, err := uc.Create(deviceID, x, y, z, timestamp)
	assert.NoError(t, err)
	assert.Equal(t, data, result)
	mockRepo.AssertExpectations(t)
}

func TestGyroscopeUseCase_Create_InvalidDeviceID(t *testing.T) {
	mockService := new(MockGyroscopeService)
	uc := gyroscope.NewGyroscopeUseCase(mockService)

	_, err := uc.Create("invalid-mac", 1.0, 2.0, 3.0, time.Now())
	assert.Error(t, err)
	assert.Equal(t, "invalid MAC address", err.Error())
}
