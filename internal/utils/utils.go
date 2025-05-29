package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatFieldError(fe validator.FieldError) string {
	field := fe.Field()

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "mac":
		return fmt.Sprintf("%s must be a valid MAC address", field)
	case "base64":
		return fmt.Sprintf("%s must be a valid base64-encoded string", field)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
