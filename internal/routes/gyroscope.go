package routes

import (
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var gyroscopeService = services.AddGyroscope

func SetGyroscopeService(service func(model models.GyroscopeModel) error) {
	gyroscopeService = service
}

func SaveGyroscope(c *gin.Context) {
	var gyroscope models.GyroscopeModel

	if err := c.ShouldBindJSON(&gyroscope); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect or missing parameters"})
		return
	}

	err := gyroscopeService(gyroscope)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving gyroscope to database: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Gyroscope Saved Successfully"})
}
