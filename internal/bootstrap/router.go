package bootstrap

import (
	"v3-test/internal/config/routers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(controllers Controllers) *gin.Engine {
	r := gin.Default()
	routers.TelemetryRouter(r, controllers.GpsController)
	return r
}
