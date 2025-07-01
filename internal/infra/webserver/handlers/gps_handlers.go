package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/igorlopes88/desafio-v3/internal/dto"
	"github.com/igorlopes88/desafio-v3/internal/entity"
	"github.com/igorlopes88/desafio-v3/internal/infra/database"
)

type GpsHandler struct {
	GpsDB database.GpsInterface
}

func NewGpsHandler(db database.GpsInterface) *GpsHandler {
	return &GpsHandler{
		GpsDB: db,
	}
}

// Register GPS data
// @Summary Register GPS
// @Description Record GPS data
// @Tags telemetry
// @Accept json
// @Produce json
// @Param request body dto.RegisterGpsInput true "gps cordenate"
// @Success 201
// @Failure 400 {object} Error
// @Router /telemetry/gps [post]
// @Security ApiKeyAuth
func (h *GpsHandler) Register(w http.ResponseWriter, r *http.Request) {
	var gps dto.RegisterGpsInput
	err := json.NewDecoder(r.Body).Decode(&gps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	g, err := entity.NewGps(gps.UserId, gps.MacAddress, gps.Latitude, gps.Longitude, gps.TimeStamp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.GpsDB.Register(g)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
