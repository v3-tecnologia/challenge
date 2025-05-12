package router

import (
	"challenge-v3-backend/internal/application/usecase"
	"challenge-v3-backend/internal/infrastructure/repository"
	"challenge-v3-backend/internal/interface/controller"
	services "challenge-v3-backend/pkg/services/aws"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	controllers := setupControllers(db)

	applyGlobalMiddlewares(router)

	group := router.Group("")

	setupPublicRoutes(group, controllers)

	return router
}

func setupControllers(db *gorm.DB) *Controllers {
	gpsRepository := repository.NewGPSRepository(db)
	gyroscopeRepository := repository.NewGyroscopeTelemetryRepository(db)
	picturesRepository := repository.NewPicturesRepository(db)

	s3Client, _ := services.NewS3Client()
	rekognitionClient, _ := services.NewRekognitionClient()

	gpsUseCase := usecase.NewGPSUseCase(gpsRepository)
	gyroscopeUseCase := usecase.NewGyroscopeUseCase(gyroscopeRepository)
	picturesUseCase := usecase.NewPicturesUseCase(picturesRepository, s3Client, rekognitionClient)

	gpsController := controller.NewGPSController(gpsUseCase)
	gyroscopyController := controller.NewGyroscopeController(gyroscopeUseCase)
	picturesController := controller.NewPicturesController(picturesUseCase)

	return &Controllers{
		GPSTelemetryController: gpsController,
		GyroscopeController:    gyroscopyController,
		PicturesController:     picturesController,
	}
}

type Controllers struct {
	GPSTelemetryController *controller.GPSController
	GyroscopeController    *controller.GyroscopeController
	PicturesController     *controller.PicturesController
}

func applyGlobalMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
}

func setupPublicRoutes(router *gin.RouterGroup, controllers *Controllers) {
	telemetryGroup := router.Group("/telemetry")

	telemetryGroup.POST("/gps", controllers.GPSTelemetryController.Create)
	telemetryGroup.POST("/gyroscope", controllers.GyroscopeController.Create)
	telemetryGroup.POST("/photo", controllers.PicturesController.Create)
}
