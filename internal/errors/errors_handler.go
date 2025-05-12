package errors

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewErrorResponse(status int, message string, details interface{}) ErrorResponse {
	return ErrorResponse{
		Status:  status,
		Message: message,
		Details: details,
	}
}

func NewValidationErrorResponse(validationErrors []ValidationError) ErrorResponse {
	return ErrorResponse{
		Status:  400,
		Message: "Falha na validação de dados",
		Details: validationErrors,
	}
}

func HandleValidationErrors(err error) ErrorResponse {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var details []ValidationError

		for _, e := range validationErrors {
			field := strings.ToLower(e.Field())
			message := getErrorMessage(e)

			details = append(details, ValidationError{
				Field:   field,
				Message: message,
			})
		}

		return NewValidationErrorResponse(details)
	}

	return NewErrorResponse(
		http.StatusBadRequest,
		"Formato de dados inválido",
		err.Error(),
	)
}

func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "Este campo é obrigatório"
	case "email":
		return "Endereço de e-mail inválido"
	case "min":
		return "Valor abaixo do mínimo permitido"
	case "max":
		return "Valor acima do máximo permitido"
	default:
		return "Valor inválido"
	}
}

func RespondWithError(c *gin.Context, statusCode int, message string, details interface{}) {
	errResponse := NewErrorResponse(statusCode, message, details)
	c.JSON(statusCode, errResponse)
}
