package controller

import (
	"challenge-v3-backend/internal/application/usecase"
	"challenge-v3-backend/internal/errors"
	"challenge-v3-backend/internal/interface/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GyroscopeController struct {
	gyroscopeUseCase *usecase.GyroscopeUseCaseImpl
}

func NewGyroscopeController(gyroscopeUseCase *usecase.GyroscopeUseCaseImpl) *GyroscopeController {
	return &GyroscopeController{
		gyroscopeUseCase: gyroscopeUseCase,
	}
}

func (h *GyroscopeController) Create(ctx *gin.Context) {
	var input dto.CreateGyroscopeRequestDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		errorResponse := errors.HandleValidationErrors(err)
		ctx.JSON(errorResponse.Status, errorResponse)
		return
	}

	err := h.gyroscopeUseCase.Create(ctx.Request.Context(), input)

	if err != nil {
		errors.RespondWithError(ctx, http.StatusInternalServerError, "error creating gyroscope telemetry", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Gyroscope telemetry created successfully"})
}
