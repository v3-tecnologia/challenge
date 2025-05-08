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

type MockGyroscopeService struct {
	mock.Mock
}

func (m *MockGyroscopeService) Create(data *d.GyroscopeData) error {
	args := m.Called(data)
	return args.Error(0)
}

func TestGyroscopeUseCase_Create_Success(t *testing.T) {
	mockService := new(MockGyroscopeService)
	uc := uc.NewGyroscopeUseCase(mockService)

	deviceID := "00:1A:2B:3C:4D:5E"
	x := gofakeit.Float64()
	y := gofakeit.Float64()
	z := gofakeit.Float64()
	timestamp := time.Now()

	data, _ := d.NewGyroscopeData(deviceID, x, y, z, timestamp)
	mockService.On("Create", data).Return(nil)

	_, err := uc.Create(deviceID, x, y, z, timestamp)
	assert.NoError(t, err)
	mockService.AssertExpectations(t)
}

func TestGyroscopeUseCase_Create_InvalidDeviceID(t *testing.T) {
	mockService := new(MockGyroscopeService)
	uc := gyroscope.NewGyroscopeUseCase(mockService)

	_, err := uc.Create("invalid-mac", 1.0, 2.0, 3.0, time.Now())
	assert.Error(t, err)
	assert.Equal(t, "invalid MAC address", err.Error())
}
