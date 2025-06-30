package controller

import (
	"challenge-cloud/internal/models"
	service "challenge-cloud/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type GyroscopeController struct {
	Service service.GyroscopeServiceInterface
}

func NewGyroscopeController(s service.GyroscopeServiceInterface) *GyroscopeController {
	return &GyroscopeController{Service: s}
}

func (c *GyroscopeController) CreateGyroscope(w http.ResponseWriter, r *http.Request) {
	var data models.Gyroscope
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			http.Error(w, "Unexpected validation error", http.StatusBadRequest)
			return
		}

		messageMap := make(map[string]string)
		errorMap := make(map[string]string)

		for _, fieldErr := range validationErrors {
			field := fieldErr.Field()
			if field == "X" || field == "Y" || field == "Z" {
				messageMap[field] = fmt.Sprintf("Field axis %s is require", field)
			} else {
				messageMap[field] = fmt.Sprintf("Field %s is require", field)
			}
			errorMap[field] = fieldErr.Error()
		}

		response := map[string]interface{}{
			"message": messageMap,
			"error":   errorMap,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
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
