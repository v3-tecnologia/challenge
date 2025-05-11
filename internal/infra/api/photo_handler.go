package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/iamrosada0/v3/internal/usecase/photos"
)

type PhotoHandlers struct {
	CreatePhotoUseCase *usecase.CreatePhotoUseCase
}

func NewPhotoHandlers(createPhotoUseCase *usecase.CreatePhotoUseCase) *PhotoHandlers {
	return &PhotoHandlers{
		CreatePhotoUseCase: createPhotoUseCase,
	}
}

func (h *PhotoHandlers) CreatePhotoHandler(c *gin.Context) {
	var input usecase.PhotoInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid fields"})
		return
	}

	photo, err := h.CreatePhotoUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"photo":      photo,
		"recognized": photo.Recognized,
	})
}
