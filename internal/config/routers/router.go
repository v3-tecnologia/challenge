package routers

import (
	controllers "v3-test/internal/controllers/telemetries"

	"github.com/gin-gonic/gin"
)

func SetupRouter(gpsController controllers.GpsController) *gin.Engine {
	r := gin.Default()

	TelemetryRouter(r, gpsController)

	return r
}
