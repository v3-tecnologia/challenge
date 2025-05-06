package controller

import (
	"io"
	"net/http"

	"github.com/mkafonso/go-cloud-challenge/api/rest"
	"github.com/mkafonso/go-cloud-challenge/recognition"
	"github.com/mkafonso/go-cloud-challenge/repository"
	"github.com/mkafonso/go-cloud-challenge/storage"
	"github.com/mkafonso/go-cloud-challenge/usecase"
	appError "github.com/mkafonso/go-cloud-challenge/usecase/errors"
	factory "github.com/mkafonso/go-cloud-challenge/usecase/factories"
)

type SavePhotoHandler struct {
	usecase *usecase.SavePhoto
}

func NewSavePhotoHandler(
	repo repository.PhotoRepositoryInterface,
	recognizer recognition.FaceRecognitionService,
	storage storage.PhotoStorageService,
) *SavePhotoHandler {
	usecase := factory.SavePhotoDataFactory(repo, recognizer, storage)
	return &SavePhotoHandler{usecase: usecase}
}

func (h *SavePhotoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB max memory
	if err != nil {
		rest.HandleErrorJSON(w, appError.NewErrorBadRequest(
			"invalid multipart data",
			"please upload the file as multipart/form-data",
		))
		return
	}

	deviceID := r.FormValue("device_id")
	timestamp := r.FormValue("timestamp")

	file, _, err := r.FormFile("photo")
	if err != nil {
		rest.HandleErrorJSON(w, appError.NewErrorBadRequest(
			"photo file is required",
			"please upload an image with the field name 'photo'",
		))
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		rest.HandleErrorJSON(w, appError.NewErrorInternalServer(
			"could not read uploaded file",
			"try again or upload a different file",
		))
		return
	}

	req := &usecase.SavePhotoRequest{
		DeviceID:  deviceID,
		FileBytes: fileBytes,
		Timestamp: timestamp,
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
