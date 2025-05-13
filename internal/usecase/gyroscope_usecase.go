package usecase

import (
	"v3/internal/domain"
	"v3/internal/repository/gyroscope"
)

// GyroscopeUseCase define os métodos do CreateGyroscopeUseCase
type GyroscopeUseCase interface {
	Execute(input domain.GyroscopeDto) (*domain.Gyroscope, error)
}

type CreateGyroscopeUseCase struct {
	Repo gyroscope.GyroscopeRepository
}

func NewCreateGyroscopeUseCase(repo gyroscope.GyroscopeRepository) *CreateGyroscopeUseCase {
	return &CreateGyroscopeUseCase{Repo: repo}
}

func (uc *CreateGyroscopeUseCase) Execute(input domain.GyroscopeDto) (*domain.Gyroscope, error) {
	gyroData, err := domain.NewGyroscopeData(&domain.GyroscopeDto{
		DeviceID:  input.DeviceID,
		Timestamp: input.Timestamp,
		X:         input.X,
		Y:         input.Y,
		Z:         input.Z,
	})
	if err != nil {
		return nil, err
	}
	savedGyro, err := uc.Repo.Create(gyroData)
	if err != nil {
		return nil, domain.ErrSaveGyroscopeData
	}
	return savedGyro, nil
}
