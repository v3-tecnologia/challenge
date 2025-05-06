package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mkafonso/go-cloud-challenge/api/rest"
	"github.com/mkafonso/go-cloud-challenge/usecase"
	appError "github.com/mkafonso/go-cloud-challenge/usecase/errors"
)

type SaveGPSDataRequest struct {
	DeviceID  string   `json:"device_id"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Timestamp string   `json:"timestamp"`
}

type SaveGPSDataHandler struct {
	usecase *usecase.SaveGPSData
}

func NewSaveGPSDataHandler(uc *usecase.SaveGPSData) *SaveGPSDataHandler {
	return &SaveGPSDataHandler{usecase: uc}
}

func (h *SaveGPSDataHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var body SaveGPSDataRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		rest.HandleErrorJSON(w, appError.NewErrorBadRequest(
			"invalid request body",
			"please send a valid JSON payload",
		))
		return
	}

	req := &usecase.SaveGPSDataRequest{
		DeviceID:  body.DeviceID,
		Latitude:  body.Latitude,
		Longitude: body.Longitude,
		Timestamp: body.Timestamp,
	}

	_, err := h.usecase.Execute(r.Context(), req)
	if err != nil {
		rest.HandleErrorJSON(w, err)
		return
	}

	rest.WriteJSON(w, http.StatusCreated, rest.JSONResponse{
		Error:      false,
		StatusCode: http.StatusCreated,
		Message:    "gps data has been saved",
	})
}
