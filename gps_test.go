package main

import (
	"testing"
)

func TestBuildGpsSchema(t *testing.T) {
	ctx, err := NewAppContext()
	if err != nil {
		t.Errorf("error creating context: %v", err)
	}
	if ctx.GpsSchema == nil {
		t.Error("GpsSchema was not built")
	}
}
