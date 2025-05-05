package rest

import (
	"net/http"

	appError "github.com/mkafonso/go-cloud-challenge/usecase/errors"
)

func HandleErrorJSON(w http.ResponseWriter, err error) error {
	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	switch customErr := err.(type) {

	case *appError.ErrorBadRequest:
		payload.Action = customErr.GetAction()
		payload.StatusCode = http.StatusBadRequest
		payload.Name = "ErrorBadRequest"
		WriteJSON(w, http.StatusBadRequest, payload)

	default:
		payload.Message = "internal server error"
		payload.Action = "something went wrong on our end, please try again later"
		payload.StatusCode = http.StatusInternalServerError
		payload.Name = "ErrorInternalServerError"
		WriteJSON(w, http.StatusInternalServerError, payload)
	}

	return nil
}
