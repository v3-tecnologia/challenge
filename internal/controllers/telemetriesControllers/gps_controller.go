package telemetriesControllers

import (
	"net/http"
	"v3-test/internal/dtos/telemetriesDtos"
	"v3-test/internal/usecases/telemetriesUsecases"
	"v3-test/internal/validators"

	"github.com/gin-gonic/gin"
)

type GpsController struct {
	usecase telemetriesUsecases.IGpsUsecase
}

func NewGpsController(usecase telemetriesUsecases.IGpsUsecase) GpsController {
	return GpsController{usecase: usecase}
}

func (c *GpsController) CreateGps(ctx *gin.Context) {

	var gpsDto telemetriesDtos.CreateGpsDto

	if !validators.BindAndValidate(ctx, &gpsDto) {
		return
	}

	gpsModel, err := c.usecase.CreateGps(gpsDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create GPS data"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": gpsModel})
}
