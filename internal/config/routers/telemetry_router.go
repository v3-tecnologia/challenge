package routers

import (
	controllers "v3-test/internal/controllers/telemetries"

	"github.com/gin-gonic/gin"
)

func TelemetryRouter(r *gin.Engine, gpsController controllers.GpsController) {
	api := r.Group("/telemetry")
	{
		api.POST("/gps", gpsController.CreateGps)
	}
}
