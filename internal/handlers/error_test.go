package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bielgennaro/v3-challenge-cloud/internal/errors"
)

func TestHandleError_AppError_BadRequest(t *testing.T) {
	appError := errors.NewErrorBadRequest("invalid_data", "The provided data is invalid")

	rr := httptest.NewRecorder()

	HandleError(rr, appError)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("esperado status %d, mas recebeu %d", http.StatusBadRequest, status)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("erro ao decodificar a resposta: %v", err)
	}

	if response["error"].(map[string]interface{})["code"] != "invalid_data" {
		t.Errorf("esperado código de erro 'invalid_data', mas obteve %v", response["error"].(map[string]interface{})["code"])
	}
	if response["error"].(map[string]interface{})["message"] != "Invalid Request" {
		t.Errorf("esperado mensagem 'Invalid Request', mas obteve %v", response["error"].(map[string]interface{})["message"])
	}
	if response["error"].(map[string]interface{})["details"] != "The provided data is invalid" {
		t.Errorf("esperado detalhes 'The provided data is invalid', mas obteve %v", response["error"].(map[string]interface{})["details"])
	}
}

func TestHandleError_AppError_NotFound(t *testing.T) {
	appError := errors.NewErrorNotFound("resource_not_found", "The requested resource could not be found")

	rr := httptest.NewRecorder()

	HandleError(rr, appError)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("esperado status %d, mas recebeu %d", http.StatusNotFound, status)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("erro ao decodificar a resposta: %v", err)
	}

	if response["error"].(map[string]interface{})["code"] != "resource_not_found" {
		t.Errorf("esperado código de erro 'resource_not_found', mas obteve %v", response["error"].(map[string]interface{})["code"])
	}

	if response["error"].(map[string]interface{})["message"] != "Resource Not Found" {
		t.Errorf("esperado mensagem 'Resource Not Found', mas obteve %v", response["error"].(map[string]interface{})["message"])
	}

	if response["error"].(map[string]interface{})["details"] != "The requested resource could not be found" {
		t.Errorf("esperado detalhes 'The requested resource could not be found', mas obteve %v", response["error"].(map[string]interface{})["details"])
	}
}

func TestHandleError_AppError_Internal(t *testing.T) {
	appError := errors.NewErrorInternal("server_error", "An internal error occurred on the server")

	rr := httptest.NewRecorder()

	HandleError(rr, appError)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("esperado status %d, mas recebeu %d", http.StatusInternalServerError, status)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("erro ao decodificar a resposta: %v", err)
	}

	if response["error"].(map[string]interface{})["code"] != "server_error" {
		t.Errorf("esperado código de erro 'server_error', mas obteve %v", response["error"].(map[string]interface{})["code"])
	}

	if response["error"].(map[string]interface{})["message"] != "Internal Server Error" {
		t.Errorf("esperado mensagem 'Internal Server Error', mas obteve %v", response["error"].(map[string]interface{})["message"])
	}

	if response["error"].(map[string]interface{})["details"] != "An internal error occurred on the server" {
		t.Errorf("esperado detalhes 'An internal error occurred on the server', mas obteve %v", response["error"].(map[string]interface{})["details"])
	}
}
