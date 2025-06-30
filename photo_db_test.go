package main

import "testing"

func TestInsertPhoto(t *testing.T) {
	ctx, err := NewAppContext()
	if err != nil {
		t.Fatal(err)
	}
	photoData := PhotoData{}
	_, err = ctx.insertPhoto(photoData)
	if err == nil {
		t.Error("inserted empty PhotoData")
	}

	photoData.UniqueId = "1"
	photoData.Timestamp = "1"
	photoData.Photo = "1"
	_, err = ctx.insertPhoto(photoData)
	if err != nil {
		t.Error("could not insert PhotoData")
	}
}
