package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/KaiRibeiro/challenge/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GpsService interface {
	AddGps(gps models.GpsModel) error
}

type GpsHandler struct {
	Service services.GpsService
}

func NewGpsHandler(s services.GpsService) *GpsHandler {
	return &GpsHandler{Service: s}
}

func (h *GpsHandler) SaveGps(c *gin.Context) {
	var gps models.GpsModel

	if err := c.ShouldBindJSON(&gps); err != nil {
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

	err := h.Service.AddGps(gps)
	if err != nil {
		log.Printf("Error processing SaveGps: %v", err)
		var dbErr *custom_errors.DBError

		if errors.As(err, &dbErr) {
			c.JSON(dbErr.Status(), gin.H{"error": "A database error occurred", "details": "Please try again later."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected server error occurred", "details": "Please try again later."})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "GPS Saved Successfully"})
}
