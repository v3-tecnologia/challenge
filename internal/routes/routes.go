package routes

import (
	v1 "go-challenge/api/v1"
	"go-challenge/internal/messaging"
	"go-challenge/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(service *services.TelemetryService, nats *messaging.NATSProducer) *gin.Engine {
	r := gin.Default()

	handler := v1.NewTelemetryHandler(service, nats)

	api := r.Group("/telemetry")
	{
		api.POST("/gyroscope", handler.PostGyroscopeHandler)
		api.POST("/gps", handler.PostGPSHandler)
		api.POST("/photo", handler.PostPhotoHandler)
	}

	return r
}
