// package routes

package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/KaiRibeiro/challenge/internal/logs"
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/KaiRibeiro/challenge/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GyroscopeService interface {
	AddGyroscope(gyroscope models.GyroscopeModel) error
}

type GyroscopeHandler struct {
	Service services.GyroscopeService
}

func NewGyroscopeHandler(s services.GyroscopeService) *GyroscopeHandler {
	return &GyroscopeHandler{Service: s}
}

func (h *GyroscopeHandler) SaveGyroscope(c *gin.Context) {
	var gyroscope models.GyroscopeModel

	if err := c.ShouldBindJSON(&gyroscope); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]string, len(ve))
			for i, fe := range ve {
				out[i] = utils.FormatFieldError(fe)
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			logs.Logger.Error("Bad gps request",
				"error", err,
			)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logs.Logger.Error("Bad gps request",
			"error", err,
		)
		return
	}

	err := h.Service.AddGyroscope(gyroscope)
	if err != nil {
		log.Printf("Error processing SaveGyroscope: %v", err)
		var dbErr *custom_errors.DBError

		if errors.As(err, &dbErr) {
			c.JSON(dbErr.Status(), gin.H{"error": "A database error occurred", "details": "Please try again later."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected server error occurred", "details": "Please try again later."})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Gyroscope Saved Successfully"})
}
