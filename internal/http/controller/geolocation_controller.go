package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wellmtx/challenge/internal/dtos"
	"github.com/wellmtx/challenge/internal/service"
)

type GeolocationController struct {
	geolocationService *service.GeolocationService
}

func NewGeolocationController(geolocationService *service.GeolocationService) *GeolocationController {
	return &GeolocationController{
		geolocationService: geolocationService,
	}
}

func (g *GeolocationController) SaveData(c *gin.Context) {
	var data dtos.GeolocationDataDto
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := g.geolocationService.SaveData(data)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success", "data": result})
}
