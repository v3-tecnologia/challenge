package routers

import (
	controllers "v3-test/internal/controllers/telemetries"

	"github.com/gin-gonic/gin"
)

func TelemetryRouter(r *gin.Engine) {
	api := r.Group("/telemetry")
	{
		gpsController := controllers.NewGpsController()
		api.POST("/gps", gpsController.CreateGps)
	}
}
