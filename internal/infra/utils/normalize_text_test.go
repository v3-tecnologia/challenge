package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/infra/utils"
)

func TestNormalizeText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World!", "hello_world"},
		{"Go is awesome!", "go_is_awesome"},
		{"12345", "12345"},
		{"", ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := utils.NormalizeText(test.input)
			require.Equal(t, test.expected, result)
		})
	}
}
