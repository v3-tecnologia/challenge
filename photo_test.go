package main

import (
	"testing"
)

func TestBuildPhotoSchema(t *testing.T) {
	ctx, err := NewAppContext()
	if err != nil {
		t.Error(err)
	}
	if ctx.PhotoSchema == nil {
		t.Error("PhotoSchema was not built")
	}
}
