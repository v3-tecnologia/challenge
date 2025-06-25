package bootstrap

import (
	usecases "v3-test/internal/usecases/telemetries"
)

type Usecases struct {
	GpsUsecase       usecases.GpsUsecase
	GyroscopeUsecase usecases.GyroscopeUsecase
}

func SetupUsecases(repos Repositories) Usecases {
	return Usecases{
		GpsUsecase:       usecases.NewGpsUsecase(repos.GpsRepo),
		GyroscopeUsecase: usecases.NewGyroscopeUsecase(repos.GyroscopeRepo),
	}
}
