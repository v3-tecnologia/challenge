package controller

import (
	"challenge-cloud/internal/models"
	service "challenge-cloud/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type GyroscopeController struct {
	Service *service.GyroscopeService
}

func NewGyroscopeController(s *service.GyroscopeService) *GyroscopeController {
	return &GyroscopeController{Service: s}
}

func (c *GyroscopeController) CreateGyroscope(w http.ResponseWriter, r *http.Request) {
	var data models.Gyroscope
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if data.MAC == "" {
		http.Error(w, "MAC is required", http.StatusBadRequest)
		return
	}
	if data.X == 0 {
		http.Error(w, "Axis X is required", http.StatusBadRequest)
		return
	}
	if data.Y == 0 {
		http.Error(w, "Axis Y is required", http.StatusBadRequest)
		return
	}
	if data.Z == 0 {
		http.Error(w, "Axis Z is required", http.StatusBadRequest)
		return
	}

	if err := c.Service.Save(&data); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func (c *GyroscopeController) GetGyroscope(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
		fmt.Println("page invalid, n 1")
	}
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		sizeInt = 10
		fmt.Println("limit invalid, n 10")
	}

	gyro, err := c.Service.GetAll(pageInt, sizeInt)
	if err != nil {
		http.Error(w, "gyroscope not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gyro)
}
