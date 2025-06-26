package bootstrap

import (
	routers "v3-test/internal/config/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter(controllers Controllers) *gin.Engine {
	r := gin.Default()
	routers.TelemetryRouter(r, controllers.GpsController, controllers.GyroscopeController, controllers.PhotoController)
	return r
}
