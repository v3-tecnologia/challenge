package controller

import (
	"challenge-v3-backend/internal/application/usecase"
	"challenge-v3-backend/internal/errors"
	"challenge-v3-backend/internal/interface/dto"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type PicturesController struct {
	picturesUseCase *usecase.PicturesUseCaseImpl
}

func NewPicturesController(picturesUseCase *usecase.PicturesUseCaseImpl) *PicturesController {
	return &PicturesController{
		picturesUseCase: picturesUseCase,
	}
}

func (h *PicturesController) Create(ctx *gin.Context) {
	deviceId := ctx.Query("device_id")

	if deviceId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "device_id is required"})
		return
	}

	createdAt := ctx.Query("created_at")

	if createdAt == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "created_at is required"})
		return
	}

	file, fileHeader, err := ctx.Request.FormFile("photo")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "photo is required"})
		return
	}

	defer file.Close()

	fileSizeLimit := 5 * 1024 * 1024

	fileType := fileHeader.Header.Get("Content-Type")

	if fileHeader.Size > int64(fileSizeLimit) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "photo size must be less than 5MB"})
		return
	}

	if !isValidImageType(fileType) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "photo must be a valid image type"})
		return
	}

	imageData, err := ioutil.ReadAll(file)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "photo must be a valid image type"})
		return
	}

	var input dto.CreatePictureRequestDTO

	input.DeviceId = deviceId
	input.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	input.PictureData = imageData
	input.PictureType = fileType

	result, err := h.picturesUseCase.Create(ctx.Request.Context(), input)

	if err != nil {
		errors.RespondWithError(ctx, http.StatusInternalServerError, "error getting ", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func isValidImageType(contentType string) bool {
	validTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
		"image/bmp":  true,
	}
	return validTypes[contentType]
}
