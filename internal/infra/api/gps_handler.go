package api

import (
	"net/http"

	"v3/internal/domain"
	usecase "v3/internal/usecase"

	"github.com/gin-gonic/gin"
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

func (h *GPSHandlers) CreateGPSHandler(c *gin.Context) {
	var input domain.GPSDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrMissingGPSInvalidFields})
		return
	}

	gps, err := h.CreateGPSUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gps)
}
