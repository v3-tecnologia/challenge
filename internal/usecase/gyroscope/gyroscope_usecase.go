package gyroscope

import (
	"github.com/iamrosada0/v3/internal/domain"
	"github.com/iamrosada0/v3/internal/repository/gyroscope"
)

type GyroscopeInputDto struct {
	DeviceID  string  `json:"deviceId"`
	Timestamp int64   `json:"timestamp"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
}

type CreateGyroscopeUseCase struct {
	Repo gyroscope.GyroscopeRepository
}

func NewCreateGyroscopeUseCase(repo gyroscope.GyroscopeRepository) *CreateGyroscopeUseCase {
	return &CreateGyroscopeUseCase{Repo: repo}
}

func (uc *CreateGyroscopeUseCase) Execute(input GyroscopeInputDto) (*domain.Gyroscope, error) {
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

	savedGyro, err := uc.Repo.Create(gyro)
	if err != nil {
		return nil, domain.ErrSaveGyroscopeData
	}

	return savedGyro, nil
}
