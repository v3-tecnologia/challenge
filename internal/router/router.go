package router

import (
	controller "challenge-cloud/internal/controllers"

	"github.com/gorilla/mux"
)

type Controllers struct {
	Gyro  *controller.GyroscopeController
	GPS   *controller.GPSController
	Photo *controller.PhotoController
	Auth  *controller.AuthController
}

func LoadRouter(c Controllers) *mux.Router {
	api := mux.NewRouter()

	api.HandleFunc("/auth/login", c.Auth.LoginHandler).Methods("POST")

	api.HandleFunc("/telemetry/gyroscope", c.Gyro.CreateGyroscope).Methods("POST")
	api.HandleFunc("/telemetry/gyroscope", c.Gyro.GetGyroscope).Methods("GET")

	api.HandleFunc("/telemetry/gps", c.GPS.CreateGPS).Methods("POST")
	api.HandleFunc("/telemetry/gps", c.GPS.GetGPS).Methods("GET")

	api.HandleFunc("/telemetry/photo", c.Photo.CreatePhoto).Methods("POST")
	api.HandleFunc("/telemetry/photo", c.Photo.GetPhoto).Methods("GET")
	return api
}
