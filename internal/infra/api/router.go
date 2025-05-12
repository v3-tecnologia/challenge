package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	gyroscopeHandlers *GyroscopeHandlers,
	gpsHandlers *GPSHandlers,
	photoHandlers *PhotoHandlers,
) *gin.Engine {
	router := gin.Default()

	gyroscopeHandlers.SetupRoutes(router)
	gpsHandlers.SetupRoutes(router)
	photoHandlers.SetupRoutes(router)

	return router
}
