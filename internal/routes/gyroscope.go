package routes

import (
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveGyroscope(c *gin.Context) {
	var gyroscope models.GyroscopeModel

	if err := c.ShouldBindJSON(&gyroscope); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Gyroscope Saved Successfully"})
}
