package usecases

import (
	dtos "v3-test/internal/dtos/telemetries"
	models "v3-test/internal/models/telemetries"
	"v3-test/internal/repositories/telemetries"
)

type GyroscopeUsecase struct {
	repo telemetries.GyroscopeRepository
}

func NewGyroscopeUsecase(repo telemetries.GyroscopeRepository) GyroscopeUsecase {
	return GyroscopeUsecase{repo: repo}
}

func (u *GyroscopeUsecase) CreateGyroscope(gyroscopeDto dtos.CreateGyroscopeDto) (models.GyroscopeModel, error) {
	gyroscopeModel := models.GyroscopeModel{
		X: gyroscopeDto.X,
		Y: gyroscopeDto.Y,
		Z: gyroscopeDto.Z,
	}

	newGyroscope, err := u.repo.CreateGyroscope(gyroscopeModel)
	if err != nil {
		return models.GyroscopeModel{}, err
	}

	return newGyroscope, nil
}
