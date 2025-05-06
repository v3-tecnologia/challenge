package controller_test

import (
	"github.com/wellmtx/challenge/internal/http/controller"
	"github.com/wellmtx/challenge/internal/infra/database"
	"github.com/wellmtx/challenge/internal/infra/providers"
	"github.com/wellmtx/challenge/internal/infra/repositories"
	"github.com/wellmtx/challenge/internal/service"
)

var (
	db                    *database.Database
	geolocationController *controller.GeolocationController
	gyroscopeController   *controller.GyroscopeController
	photoController       *controller.PhotoController
	gyroscopeService      *service.GyroscopeService
	photoService          *service.PhotoService
	geolocationService    *service.GeolocationService
	geolocationRepository repositories.GeolocationRepository
	gyroscopeRepository   repositories.GyroscopeRepository
	photoRepository       repositories.PhotoRepository
	recognitionProvider   providers.RecognitionProvider
)

func init() {
	db = database.NewDatabase(
		"test",
		"test",
		"test",
		"test",
		true,
	)
	if err := db.Connect(); err != nil {
		panic(err)
	}

	geolocationRepository = repositories.NewGeolocationRepository(db)
	gyroscopeRepository = repositories.NewGyroscopeRepository(db)
	geolocationService = service.NewGeolocationService(geolocationRepository)
	gyroscopeService = service.NewGyroscopeService(gyroscopeRepository)
	geolocationController = controller.NewGeolocationController(geolocationService)
	gyroscopeController = controller.NewGyroscopeController(gyroscopeService)

	photoRepository = repositories.NewPhotoRepository(db)
	recognitionProvider = providers.NewRecognitionProviderInMemory()
	photoService = service.NewPhotoService(photoRepository, recognitionProvider)
	photoController = controller.NewPhotoController(photoService)
	photoController.TestMode = true
}
