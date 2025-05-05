package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mkafonso/go-cloud-challenge/api/rest"
	"github.com/mkafonso/go-cloud-challenge/usecase"
	appError "github.com/mkafonso/go-cloud-challenge/usecase/errors"
)

type SavePhotoRequest struct {
	DeviceID  string `json:"device_id"`
	FilePath  string `json:"file_path"`
	Timestamp string `json:"timestamp"`
}

type SavePhotoHandler struct {
	usecase *usecase.SavePhoto
}

func NewSavePhotoHandler(uc *usecase.SavePhoto) *SavePhotoHandler {
	return &SavePhotoHandler{usecase: uc}
}

func (h *SavePhotoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var body SavePhotoRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		rest.HandleErrorJSON(w, appError.NewErrorBadRequest(
			"invalid request body",
			"please send a valid JSON payload",
		))
		return
	}

	req := &usecase.SavePhotoRequest{
		DeviceID:  body.DeviceID,
		FilePath:  body.FilePath,
		Timestamp: body.Timestamp,
	}

	res, err := h.usecase.Execute(r.Context(), req)
	if err != nil {
		rest.HandleErrorJSON(w, err)
		return
	}

	rest.WriteJSON(w, http.StatusCreated, rest.JSONResponse{
		Error:      false,
		StatusCode: http.StatusCreated,
		Message:    "photo has been saved",
		Data: map[string]any{
			"recognized": res.Recognized,
		},
	})
}
