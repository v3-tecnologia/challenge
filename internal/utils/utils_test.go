package utils

import (
	"errors"
	"testing"

	"github.com/Kairum-Labs/should"
	"github.com/go-playground/validator/v10"
)

type Dummy struct {
	Name  string `validate:"required,min=3,max=10"`
	Email string `validate:"required,email"`
}

const WorkingGIF = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII="

func TestIsBase64(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", false},
		{"invalid", false},
		{WorkingGIF, true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			should.BeEqual(t, IsValidImageBase64(test.input), test.expected)
		})
	}
}

func TestTranslateValidationErrors(t *testing.T) {
	validate := validator.New()

	t.Run("erros de validação conhecidos", func(t *testing.T) {
		d := Dummy{
			Name:  "Jo",             // min falha
			Email: "invalido-email", // email falha
		}

		err := validate.Struct(d)
		errs := TranslateValidationErrors(err)

		should.Contain(t, errs, "name deve ter pelo menos 3 caracteres")
		should.Contain(t, errs, "email deve ser um email válido")
	})

	t.Run("campo ausente (required)", func(t *testing.T) {
		d := Dummy{}
		err := validate.Struct(d)
		errs := TranslateValidationErrors(err)

		should.Contain(t, errs, "name é obrigatório")
		should.Contain(t, errs, "email é obrigatório")
	})

	t.Run("campo maior que o máximo", func(t *testing.T) {
		d := Dummy{
			Name:  "12345678901", // 11 caracteres (max = 10)
			Email: "a@a.com",
		}
		err := validate.Struct(d)
		errs := TranslateValidationErrors(err)

		should.Contain(t, errs, "name deve ter no máximo 10 caracteres")
	})

	t.Run("erro genérico", func(t *testing.T) {
		err := errors.New("erro qualquer")
		errs := TranslateValidationErrors(err)

		should.BeEqual(t, []string{"erro qualquer"}, errs)
	})
}

func TestParseTimestamp(t *testing.T) {
	t.Run("formatos válidos", func(t *testing.T) {
		tests := []struct {
			name      string
			input     string
			expectErr bool
		}{
			{
				name:      "RFC3339",
				input:     "2023-10-15T14:30:45Z",
				expectErr: false,
			},
			{
				name:      "RFC3339Nano",
				input:     "2023-10-15T14:30:45.123456789Z",
				expectErr: false,
			},
			{
				name:      "ISO8601 with timezone",
				input:     "2023-10-15T14:30:45-03:00",
				expectErr: false,
			},
			{
				name:      "ISO8601 with microseconds and timezone",
				input:     "2023-10-15T14:30:45.123456-03:00",
				expectErr: false,
			},
			{
				name:      "Simple datetime",
				input:     "2023-10-15 14:30:45",
				expectErr: false,
			},
			{
				name:      "Simple datetime with microseconds",
				input:     "2023-10-15 14:30:45.123456",
				expectErr: false,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseTimestamp(test.input)

				if test.expectErr {
					should.BeNotNil(t, err)
				} else {
					should.BeNil(t, err)
					should.BeFalse(t, result.IsZero())
				}
			})
		}
	})

	t.Run("formatos inválidos", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{
				name:  "string vazia",
				input: "",
			},
			{
				name:  "formato inválido",
				input: "invalid-timestamp",
			},
			{
				name:  "apenas data",
				input: "2023-10-15",
			},
			{
				name:  "apenas hora",
				input: "14:30:45",
			},
			{
				name:  "formato incorreto",
				input: "15/10/2023 14:30:45",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseTimestamp(test.input)

				should.BeNotNil(t, err)
				should.BeEqual(t, err, ErrInvalidTimestampFormat)
				should.BeTrue(t, result.IsZero())
			})
		}
	})

	t.Run("consistência de parsing", func(t *testing.T) {
		timestamp1, err1 := ParseTimestamp("2023-10-15T14:30:45Z")
		timestamp2, err2 := ParseTimestamp("2023-10-15T14:30:45.000000Z")

		should.BeNil(t, err1)
		should.BeNil(t, err2)
		should.BeEqual(t, timestamp1.Unix(), timestamp2.Unix())
	})

	t.Run("timezone handling", func(t *testing.T) {
		// Teste para verificar se timestamps com timezone são processados corretamente
		utc, err1 := ParseTimestamp("2023-10-15T14:30:45Z")
		withTimezone, err2 := ParseTimestamp("2023-10-15T14:30:45+05:00")

		should.BeNil(t, err1)
		should.BeNil(t, err2)
		should.BeFalse(t, utc.IsZero())
		should.BeFalse(t, withTimezone.IsZero())
	})
}
