package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/igorlopes88/desafio-v3/internal/dto"
	"github.com/igorlopes88/desafio-v3/internal/entity"
	"github.com/igorlopes88/desafio-v3/internal/infra/database"
)

type PhotoHandler struct {
	PhotoDB database.PhotoInterface
}

func NewPhotoHandler(db database.PhotoInterface) *PhotoHandler {
	return &PhotoHandler{
		PhotoDB: db,
	}
}


func (p *PhotoHandler) Register(w http.ResponseWriter, r *http.Request) {
	var photo dto.RegisterPhotoInput
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	g, err := entity.NewPhoto(photo.UserId, photo.MacAddress, photo.Image, photo.TimeStamp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = p.PhotoDB.Register(g)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
