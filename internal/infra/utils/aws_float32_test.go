package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/infra/utils"
)

func TestAwsFloat32(t *testing.T) {
	value := float32(1.23)
	result := utils.AwsFloat32(value)

	require.Equal(t, value, *result, "Expected the float32 value to be equal")
	// Check if the pointer is not nil
	require.NotNil(t, result, "Expected the pointer to be not nil")
}
