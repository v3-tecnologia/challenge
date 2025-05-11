package gyroscope

import (
	"errors"

	"github.com/iamrosada0/v3/internal/domain"
	"github.com/iamrosada0/v3/internal/repository/gyroscope"
)

// GyroscopeInputDto defines the input structure for creating a Gyroscope.
type GyroscopeInputDto struct {
	DeviceID  string  `json:"deviceId"`
	Timestamp int64   `json:"timestamp"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
}

// CreateGyroscopeUseCase handles the creation of a Gyroscope entity.
type CreateGyroscopeUseCase struct {
	Repo gyroscope.GyroscopeRepository
}

// NewCreateGyroscopeUseCase creates a new instance of CreateGyroscopeUseCase.
func NewCreateGyroscopeUseCase(repo gyroscope.GyroscopeRepository) *CreateGyroscopeUseCase {
	return &CreateGyroscopeUseCase{Repo: repo}
}

// Execute creates a new Gyroscope entity and saves it to the database.
func (uc *CreateGyroscopeUseCase) Execute(input GyroscopeInputDto) (*domain.Gyroscope, error) {
	// Transform input DTO to domain entity
	gyro, err := domain.NewGyroscopeData(&domain.GyroscopeDto{
		DeviceID:  input.DeviceID,
		Timestamp: input.Timestamp,
		X:         input.X,
		Y:         input.Y,
		Z:         input.Z,
	})
	if err != nil {
		return nil, err
	}

	// Save to repository
	savedGyro, err := uc.Repo.Create(gyro)
	if err != nil {
		return nil, errors.New("failed to save gyroscope data")
	}

	return savedGyro, nil
}
