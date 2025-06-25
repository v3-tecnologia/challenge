package telemetries

import (
	"net/http"
	dtos "v3-test/internal/dtos/telemetries"
	usecases "v3-test/internal/usecases/telemetries"
	"v3-test/internal/validators"

	"github.com/gin-gonic/gin"
)

type GpsController struct {
	usecase usecases.GpsUsecase
}

func NewGpsController(usecase usecases.GpsUsecase) GpsController {
	return GpsController{usecase: usecase}
}

func (c *GpsController) CreateGps(ctx *gin.Context) {

	var gpsDto dtos.CreateGpsDto

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
