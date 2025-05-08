package common

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       any
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func Validate(data any) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.Error = true
			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
