package utils_test

import (
	"testing"

	"github.com/KaiRibeiro/challenge/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	MAC     string `validate:"mac"`
	Image   string `validate:"base64"`
	Name    string `validate:"required"`
	Invalid string `validate:"numeric"`
}

func TestFormatFieldError(t *testing.T) {
	validate := validator.New()

	test := TestStruct{}

	err := validate.Struct(test)
	assert.Error(t, err)

	validationErrors := err.(validator.ValidationErrors)

	for _, fe := range validationErrors {
		output := utils.FormatFieldError(fe)
		switch fe.Field() {
		case "MAC":
			assert.Equal(t, "MAC must be a valid MAC address", output)
		case "Image":
			assert.Equal(t, "Image must be a valid base64-encoded string", output)
		case "Name":
			assert.Equal(t, "Name is required", output)
		case "Invalid":
			assert.Equal(t, "Invalid is invalid", output)
		default:
			t.Errorf("unexpected field: %s", fe.Field())
		}
	}
}
