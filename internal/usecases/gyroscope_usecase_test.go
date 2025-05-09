package usecases

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ricardoraposo/challenge/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQueries struct {
	mock.Mock
}

func (m *MockQueries) InsertDevice(ctx context.Context, deviceID string) (repository.Device, error) {
	args := m.Called(ctx, deviceID)
	return args.Get(0).(repository.Device), args.Error(1)
}

func (m *MockQueries) InsertGyroscopeReading(ctx context.Context, arg repository.InsertGyroscopeReadingParams) (repository.GyroscopeReading, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repository.GyroscopeReading), args.Error(1)
}

func (m *MockQueries) GetDeviceByID(ctx context.Context, arg string) (repository.Device, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repository.Device), args.Error(1)
}

var (
	deviceID = "test-device"
	now      = pgtype.Timestamp{Time: time.Now(), Valid: true}
	insertGyroscopeParams   = repository.InsertGyroscopeReadingParams{
		DeviceID:    deviceID,
		X:           1.0,
		Y:           2.0,
		Z:           3.0,
		CollectedAt: now,
	}

	expectedDevice = repository.Device{
		DeviceID:     deviceID,
		RegisteredAt: now,
	}

	expectedReading = repository.GyroscopeReading{
		DeviceID:    deviceID,
		X:           1.0,
		Y:           2.0,
		Z:           3.0,
		CollectedAt: now,
	}
)

func Test_CreateReading_WhenDeviceExists_ShouldNotInsertDevice(t *testing.T) {
	t.Parallel()

	mockQueries := new(MockQueries)
	uc := NewGyroscopeUseCase(mockQueries)

	mockQueries.On("GetDeviceByID", mock.Anything, deviceID).Return(expectedDevice, nil)
	mockQueries.On("InsertGyroscopeReading", mock.Anything, insertGyroscopeParams).Return(expectedReading, nil)

	reading, err := uc.CreateGyroscopeReading(context.Background(), insertGyroscopeParams)

	assert.NoError(t, err)
	assert.Equal(t, expectedReading, reading)

	mockQueries.AssertExpectations(t)
}

func Test_CreateReading_WhenDeviceMissing_ShouldInsertDeviceFirst(t *testing.T) {
	t.Parallel()

	mockQueries := new(MockQueries)
	uc := NewGyroscopeUseCase(mockQueries)

	deviceID := "test-device"

	mockQueries.On("GetDeviceByID", mock.Anything, deviceID).Return(repository.Device{}, sql.ErrNoRows)
	mockQueries.On("InsertDevice", mock.Anything, deviceID).Return(expectedDevice, nil)
	mockQueries.On("InsertGyroscopeReading", mock.Anything, insertGyroscopeParams).Return(expectedReading, nil)

	reading, err := uc.CreateGyroscopeReading(context.Background(), insertGyroscopeParams)

	assert.NoError(t, err)
	assert.Equal(t, expectedReading, reading)

	mockQueries.AssertExpectations(t)
}
