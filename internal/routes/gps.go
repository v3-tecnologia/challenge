package routes

import (
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var gpsService = services.AddGps

func SetGpsService(service func(models.GpsModel) error) {
	gpsService = service
}

func SaveGps(c *gin.Context) {
	var gps models.GpsModel

	if err := c.ShouldBindJSON(&gps); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect or missing parameters"})
		return
	}

	err := gpsService(gps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving gps to database: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "GPS Saved Successfully"})
}
