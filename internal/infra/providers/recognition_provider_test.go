package providers_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/infra/providers"
)

func TestRecognitionProvider_CompareFaces(t *testing.T) {
	recognitionProvider := providers.NewRecognitionProviderInMemory()

	matched, err := recognitionProvider.CompareFaces(
		[]byte("test"),
		[]byte("test"),
	)

	require.NoError(t, err)
	require.True(t, matched)

	matched, err = recognitionProvider.CompareFaces(
		[]byte("test"),
		[]byte("test2"),
	)
	require.NoError(t, err)
	require.False(t, matched)
}
