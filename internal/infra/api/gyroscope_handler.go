package api

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/iamrosada0/v3/internal/usecase/gyroscope"
)

type GyroscopeHandlers struct {
	CreateGyroscopeUseCase *usecase.CreateGyroscopeUseCase
}

func NewGyroscopeHandlers(createGyroscopeUseCase *usecase.CreateGyroscopeUseCase) *GyroscopeHandlers {
	return &GyroscopeHandlers{
		CreateGyroscopeUseCase: createGyroscopeUseCase,
	}
}

func (h *GyroscopeHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/telemetry")
	{
		api.POST("/gyroscope", h.CreateGyroscopeHandler)
	}
}
