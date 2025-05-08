package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iamrosada0/v3/internal/infra/api/handlers"
)

func SetupRouter(gyroHandler *handlers.GyroscopeHandler) *gin.Engine {
	router := gin.Default()

	telemetry := router.Group("/telemetry")
	{
		telemetry.POST("/gyroscope", gyroHandler.Create)
	}

	return router
}
