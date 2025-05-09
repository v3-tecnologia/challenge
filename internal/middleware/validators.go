package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       any
}

var structValidator = validator.New(validator.WithRequiredStructEnabled())

func validate(data any) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := structValidator.Struct(data)
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

func ValidateStruct[T any](c *fiber.Ctx) error {
	var params T

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	validationErrors := validate(params)
	if len(validationErrors) > 0 {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"errors": validationErrors})
	}

	// Store the parsed and validated data in Locals to be used in the handler
	c.Locals("validatedBody", params)

	return c.Next()
}
