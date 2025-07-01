package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/igorlopes88/desafio-v3/internal/dto"
	"github.com/igorlopes88/desafio-v3/internal/entity"
	"github.com/igorlopes88/desafio-v3/internal/infra/database"
)

type GyroscopeHandler struct {
	GyroscopeDB database.GyroscopeInterface
}

func NewGyroscopeHandler(db database.GyroscopeInterface) *GyroscopeHandler {
	return &GyroscopeHandler{
		GyroscopeDB: db,
	}
}

// Register Gyroscope data
// @Summary Register gyroscope
// @Description Record Gyroscope data
// @Tags telemetry
// @Accept json
// @Produce json
// @Param request body dto.RegisterGyroscopeInput true "gyroscope cordenate"
// @Success 201
// @Failure 400 {object} Error
// @Router /telemetry/gyroscope [post]
// @Security ApiKeyAuth
func (h *GyroscopeHandler) Register(w http.ResponseWriter, r *http.Request) {
	var gyro dto.RegisterGyroscopeInput
	err := json.NewDecoder(r.Body).Decode(&gyro)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	g, err := entity.NewGyroscope(gyro.UserId, gyro.MacAddress, gyro.XAxis, gyro.YAxis, gyro.ZAxis, gyro.TimeStamp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.GyroscopeDB.Register(g)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
