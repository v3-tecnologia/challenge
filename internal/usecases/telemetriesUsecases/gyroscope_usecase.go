package telemetriesUsecases

import (
	"time"
	dtos "v3-test/internal/dtos/telemetriesDtos"
	models "v3-test/internal/models/telemetriesModels"
	"v3-test/internal/repositories/telemetriesRepositories"
)

type IGyroscopeUsecase interface {
	CreateGyroscope(gyroscopeDto dtos.CreateGyroscopeDto) (models.GyroscopeModel, error)
}

type GyroscopeUsecase struct {
	repo telemetriesRepositories.GyroscopeRepository
}

func NewGyroscopeUsecase(repo telemetriesRepositories.GyroscopeRepository) GyroscopeUsecase {
	return GyroscopeUsecase{repo: repo}
}

func (u *GyroscopeUsecase) CreateGyroscope(gyroscopeDto dtos.CreateGyroscopeDto) (models.GyroscopeModel, error) {
	gyroscopeModel := models.GyroscopeModel{
		X:         gyroscopeDto.X,
		Y:         gyroscopeDto.Y,
		Z:         gyroscopeDto.Z,
		Timestamp: time.Now(),
	}

	newGyroscope, err := u.repo.CreateGyroscope(gyroscopeModel)
	if err != nil {
		return models.GyroscopeModel{}, err
	}

	return newGyroscope, nil
}
