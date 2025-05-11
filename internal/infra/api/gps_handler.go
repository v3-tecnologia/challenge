package api

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/iamrosada0/v3/internal/usecase/gps"
)

type GPSHandlers struct {
	CreateGPSUseCase *usecase.CreateGPSUseCase
}

func NewGPSHandlers(createGPSUseCase *usecase.CreateGPSUseCase) *GPSHandlers {
	return &GPSHandlers{
		CreateGPSUseCase: createGPSUseCase,
	}
}

func (h *GPSHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/telemetry")
	{
		api.POST("/gps", h.CreateGPSHandler)
	}
}
