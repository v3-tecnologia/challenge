package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int
	Code       string
	Message    string
	Details    string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Code, e.Message, e.Details)
}

func (e *AppError) GetStatusCode() int {
	return e.StatusCode
}

func (e *AppError) ToHTTPResponse() map[string]interface{} {
	return map[string]interface{}{
		"error": map[string]interface{}{
			"code":    e.Code,
			"message": e.Message,
			"details": e.Details,
		},
	}
}

// (400)
func NewErrorBadRequest(code, details string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Code:       code,
		Message:    "Invalid Request",
		Details:    details,
	}
}

// (404)
func NewErrorNotFound(code, details string) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Code:       code,
		Message:    "Resource Not Found",
		Details:    details,
	}
}

// (500)
func NewErrorInternal(code, details string) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Code:       code,
		Message:    "Internal Server Error",
		Details:    details,
	}
}
