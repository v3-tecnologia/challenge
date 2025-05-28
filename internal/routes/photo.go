package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/KaiRibeiro/challenge/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PhotoService interface {
	AddPhoto(photo models.PhotoModel) error
}

type PhotoHandler struct {
	Service services.PhotoService
}

func NewPhotoHandler(s services.PhotoService) *PhotoHandler {
	return &PhotoHandler{Service: s}
}

func (h *PhotoHandler) SavePhoto(c *gin.Context) {
	var photo models.PhotoModel

	if err := c.ShouldBindJSON(&photo); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]string, len(ve))
			for i, fe := range ve {
				out[i] = utils.FormatFieldError(fe)
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recognized, err := h.Service.AddPhoto(photo)
	if err != nil {
		log.Printf("Error processing SavePhoto: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing photo: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Photo Saved Successfully", "recognized": recognized})
}
