package routes

import (
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SavePhoto(c *gin.Context) {
	var photo models.PhotoModel

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.AddPhoto(photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving photo to database: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Photo saved Successfully"})
}
