package controllers

import (
	"net/http"
	dtos "v3-test/internal/dtos/telemetries"
	usecases "v3-test/internal/usecases/telemetries"
	"v3-test/internal/validators"

	"github.com/gin-gonic/gin"
)

type gpsController struct {
	usecase usecases.GpsUsecase
}

func NewGpsController() gpsController {
	return gpsController{
		usecase: usecases.NewGpsUsecase(),
	}
}

func (gpsController *gpsController) CreateGps(ctx *gin.Context) {
	var gpsDto dtos.GpsDto

	isValid := validators.BindAndValidate(ctx, &gpsDto)
	if !isValid {
		return
	}

	gpsModel, err := gpsController.usecase.CreateGps(gpsDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create GPS data"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": gpsModel})
}
