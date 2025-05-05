package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	controller "github.com/mkafonso/go-cloud-challenge/api"
	"github.com/mkafonso/go-cloud-challenge/recognition"
	"github.com/mkafonso/go-cloud-challenge/repository"
	"github.com/mkafonso/go-cloud-challenge/storage"
	"github.com/mkafonso/go-cloud-challenge/usecase"
)

func TelemetryModuleRouter(
	gpsRepo repository.GPSRepositoryInterface,
	gyroRepo repository.GyroscopeRepositoryInterface,
	photoRepo repository.PhotoRepositoryInterface,
	storageService storage.PhotoStorageService,
	recognizer recognition.FaceRecognitionService,
) http.Handler {
	router := chi.NewRouter()

	saveGPS := usecase.NewSaveGPSData(gpsRepo)
	router.Post("/gps", controller.NewSaveGPSDataHandler(saveGPS).Handle)

	saveGyro := usecase.NewSaveGyroscopeData(gyroRepo)
	router.Post("/gyroscope", controller.NewSaveGyroscopeDataHandler(saveGyro).Handle)

	savePhoto := usecase.NewSavePhoto(photoRepo, recognizer, storageService)
	router.Post("/photo", controller.NewSavePhotoHandler(savePhoto).Handle)

	return router
}
