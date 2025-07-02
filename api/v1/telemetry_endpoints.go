package v1

import (
	"net/http"
	"time"

	"go-challenge/internal/models"
	"go-challenge/internal/services"
	"go-challenge/internal/messaging"

	"github.com/gin-gonic/gin"
)

type TelemetryHandler struct {
	Service      *services.TelemetryService
	NATSProducer *messaging.NATSProducer
}

func NewTelemetryHandler(service *services.TelemetryService, natsProducer *messaging.NATSProducer) *TelemetryHandler {
	return &TelemetryHandler{Service: service, NATSProducer: natsProducer}
}

// PostGyroscopeHandler godoc
// @Summary Envia dados do giroscópio
// @Description Recebe dados do giroscópio e os salva no banco de dados
// @Tags Telemetry
// @Accept json
// @Produce json
// @Param gyroscope body models.GyroscopeDTO true "Dados do Giroscópio"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /telemetry/gyroscope [post]
func (h *TelemetryHandler) PostGyroscopeHandler(c *gin.Context) {
	var dto models.GyroscopeDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Converte o DTO para o modelo
	data := models.Gyroscope{
		X: dto.X,
		Y: dto.Y,
		Z: dto.Z,
	}

	if err := h.Service.SaveGyroscopeData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save gyroscope data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gyroscope data saved successfully"})
}

// PostGPSHandler godoc
// @Summary Envia dados do GPS
// @Description Recebe dados do GPS e os salva no banco de dados
// @Tags Telemetry
// @Accept json
// @Produce json
// @Param gps body models.GPSDTO true "Dados do GPS"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /telemetry/gps [post]
func (h *TelemetryHandler) PostGPSHandler(c *gin.Context) {
	var dto models.GPSDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Converte o DTO para o modelo
	data := models.GPS{
		Latitude:  dto.Latitude,
		Longitude: dto.Longitude,
	}

	if err := h.Service.SaveGPSData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save GPS data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GPS data saved successfully"})
}

// PostPhotoHandler godoc
// @Summary Envia dados de uma foto
// @Description Recebe dados de uma foto e os salva no banco de dados
// @Tags Telemetry
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Arquivo da Foto"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /telemetry/photo [post]
func (h *TelemetryHandler) PostPhotoHandler(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open image file"})
		return
	}
	defer f.Close()

	imageData := make([]byte, file.Size)
	_, err = f.Read(imageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image file"})
		return
	}

	data := models.Photo{
		Image: imageData,
	}

	if err := h.Service.SavePhotoData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo data"})
		return
	}

	// Publica no NATS
	photoMsg := messaging.PhotoMessage{
		Photo:     imageData,
		DeviceID:  file.Filename, // ou gere um UUID se preferir
		Timestamp: time.Now().Unix(),
	}
	if err := h.NATSProducer.PublishPhoto(photoMsg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish photo to queue"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo data saved and published successfully"})
}
