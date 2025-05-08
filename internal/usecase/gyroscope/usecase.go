package gyroscope

import (
	repository "github.com/iamrosada0/v3/internal/repository/gyroscope"
)

type GyroscopeUseCase struct {
	Repo repository.GyroscopeRepository
}
