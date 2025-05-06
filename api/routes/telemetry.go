package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	controller "github.com/mkafonso/go-cloud-challenge/api"
	"github.com/mkafonso/go-cloud-challenge/recognition"
	"github.com/mkafonso/go-cloud-challenge/repository"
	"github.com/mkafonso/go-cloud-challenge/storage"
)

func TelemetryModuleRouter(
	gpsRepo repository.GPSRepositoryInterface,
	gyroRepo repository.GyroscopeRepositoryInterface,
	photoRepo repository.PhotoRepositoryInterface,
	storageService storage.PhotoStorageService,
	recognizer recognition.FaceRecognitionService,
) http.Handler {
	router := chi.NewRouter()

	saveGPS := controller.NewSaveGPSDataHandler(gpsRepo)
	router.Post("/gps", saveGPS.Handle)

	saveGyro := controller.NewSaveGyroscopeDataHandler(gyroRepo)
	router.Post("/gyroscope", saveGyro.Handle)

	savePhoto := controller.NewSavePhotoHandler(photoRepo, recognizer, storageService)
	router.Post("/photo", savePhoto.Handle)

	return router
}
