package services_test

import (
	"os"
	"testing"

	"github.com/KaiRibeiro/challenge/internal/logs"
)

func TestMain(m *testing.M) {
	logs.Init()
	os.Exit(m.Run())
}
