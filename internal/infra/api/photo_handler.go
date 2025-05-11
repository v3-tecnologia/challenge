package api

import (
	"errors"
	"net/http"
	"v3/internal/domain"
	"v3/internal/usecase"

	"github.com/gin-gonic/gin"
)

type PhotoHandlers struct {
	CreatePhotoUseCase *usecase.CreatePhotoUseCase
}

func NewPhotoHandlers(createPhotoUseCase *usecase.CreatePhotoUseCase) *PhotoHandlers {
	return &PhotoHandlers{
		CreatePhotoUseCase: createPhotoUseCase,
	}
}

func (h *PhotoHandlers) SetupRoutes(router *gin.Engine) {
	router.POST("/api/telemetry/photo", h.CreatePhotoHandler)
}

func (h *PhotoHandlers) CreatePhotoHandler(c *gin.Context) {
	var input domain.PhotoDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrMissingPhotoInvalidFields.Error()})
		return
	}

	photo, err := h.CreatePhotoUseCase.Execute(input)
	if err != nil {
		if errors.Is(err, domain.ErrDeviceIDPhoto) || errors.Is(err, domain.ErrTimestampPhoto) || errors.Is(err, domain.ErrPhotoData) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, photo)
}
