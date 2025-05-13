package usecases

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ricardoraposo/challenge/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (m *MockQueries) InsertGPSReading(ctx context.Context, arg repository.InsertGPSReadingParams) (repository.GpsReading, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repository.GpsReading), args.Error(1)
}

var (
	insertGpsParams = repository.InsertGPSReadingParams{
		DeviceID:    deviceID,
		Latitude:    pgtype.Numeric{Exp: 1, Valid: true},
		Longitude:   pgtype.Numeric{Exp: 2, Valid: true},
		CollectedAt: now,
	}

	expectedGpsReading = repository.GpsReading{
		DeviceID:    deviceID,
		Latitude:    pgtype.Numeric{Exp: 1, Valid: true},
		Longitude:   pgtype.Numeric{Exp: 2, Valid: true},
		CollectedAt: now,
	}
)

func Test_CreateGpsReading_WhenDeviceExists_ShouldNotInsertDevice(t *testing.T) {
	t.Parallel()

	mockQueries := new(MockQueries)
	uc := NewGPSUseCase(mockQueries)

	mockQueries.On("GetDeviceByID", mock.Anything, deviceID).Return(expectedDevice, nil)
	mockQueries.On("InsertGPSReading", mock.Anything, insertGpsParams).Return(expectedGpsReading, nil)

	reading, err := uc.CreateGPSReading(context.Background(), insertGpsParams)

	assert.NoError(t, err)
	assert.Equal(t, expectedGpsReading, reading)

	mockQueries.AssertExpectations(t)
}

func Test_CreateGpsReading_WhenDeviceMissing_ShouldInsertDeviceFirst(t *testing.T) {
	t.Parallel()

	mockQueries := new(MockQueries)
	uc := NewGPSUseCase(mockQueries)

	deviceID := "test-device"

	mockQueries.On("GetDeviceByID", mock.Anything, deviceID).Return(repository.Device{}, sql.ErrNoRows)
	mockQueries.On("InsertDevice", mock.Anything, deviceID).Return(expectedDevice, nil)
	mockQueries.On("InsertGPSReading", mock.Anything, insertGpsParams).Return(expectedGpsReading, nil)

	reading, err := uc.CreateGPSReading(context.Background(), insertGpsParams)

	assert.NoError(t, err)
	assert.Equal(t, expectedGpsReading, reading)

	mockQueries.AssertExpectations(t)
}
