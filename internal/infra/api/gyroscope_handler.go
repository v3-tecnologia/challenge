package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada0/v3/internal/domain"
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

func (h *GyroscopeHandlers) CreateGyroscopeHandler(c *gin.Context) {
	var input domain.GyroscopeDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrMissingGPSInvalidFields})
		return
	}

	gyro, err := h.CreateGyroscopeUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gyro)
}
