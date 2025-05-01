package routes

import (
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveGps(c *gin.Context) {
	var gps models.GpsModel

	if err := c.ShouldBindJSON(&gps); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.AddGps(gps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving gps to database: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "GPS Saved Successfully"})
}
