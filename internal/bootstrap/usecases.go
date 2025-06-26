package bootstrap

import (
	"v3-test/internal/usecases/commonUsecases"
	"v3-test/internal/usecases/telemetriesUsecases"
)

type Usecases struct {
	GpsUsecase       telemetriesUsecases.GpsUsecase
	GyroscopeUsecase telemetriesUsecases.GyroscopeUsecase
	PhotoUsecase     commonUsecases.PhotoUsecase
}

func SetupUsecases(repos Repositories) Usecases {
	return Usecases{
		GpsUsecase:       telemetriesUsecases.NewGpsUsecase(repos.GpsRepo),
		GyroscopeUsecase: telemetriesUsecases.NewGyroscopeUsecase(repos.GyroscopeRepo),
		PhotoUsecase:     commonUsecases.NewPhotoUsecase(repos.PhotoRepo),
	}
}
