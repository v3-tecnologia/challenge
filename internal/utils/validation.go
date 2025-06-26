package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// TranslateValidationErrors converte erros de validação do validator.v10
// em um mapa de mensagens amigáveis para cada campo.
//
// Recebe um erro retornado pelo método de binding do Gin (que usa validator.v10 internamente)
// e retorna um map[string]string, onde a chave é o nome do campo e o valor é a mensagem traduzida.
//
// Suporta as tags "required", "email" e "min". Outros erros são retornados como "Valor inválido".
//
// Exemplo de uso:
//
//	err := c.ShouldBindJSON(&req)
//	if err != nil {
//	    errs := TranslateValidationErrors(err)
//	    c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
//	    return
//	}
func TranslateValidationErrors(err error) []string {
	var errors []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := strings.ToLower(fieldError.Field())
			switch fieldError.Tag() {
			case "required":
				errors = append(errors, fmt.Sprintf("%s é obrigatório", field))
			case "email":
				errors = append(errors, fmt.Sprintf("%s deve ser um email válido", field))
			case "min":
				errors = append(errors, fmt.Sprintf("%s deve ter pelo menos %s caracteres", field, fieldError.Param()))
			case "max":
				errors = append(errors, fmt.Sprintf("%s deve ter no máximo %s caracteres", field, fieldError.Param()))
			default:
				errors = append(errors, fmt.Sprintf("%s é inválido", field))
			}
		}
	} else {
		errors = append(errors, err.Error())
	}

	return errors
}
