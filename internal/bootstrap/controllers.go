package bootstrap

import (
	"v3-test/internal/controllers/telemetries"
)

type Controllers struct {
	GpsController       telemetries.GpsController
	GyroscopeController telemetries.GyroscopeController
}

func SetupControllers(usecases Usecases) Controllers {
	return Controllers{
		GpsController:       telemetries.NewGpsController(usecases.GpsUsecase),
		GyroscopeController: telemetries.NewGyroscopeController(usecases.GyroscopeUsecase),
	}
}
