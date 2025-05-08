package handlers

import (
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

}
