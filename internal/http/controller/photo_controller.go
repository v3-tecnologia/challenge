package controller

import (
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
