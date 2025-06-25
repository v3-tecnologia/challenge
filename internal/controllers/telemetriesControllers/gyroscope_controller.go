package telemetriesControllers

import (
	"net/http"
	dtos "v3-test/internal/dtos/telemetriesDtos"
	usecases "v3-test/internal/usecases/telemetriesUsecases"
	"v3-test/internal/validators"

	"github.com/gin-gonic/gin"
)

type GyroscopeController struct {
	usecase usecases.GyroscopeUsecase
}

func NewGyroscopeController(usecase usecases.GyroscopeUsecase) GyroscopeController {
	return GyroscopeController{
		usecase: usecase,
	}
}

func (g *GyroscopeController) CreateGyroscope(ctx *gin.Context) {
	var gyroscopeDto dtos.CreateGyroscopeDto
	if !validators.BindAndValidate(ctx, &gyroscopeDto) {
		return
	}

	gyroscopeModel, err := g.usecase.CreateGyroscope(gyroscopeDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create gyroscope data"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": gyroscopeModel})
}
