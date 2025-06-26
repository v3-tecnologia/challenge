package routers

import (
	"v3-test/internal/controllers/commonController"
	controllers "v3-test/internal/controllers/telemetriesControllers"

	"github.com/gin-gonic/gin"
)

func TelemetryRouter(
	r *gin.Engine, gpsController controllers.GpsController,
	gyroscopeController controllers.GyroscopeController,
	photoController commonController.PhotoController,
) {
	api := r.Group("/telemetry")
	{
		api.POST("/gps", gpsController.CreateGps)
		api.POST("/gyroscope", gyroscopeController.CreateGyroscope)
		api.POST("/photo", photoController.UploadTelemetryPhoto)
	}
}
