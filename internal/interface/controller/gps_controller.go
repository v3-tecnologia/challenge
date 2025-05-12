package controller

import (
	"challenge-v3-backend/internal/application/usecase"
	"challenge-v3-backend/internal/errors"
	"challenge-v3-backend/internal/interface/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GPSController struct {
	gpsUseCase *usecase.GPSUseCaseImpl
}

func NewGPSController(gpsUseCase *usecase.GPSUseCaseImpl) *GPSController {
	return &GPSController{
		gpsUseCase: gpsUseCase,
	}
}

func (h *GPSController) Create(ctx *gin.Context) {
	var input dto.CreateGPSRequestDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		errorResponse := errors.HandleValidationErrors(err)
		ctx.JSON(errorResponse.Status, errorResponse)
		return
	}

	err := h.gpsUseCase.Create(ctx.Request.Context(), input)

	if err != nil {
		errors.RespondWithError(ctx, http.StatusInternalServerError, "error creating gps telemetry", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "gps telemetry created successfully"})
}
