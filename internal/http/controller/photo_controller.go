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

func (p *PhotoController) RecognizePhoto(c *gin.Context) {
	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"error": "Unable to get image from request"})
		return
	}

	savePath := "./uploads/" + utils.NormalizeText(image.Filename)
	if !p.TestMode {
		err = c.SaveUploadedFile(image, savePath)
		if err != nil {
			c.JSON(400, gin.H{"error": "Unable to save image"})
			return
		}
	}

	file, err := image.Open()
	if err != nil {
		c.JSON(400, gin.H{"error": "Unable to open image"})
		return
	}
	defer file.Close()

	imageBytes := make([]byte, image.Size)
	_, err = file.Read(imageBytes)
	if err != nil {
		c.JSON(400, gin.H{"error": "Unable to read image"})
		return
	}

	result, err := p.photoService.RecognizePhoto(dtos.SavePhotoDTO{
		Image:    imageBytes,
		FilePath: savePath,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to recognize photo"})
		return
	}

	c.JSON(200, result)
}
