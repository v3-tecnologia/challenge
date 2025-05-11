package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada0/v3/internal/usecase"
)

// GyroscopeHandlers manages the API endpoints for Gyroscope entities.
type GyroscopeHandlers struct {
	CreateGyroscopeUseCase *usecase.CreateGyroscopeUseCase
}

// NewGyroscopeHandlers creates a new instance of GyroscopeHandlers.
func NewGyroscopeHandlers(createGyroscopeUseCase *usecase.CreateGyroscopeUseCase) *GyroscopeHandlers {
	return &GyroscopeHandlers{
		CreateGyroscopeUseCase: createGyroscopeUseCase,
	}
}

// SetupRoutes configures the API routes for Gyroscope endpoints.
func (h *GyroscopeHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/telemetry")
	{
		api.POST("/gyroscope", h.CreateGyroscopeHandler)
	}
}

// CreateGyroscopeHandler handles the POST /telemetry/gyroscope endpoint.
func (h *GyroscopeHandlers) CreateGyroscopeHandler(c *gin.Context) {
	var input usecase.GyroscopeInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid fields"})
		return
	}

	gyro, err := h.CreateGyroscopeUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gyro)
}
