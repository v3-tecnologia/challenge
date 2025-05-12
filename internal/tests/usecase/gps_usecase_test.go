package usecase

import (
	"errors"
	"math"
	"testing"
	"time"
	"v3/internal/domain"
	"v3/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGPSRepository struct {
	mock.Mock
}

func (m *MockGPSRepository) Create(d *domain.GPS) (*domain.GPS, error) {
	args := m.Called(d)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.GPS), args.Error(1)
}

func TestCreateGPSUseCase_Execute(t *testing.T) {
	tests := []struct {
		name           string
		input          domain.GPSDto
		setupMock      func(repo *MockGPSRepository)
		wantErr        error
		validateResult func(t *testing.T, gps *domain.GPS)
	}{
		{
			name: "Successful GPS creation",
			input: domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			setupMock: func(repo *MockGPSRepository) {
				gps := &domain.GPS{
					ID:        "mock-id",
					DeviceID:  "00:0a:95:9d:68:16",
					Latitude:  40.7128,
					Longitude: -74.0060,
					Timestamp: time.Unix(time.Now().Unix(), 0).UTC(),
				}
				repo.On("Create", mock.AnythingOfType("*domain.GPS")).Return(gps, nil)
			},
			wantErr: nil,
			validateResult: func(t *testing.T, gps *domain.GPS) {
				assert.NotNil(t, gps)
				assert.Equal(t, "mock-id", gps.ID)
				assert.Equal(t, "00:0a:95:9d:68:16", gps.DeviceID)
				assert.Equal(t, 40.7128, gps.Latitude)
				assert.Equal(t, -74.0060, gps.Longitude)
				assert.WithinDuration(t, time.Now().UTC(), gps.Timestamp, time.Second)
			},
		},
		{
			name: "Invalid DeviceID",
			input: domain.GPSDto{
				DeviceID:  "invalid-mac",
				Timestamp: time.Now().Unix(),
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			setupMock: func(repo *MockGPSRepository) {
				// Repo.Create not called due to NewGPSData error
			},
			wantErr: domain.ErrDeviceIDGPS,
			validateResult: func(t *testing.T, gps *domain.GPS) {
				assert.Nil(t, gps)
			},
		},
		{
			name: "Zero Timestamp",
			input: domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: 0,
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			setupMock: func(repo *MockGPSRepository) {
				// Repo.Create not called due to NewGPSData error
			},
			wantErr: domain.ErrTimestampGPS,
			validateResult: func(t *testing.T, gps *domain.GPS) {
				assert.Nil(t, gps)
			},
		},
		{
			name: "Invalid Latitude (NaN)",
			input: domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				Latitude:  math.NaN(),
				Longitude: -74.0060,
			},
			setupMock: func(repo *MockGPSRepository) {
				// Repo.Create not called due to NewGPSData error
			},
			wantErr: domain.ErrInvalidGPSValues,
			validateResult: func(t *testing.T, gps *domain.GPS) {
				assert.Nil(t, gps)
			},
		},
		{
			name: "Repository error",
			input: domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			setupMock: func(repo *MockGPSRepository) {
				repo.On("Create", mock.AnythingOfType("*domain.GPS")).Return(nil, errors.New("database error"))
			},
			wantErr: domain.ErrSaveGPSData,
			validateResult: func(t *testing.T, gps *domain.GPS) {
				assert.Nil(t, gps)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := &MockGPSRepository{}
			tt.setupMock(repo)

			uc := usecase.NewCreateGPSUseCase(repo)

			result, err := uc.Execute(tt.input)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}

			tt.validateResult(t, result)

			repo.AssertExpectations(t)
		})
	}
}
