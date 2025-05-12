package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wellmtx/challenge/internal/dtos"
	"github.com/wellmtx/challenge/internal/infra/utils"
	"github.com/wellmtx/challenge/internal/service"
)

type PhotoController struct {
	photoService *service.PhotoService
	TestMode     bool
}

func NewPhotoController(photoService *service.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: photoService,
	}
}

// RecognizePhoto godoc
// @Summary      Recognize photo
// @Description  Recognize photo
// @Tags         telemetry
// @Accept       multipart/form-data
// @Produce      json
// @Param        image  formData  file  true  "Image file"
// @Param        mac_address  formData  string  false  "MAC address"
// @Param        timestamp  formData  string  false  "Timestamp"
// @Success      200  {object}  dtos.SavePhotoResponseDTO
// @Failure      400  {object}  dtos.ErrorResponseDTO
// @Failure      500  {object}  dtos.ErrorResponseDTO
// @Router       /telemetry/photo [post]
func (p *PhotoController) RecognizePhoto(c *gin.Context) {
	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, dtos.ErrorResponseDTO{
			Message: "Invalid image file",
			Code:    400,
		})
		return
	}
	macAddress := c.PostForm("mac_address")
	timestamp := c.PostForm("timestamp")
	if macAddress == "" {
		c.JSON(400, dtos.ErrorResponseDTO{
			Message: "MAC address is required",
			Code:    400,
		})
		return
	}
	if timestamp == "" {
		c.JSON(400, dtos.ErrorResponseDTO{
			Message: "Timestamp is required",
			Code:    400,
		})
		return
	}

	timestampInt, err := strconv.Atoi(timestamp)
	if err != nil {
		c.JSON(400, dtos.ErrorResponseDTO{
			Message: "Invalid timestamp",
			Code:    400,
		})
		return
	}

	savePath := "./uploads/" + utils.NormalizeText(image.Filename)
	if !p.TestMode {
		err = c.SaveUploadedFile(image, savePath)
		if err != nil {
			c.JSON(400, dtos.ErrorResponseDTO{
				Message: "Unable to save image",
				Code:    400,
			})
			return
		}
	}

	file, err := image.Open()
	if err != nil {
		c.JSON(400, dtos.ErrorResponseDTO{
			Message: "Unable to open image",
			Code:    400,
		})
		return
	}
	defer file.Close()

	imageBytes := make([]byte, image.Size)
	_, err = file.Read(imageBytes)
	if err != nil {
		c.JSON(400, dtos.ErrorResponseDTO{
			Message: "Unable to read image",
			Code:    400,
		})
		return
	}

	result, err := p.photoService.RecognizePhoto(dtos.SavePhotoDTO{
		Image:    imageBytes,
		FilePath: savePath,
		BaseDTO: dtos.BaseDTO{
			MacAddress: macAddress,
			Timestamp:  int64(timestampInt),
		},
	})
	if err != nil {
		c.JSON(500, dtos.ErrorResponseDTO{
			Message: "Failed to recognize photo",
			Code:    500,
		})
		return
	}

	c.JSON(200, result)
}
