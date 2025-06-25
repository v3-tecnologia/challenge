package bootstrap

import (
	"v3-test/internal/controllers/telemetries"
)

type Controllers struct {
	GpsController telemetries.GpsController
}

func SetupControllers(usecases Usecases) Controllers {
	return Controllers{
		GpsController: telemetries.NewGpsController(usecases.GpsUsecase),
	}
}
