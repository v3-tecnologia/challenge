package commonController

import (
	"net/http"
	"v3-test/internal/enums"
	"v3-test/internal/usecases/commonUsecases"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	usecase commonUsecases.PhotoUsecase
}

func NewPhotoController(usecase commonUsecases.PhotoUsecase) PhotoController {
	return PhotoController{usecase: usecase}
}

func (c *PhotoController) UploadTelemetryPhoto(ctx *gin.Context) {
	c.uploadPhoto(ctx, enums.Telemetry)
}

func (c *PhotoController) uploadPhoto(ctx *gin.Context, entity enums.PhotoEntity) {
	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from form"})
		return
	}

	photo, err := c.usecase.UploadPhoto(file, entity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload photo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"photo": photo})
}
