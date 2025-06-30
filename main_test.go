package main

import "testing"

func TestOpenDatabase(t *testing.T) {
	ctx, err := NewAppContext()
	if err != nil {
		t.Error(err)
	}
	if ctx.DB == nil {
		t.Error("database could not be opened")
	}
}

func TestBuildSchemas(t *testing.T) {
	ctx, err := NewAppContext()
	if err != nil {
		t.Fatal(err)
	}

	if ctx.GyroSchema == nil {
		t.Error("gyro schema could not be created")
	}

	if ctx.GpsSchema == nil {
		t.Error("gps schema could not be created")
	}

	if ctx.PhotoSchema == nil {
		t.Error("photo schema could not be created")
	}
}
