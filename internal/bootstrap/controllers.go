package bootstrap

import (
	"v3-test/internal/controllers/telemetriesControllers"
)

type Controllers struct {
	GpsController       telemetriesControllers.GpsController
	GyroscopeController telemetriesControllers.GyroscopeController
}

func SetupControllers(usecases Usecases) Controllers {
	return Controllers{
		GpsController:       telemetriesControllers.NewGpsController(usecases.GpsUsecase),
		GyroscopeController: telemetriesControllers.NewGyroscopeController(usecases.GyroscopeUsecase),
	}
}
