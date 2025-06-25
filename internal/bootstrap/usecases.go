package bootstrap

import (
	"v3-test/internal/usecases/telemetries"
)

type Usecases struct {
	GpsUsecase telemetries.GpsUsecase
}

func SetupUsecases(repos Repositories) Usecases {
	return Usecases{
		GpsUsecase: telemetries.NewGpsUsecase(repos.GpsRepo),
	}
}
