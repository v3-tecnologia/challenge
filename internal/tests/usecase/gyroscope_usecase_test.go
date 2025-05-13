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

type MockGyroscopeRepository struct {
	mock.Mock
}

func (m *MockGyroscopeRepository) Create(d *domain.Gyroscope) (*domain.Gyroscope, error) {
	args := m.Called(d)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Gyroscope), args.Error(1)
}

func TestCreateGyroscopeUseCase_Execute(t *testing.T) {
	tests := []struct {
		name           string
		input          domain.GyroscopeDto
		setupMock      func(repo *MockGyroscopeRepository)
		wantErr        error
		validateResult func(t *testing.T, gyro *domain.Gyroscope)
	}{
		{
			name: "Successful Gyroscope creation",
			input: domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				X:         1.5,
				Y:         -2.3,
				Z:         0.8,
			},
			setupMock: func(repo *MockGyroscopeRepository) {
				gyro := &domain.Gyroscope{
					ID:        "mock-id",
					DeviceID:  "00:0a:95:9d:68:16",
					X:         1.5,
					Y:         -2.3,
					Z:         0.8,
					Timestamp: time.Unix(time.Now().Unix(), 0).UTC(),
				}
				repo.On("Create", mock.AnythingOfType("*domain.Gyroscope")).Return(gyro, nil).Once()
			},
			wantErr: nil,
			validateResult: func(t *testing.T, gyro *domain.Gyroscope) {
				assert.NotNil(t, gyro)
				assert.Equal(t, "mock-id", gyro.ID)
				assert.Equal(t, "00:0a:95:9d:68:16", gyro.DeviceID)
				assert.Equal(t, 1.5, gyro.X)
				assert.Equal(t, -2.3, gyro.Y)
				assert.Equal(t, 0.8, gyro.Z)
				assert.WithinDuration(t, time.Now().UTC(), gyro.Timestamp, time.Second)
			},
		},
		{
			name: "Invalid DeviceID",
			input: domain.GyroscopeDto{
				DeviceID:  "invalid-mac",
				Timestamp: time.Now().Unix(),
				X:         1.5,
				Y:         -2.3,
				Z:         0.8,
			},
			setupMock: func(repo *MockGyroscopeRepository) {
				// Repo.Create should not be called
			},
			wantErr: domain.ErrDeviceIDGyroscope,
			validateResult: func(t *testing.T, gyro *domain.Gyroscope) {
				assert.Nil(t, gyro)
			},
		},
		{
			name: "Zero Timestamp",
			input: domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: 0,
				X:         1.5,
				Y:         -2.3,
				Z:         0.8,
			},
			setupMock: func(repo *MockGyroscopeRepository) {
				// Repo.Create should not be called
			},
			wantErr: domain.ErrTimestampGyroscope,
			validateResult: func(t *testing.T, gyro *domain.Gyroscope) {
				assert.Nil(t, gyro)
			},
		},
		{
			name: "Invalid X (NaN)",
			input: domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				X:         math.NaN(),
				Y:         -2.3,
				Z:         0.8,
			},
			setupMock: func(repo *MockGyroscopeRepository) {
				// Repo.Create should not be called
			},
			wantErr: domain.ErrInvalidGyroscopeValues,
			validateResult: func(t *testing.T, gyro *domain.Gyroscope) {
				assert.Nil(t, gyro)
			},
		},
		{
			name: "Invalid Y (Inf)",
			input: domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				X:         1.5,
				Y:         math.Inf(1),
				Z:         0.8,
			},
			setupMock: func(repo *MockGyroscopeRepository) {
				// Repo.Create should not be called
			},
			wantErr: domain.ErrInvalidGyroscopeValues,
			validateResult: func(t *testing.T, gyro *domain.Gyroscope) {
				assert.Nil(t, gyro)
			},
		},
		{
			name: "Repository error",
			input: domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				X:         1.5,
				Y:         -2.3,
				Z:         0.8,
			},
			setupMock: func(repo *MockGyroscopeRepository) {
				repo.On("Create", mock.AnythingOfType("*domain.Gyroscope")).Return(nil, errors.New("database error")).Once()
			},
			wantErr: domain.ErrSaveGyroscopeData,
			validateResult: func(t *testing.T, gyro *domain.Gyroscope) {
				assert.Nil(t, gyro)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := &MockGyroscopeRepository{}
			tt.setupMock(repo)

			uc := usecase.NewCreateGyroscopeUseCase(repo)

			result, err := uc.Execute(tt.input)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}

			tt.validateResult(t, result)

			repo.AssertExpectations(t)
		})
	}
}
