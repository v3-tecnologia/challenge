package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wellmtx/challenge/internal/dtos"
	_ "github.com/wellmtx/challenge/internal/infra/entities"
	"github.com/wellmtx/challenge/internal/service"
	_ "gorm.io/gorm"
)

type GyroscopeController struct {
	gyroscopeService *service.GyroscopeService
}

// SaveData godoc
// @Summary      Save gyroscope data
// @Description  Save gyroscope data
// @Tags         telemetry
// @Accept       json
// @Produce      json
// @Param        data  body      dtos.GyroscopeDataDto  true  "Gyroscope data"
// @Success      200  {object}  entities.GyroscopeEntity
// @Failure      400  {object}  dtos.ErrorResponseDTO
// @Failure      500  {object}  dtos.ErrorResponseDTO
// @Router       /telemetry/gyroscope [post]
func NewGyroscopeController(gyroscopeService *service.GyroscopeService) *GyroscopeController {
	return &GyroscopeController{
		gyroscopeService: gyroscopeService,
	}
}

func (g *GyroscopeController) SaveData(c *gin.Context) {
	var data dtos.GyroscopeDataDto
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, dtos.ErrorResponseDTO{
			Message: "Invalid request data",
			Code:    400,
		})
		return
	}

	result, err := g.gyroscopeService.SaveData(data)
	if err != nil {
		c.JSON(500, dtos.ErrorResponseDTO{
			Message: "Failed to save gyroscope data",
			Code:    500,
		})
		return
	}

	c.JSON(200, result)
}
