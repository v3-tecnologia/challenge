package service

import (
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	d "github.com/iamrosada0/v3/internal/domain/gyroscope"
	"github.com/iamrosada0/v3/internal/service/gyroscope"

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
func TestGyroscopeService_Create_Success(t *testing.T) {
	mockRepo := new(MockGyroscopeRepository)
	service := gyroscope.NewGyroscopeService(mockRepo)

	deviceID := "00:1A:2B:3C:4D:5E"
	x := gofakeit.Float64()
	y := gofakeit.Float64()
	z := gofakeit.Float64()
	timestamp := time.Now()

	data, _ := d.NewGyroscopeData(deviceID, x, y, z, timestamp)
	mockRepo.On("Create", data).Return(nil)

	result, err := service.Create(data)
	assert.NoError(t, err)
	assert.Equal(t, data, result)
	mockRepo.AssertExpectations(t)
}
func TestGyroscopeService_Create_RepositoryError(t *testing.T) {
	mockRepo := new(MockGyroscopeRepository)
	service := gyroscope.NewGyroscopeService(mockRepo)

	deviceID := "00:1A:2B:3C:4D:5E"
	x := gofakeit.Float64()
	y := gofakeit.Float64()
	z := gofakeit.Float64()
	timestamp := time.Now()

	data, _ := d.NewGyroscopeData(deviceID, x, y, z, timestamp)
	mockRepo.On("Create", data).Return(errors.New("database error"))

	_, err := service.Create(data)
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
}
