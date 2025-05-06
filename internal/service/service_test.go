package service_test

import (
	"github.com/wellmtx/challenge/internal/infra/database"
	"github.com/wellmtx/challenge/internal/infra/providers"
	"github.com/wellmtx/challenge/internal/infra/repositories"
	"github.com/wellmtx/challenge/internal/service"
)

var (
	db                    *database.Database
	gyroscopeRepository   repositories.GyroscopeRepository
	gyroscopeService      *service.GyroscopeService
	geolocationRepository repositories.GeolocationRepository
	geolocationService    *service.GeolocationService
	photoRepository       repositories.PhotoRepository
	photoService          *service.PhotoService
)

func init() {
	db = database.NewDatabase(
		"test",
		"test",
		"test",
		"test",
		true,
	)
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	gyroscopeRepository = repositories.NewGyroscopeRepository(db)
	gyroscopeService = service.NewGyroscopeService(gyroscopeRepository)
	geolocationRepository = repositories.NewGeolocationRepository(db)
	geolocationService = service.NewGeolocationService(geolocationRepository)

	recognitionProviderInMemory := providers.NewRecognitionProviderInMemory()
	photoRepository = repositories.NewPhotoRepository(db)
	photoService = service.NewPhotoService(photoRepository, recognitionProviderInMemory)
}
