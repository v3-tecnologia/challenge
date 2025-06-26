package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"

	"telemetry-api/internal/dtos/requests"
	"telemetry-api/internal/services"
	"telemetry-api/internal/utils"
)

type TelemetryHandler struct {
	service *services.TelemetryService
	nc      *nats.Conn
	logger  *zap.Logger
}

func NewTelemetryHandler(service *services.TelemetryService, nc *nats.Conn, logger *zap.Logger) *TelemetryHandler {
	return &TelemetryHandler{service: service, nc: nc, logger: logger}
}

func (h *TelemetryHandler) CreateGyroscope(c *gin.Context) {
	var gyroscope requests.CreateGyroscopeRequest
	if err := c.ShouldBindJSON(&gyroscope); err != nil {
		h.logger.Error("failed to bind JSON", zap.Error(err))
		errs := utils.TranslateValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	gyroscopeJSON, err := json.Marshal(gyroscope)
	if err != nil {
		h.logger.Error("failed to marshal gyroscope data", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao processar dados do giroscópio"})
		return
	}

	if h.nc != nil {
		if err := h.nc.Publish("telemetry.gyroscope", gyroscopeJSON); err != nil {
			h.logger.Error("failed to publish gyroscope data to NATS", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao enfileirar dados do giroscópio"})
			return
		}
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "dados do giroscópio recebidos para processamento"})
}

func (h *TelemetryHandler) CreateGPS(c *gin.Context) {
	var gps requests.CreateGPSRequest
	if err := c.ShouldBindJSON(&gps); err != nil {
		h.logger.Error("failed to bind JSON", zap.Error(err))
		errs := utils.TranslateValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	gpsJSON, err := json.Marshal(gps)
	if err != nil {
		h.logger.Error("failed to marshal gps data", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao processar dados do GPS"})
		return
	}

	if h.nc != nil {
		if err := h.nc.Publish("telemetry.gps", gpsJSON); err != nil {
			h.logger.Error("failed to publish gps data to NATS", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao enfileirar dados do GPS"})
			return
		}
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "dados do GPS recebidos para processamento"})
}

func (h *TelemetryHandler) CreateTelemetryPhoto(c *gin.Context) {
	var telemetryPhoto requests.CreateTelemetryPhotoRequest
	if err := c.ShouldBindJSON(&telemetryPhoto); err != nil {
		h.logger.Error("failed to bind JSON", zap.Error(err))
		errs := utils.TranslateValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if !utils.IsValidImageBase64(telemetryPhoto.Photo) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "formato de imagem base64 inválido"})
		return
	}

	photoJSON, err := json.Marshal(telemetryPhoto)
	if err != nil {
		h.logger.Error("failed to marshal photo data", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao processar foto"})
		return
	}

	if h.nc != nil {
		if err := h.nc.Publish("telemetry.photo", photoJSON); err != nil {
			h.logger.Error("failed to publish photo data to NATS", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao enfileirar foto"})
			return
		}
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "foto recebida para processamento"})
}

func (h *TelemetryHandler) GetGyroscopeData(c *gin.Context) {
	deviceID := c.Query("device_id")
	page, limit := h.getPaginationParams(c)

	data, err := h.service.GetGyroscopeData(deviceID, page, limit)
	if err != nil {
		h.logger.Error("failed to get gyroscope data", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao buscar dados do giroscópio"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *TelemetryHandler) GetGPSData(c *gin.Context) {
	deviceID := c.Query("device_id")
	page, limit := h.getPaginationParams(c)

	data, err := h.service.GetGPSData(deviceID, page, limit)
	if err != nil {
		h.logger.Error("failed to get GPS data", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao buscar dados do GPS"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *TelemetryHandler) GetPhotoData(c *gin.Context) {
	deviceID := c.Query("device_id")
	page, limit := h.getPaginationParams(c)

	data, err := h.service.GetPhotoData(deviceID, page, limit)
	if err != nil {
		h.logger.Error("failed to get photo data", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao buscar dados de fotos"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *TelemetryHandler) GetDevices(c *gin.Context) {
	page, limit := h.getPaginationParams(c)

	data, err := h.service.GetDevices(page, limit)
	if err != nil {
		h.logger.Error("failed to get devices", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao buscar dispositivos"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *TelemetryHandler) getPaginationParams(c *gin.Context) (page, limit int) {
	page = 1
	limit = 10

	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	return page, limit
}
