package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada0/v3/internal/usecase/gyroscope"
)

type GyroscopeHandler struct {
	UseCase *gyroscope.GyroscopeUseCase
}

func NewGyroscopeHandler(useCase *gyroscope.GyroscopeUseCase) *GyroscopeHandler {
	return &GyroscopeHandler{UseCase: useCase}
}

func (h *GyroscopeHandler) Create(c *gin.Context) {
	var input struct {
		DeviceID  string    `json:"device_id" binding:"required"`
		X         float64   `json:"x" binding:"required"`
		Y         float64   `json:"y" binding:"required"`
		Z         float64   `json:"z" binding:"required"`
		Timestamp time.Time `json:"timestamp" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid fields"})
		return
	}
	if _, err := h.UseCase.Create(input.DeviceID, input.X, input.Y, input.Z, input.Timestamp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gyroscope data received and saved"})
}
