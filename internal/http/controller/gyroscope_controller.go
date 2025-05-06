package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wellmtx/challenge/internal/dtos"
	"github.com/wellmtx/challenge/internal/service"
)

type GyroscopeController struct {
	gyroscopeService *service.GyroscopeService
}

func NewGyroscopeController(gyroscopeService *service.GyroscopeService) *GyroscopeController {
	return &GyroscopeController{
		gyroscopeService: gyroscopeService,
	}
}

func (g *GyroscopeController) SaveData(c *gin.Context) {
	var data dtos.GyroscopeDataDto
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := g.gyroscopeService.SaveData(data)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success", "data": result})
}
