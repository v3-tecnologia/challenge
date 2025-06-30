package main

import (
	"testing"
)

func TestBuildGyroSchema(t *testing.T) {
	ctx, err := NewAppContext()
	if err != nil {
		t.Fatal(err)
	}
	if ctx.GyroSchema == nil {
		t.Error("GyroSchema was not built")
	}
}
