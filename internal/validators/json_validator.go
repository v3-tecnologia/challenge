package validators

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate(ctx *gin.Context, dto interface{}) bool {
	if err := ctx.ShouldBindJSON(dto); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			var messages []string
			for _, fe := range ve {
				field := strings.ToLower(fe.Field())
				switch fe.Tag() {
				case "required":
					messages = append(messages, fmt.Sprintf("The field '%s' is required.", field))
				case "gt":
					messages = append(messages, fmt.Sprintf("The field '%s' must be greater than %s.", field, fe.Param()))
				case "lt":
					messages = append(messages, fmt.Sprintf("The field '%s' must be less than %s.", field, fe.Param()))
				case "gte":
					messages = append(messages, fmt.Sprintf("The field '%s' must be greater than or equal to %s.", field, fe.Param()))
				case "lte":
					messages = append(messages, fmt.Sprintf("The field '%s' must be less than or equal to %s.", field, fe.Param()))
				default:
					messages = append(messages, fmt.Sprintf("The field '%s' is invalid.", field))
				}
			}

			ctx.JSON(400, gin.H{"errors": messages})
			return false
		}

		// Handle type error (e.g., string into float64)
		var unmarshalTypeError *json.UnmarshalTypeError
		if errors.As(err, &unmarshalTypeError) {
			field := strings.ToLower(unmarshalTypeError.Field)
			ctx.JSON(400, gin.H{
				"error": fmt.Sprintf("The field '%s' has an invalid type. Expected: %s.", field, unmarshalTypeError.Type.String()),
			})
			return false
		}

		ctx.JSON(400, gin.H{"error": err.Error()})
		return false
	}
	return true
}
