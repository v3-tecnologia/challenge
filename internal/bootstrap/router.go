package bootstrap

import (
	"v3-test/internal/config/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter(controllers Controllers) *gin.Engine {
	r := gin.Default()
	routes.TelemetryRouter(r, controllers.GpsController, controllers.GyroscopeController, controllers.PhotoController)
	// health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	return r
}
