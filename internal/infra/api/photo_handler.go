package api

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"v3/internal/domain"
	"v3/internal/usecase"

	"github.com/gin-gonic/gin"
)

type PhotoHandlers struct {
	CreatePhotoUseCase *usecase.CreatePhotoUseCase
}

func NewPhotoHandlers(createPhotoUseCase *usecase.CreatePhotoUseCase) *PhotoHandlers {
	return &PhotoHandlers{
		CreatePhotoUseCase: createPhotoUseCase,
	}
}

func (h *PhotoHandlers) SetupRoutes(router *gin.Engine) {
	router.POST("/api/telemetry/photo", h.CreatePhotoHandler)
}

func (h *PhotoHandlers) CreatePhotoHandler(c *gin.Context) {

	if err := c.Request.ParseMultipartForm(6 << 20); err != nil {
		log.Printf("Failed to parse form data: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse form data"})
		return
	}

	deviceID := strings.TrimSpace(c.PostForm("deviceId"))
	timestampStr := strings.TrimSpace(c.PostForm("timestamp"))
	log.Printf("Received form fields: deviceId=%s, timestamp=%s\n", deviceID, timestampStr)

	if deviceID == "" || timestampStr == "" {
		log.Println("Validation failed: deviceId or timestamp is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrMissingPhotoInvalidFields.Error()})
		return
	}

	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		log.Printf("Invalid timestamp format: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrMissingPhotoInvalidFields.Error()})
		return
	}

	input := domain.PhotoDto{
		DeviceID:  deviceID,
		Timestamp: timestamp,
	}

	file, _, err := c.Request.FormFile("photo")
	if err != nil {
		log.Printf("Photo file error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrPhotoData.Error()})
		return
	}
	defer file.Close()

	photoBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Failed to read photo file: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read photo file"})
		return
	}
	if len(photoBytes) == 0 {
		log.Println("Photo file is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrPhotoData.Error()})
		return
	}
	if len(photoBytes) > 5*1024*1024 {
		log.Println("Photo size exceeds 5MB")
		c.JSON(http.StatusBadRequest, gin.H{"error": "photo size exceeds 5MB"})
		return
	}

	log.Printf("Valid input: deviceId=%s, timestamp=%d, photo size=%d bytes\n",
		input.DeviceID, input.Timestamp, len(photoBytes))

	photo, err := h.CreatePhotoUseCase.Execute(input, photoBytes)
	if err != nil {
		log.Printf("Use case failed: %v\n", err)
		if errors.Is(err, domain.ErrDeviceIDPhoto) || errors.Is(err, domain.ErrTimestampPhoto) || errors.Is(err, domain.ErrPhotoData) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, photo)
}
