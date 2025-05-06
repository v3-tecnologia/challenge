package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wellmtx/challenge/internal/dtos"
	_ "github.com/wellmtx/challenge/internal/infra/entities"
	"github.com/wellmtx/challenge/internal/service"
	_ "gorm.io/gorm"
)

type GeolocationController struct {
	geolocationService *service.GeolocationService
}

func NewGeolocationController(geolocationService *service.GeolocationService) *GeolocationController {
	return &GeolocationController{
		geolocationService: geolocationService,
	}
}

// SaveData godoc
// @Summary      Save GPS data
// @Description  Save GPS data
// @Tags         telemetry
// @Accept       json
// @Produce      json
// @Param        data  body      dtos.GeolocationDataDto  true  "GPS data"
// @Success      200  {object}  entities.GeolocationEntity
// @Failure      400  {object}  dtos.ErrorResponseDTO
// @Failure      500  {object}  dtos.ErrorResponseDTO
// @Router       /telemetry/gps [post]
func (g *GeolocationController) SaveData(c *gin.Context) {
	var data dtos.GeolocationDataDto
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, &dtos.ErrorResponseDTO{
			Message: "Invalid request data",
			Code:    400,
		})
		return
	}

	result, err := g.geolocationService.SaveData(data)
	if err != nil {
		c.JSON(500, &dtos.ErrorResponseDTO{
			Message: "Failed to save GPS data",
			Code:    500,
		})
		return
	}

	c.JSON(200, result)
}
