package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mkafonso/go-cloud-challenge/api/rest"
	"github.com/mkafonso/go-cloud-challenge/repository"
	"github.com/mkafonso/go-cloud-challenge/usecase"
	appError "github.com/mkafonso/go-cloud-challenge/usecase/errors"
	factory "github.com/mkafonso/go-cloud-challenge/usecase/factories"
)

type SaveGyroscopeDataRequest struct {
	DeviceID  string   `json:"device_id"`
	X         *float64 `json:"x"`
	Y         *float64 `json:"y"`
	Z         *float64 `json:"z"`
	Timestamp string   `json:"timestamp"`
}

type SaveGyroscopeDataHandler struct {
	usecase *usecase.SaveGyroscopeData
}

func NewSaveGyroscopeDataHandler(gyroRepo repository.GyroscopeRepositoryInterface) *SaveGyroscopeDataHandler {
	usecase := factory.SaveGyroscopeDataFactory(gyroRepo)
	return &SaveGyroscopeDataHandler{usecase: usecase}
}

func (h *SaveGyroscopeDataHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var body SaveGyroscopeDataRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		rest.HandleErrorJSON(w, appError.NewErrorBadRequest(
			"invalid request body",
			"please send a valid JSON payload",
		))
		return
	}

	req := &usecase.SaveGyroscopeDataRequest{
		DeviceID:  body.DeviceID,
		X:         body.X,
		Y:         body.Y,
		Z:         body.Z,
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
		Message:    "gyroscope data has been saved",
	})
}
