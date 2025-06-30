package main

import "testing"

func TestInsertGps(t *testing.T) {
	ctx, err := NewAppContext()
	if err != nil {
		t.Fatalf("error creating context: %v", err)
	}

	gpsData := GpsData{}
	_, err = ctx.insertGps(gpsData)
	if err == nil {
		t.Error("inserted empty GpsData")
	}

	gpsData.UniqueId = "1"
	gpsData.Timestamp = "1"
	gpsData.Latitude = 1.0
	gpsData.Longitude = 1.0
	_, err = ctx.insertGps(gpsData)
	if err != nil {
		t.Error("could not insert GpsData")
	}
}
