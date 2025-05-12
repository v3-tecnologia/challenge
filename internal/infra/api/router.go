package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter configures all API routes.
func SetupRouter(
	gyroscopeHandlers *GyroscopeHandlers,
	gpsHandlers *GPSHandlers,
	photoHandlers *PhotoHandlers,
) *gin.Engine {
	router := gin.Default()

	// Setup routes for each handler
	gyroscopeHandlers.SetupRoutes(router)
	gpsHandlers.SetupRoutes(router)
	photoHandlers.SetupRoutes(router)

	return router
}
