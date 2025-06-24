package http

import (
	"encoding/json"
	"net/http"

	"github.com/dryingcore/v3-challenge/internal/core/model"
	"github.com/dryingcore/v3-challenge/internal/core/usecase"
)

type TelemetryHandler struct {
	GyroscopeUC usecase.GyroscopeUC
	GPSUC       usecase.GPSUC
}

func NewTelemetryHandler(
	g usecase.GyroscopeUC,
	gps usecase.GPSUC,
) *TelemetryHandler {
	return &TelemetryHandler{
		GyroscopeUC: g,
		GPSUC:       gps,
	}
}

func (h *TelemetryHandler) HandleGyroscope(w http.ResponseWriter, r *http.Request) {
	var req model.Gyroscope

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid JSON payload"}`, http.StatusBadRequest)
		return
	}

	if err := h.GyroscopeUC.Register(req); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *TelemetryHandler) HandleGPS(w http.ResponseWriter, r *http.Request) {
	var req model.GPSData

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid JSON payload"}`, http.StatusBadRequest)
		return
	}

	if err := h.GPSUC.Register(req); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
