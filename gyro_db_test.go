package main

import "testing"

func TestInsertGyro(t *testing.T) {
	ctx, err := NewAppContext()
	if err != nil {
		t.Fatal(err)
	}

	gyroData := GyroData{}
	_, err = ctx.insertGyro(gyroData)
	if err == nil {
		t.Error("inserted empty GyroData")
	}

	gyroData.UniqueId = "1"
	gyroData.Timestamp = "1"
	gyroData.X = 1.0
	gyroData.Y = 1.0
	gyroData.Z = 1.0
	_, err = ctx.insertGyro(gyroData)
	if err != nil {
		t.Error("could not insert GyroData")
	}
}
