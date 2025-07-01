package router

import (
	controller "challenge-cloud/internal/controllers"
	middleware "challenge-cloud/internal/middlewares"
	"net/http"

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

	api.Handle("/telemetry/gyroscope", middleware.JWTAuth(http.HandlerFunc(c.Gyro.CreateGyroscope))).Methods("POST")
	api.Handle("/telemetry/gyroscope", middleware.JWTAuth(http.HandlerFunc(c.Gyro.GetGyroscope))).Methods("GET")

	api.Handle("/telemetry/gps", middleware.JWTAuth(http.HandlerFunc(c.GPS.CreateGPS))).Methods("POST")
	api.Handle("/telemetry/gps", middleware.JWTAuth(http.HandlerFunc(c.GPS.GetGPS))).Methods("GET")

	api.Handle("/telemetry/photo", middleware.JWTAuth(http.HandlerFunc(c.Photo.CreatePhoto))).Methods("POST")
	api.Handle("/telemetry/photo", middleware.JWTAuth(http.HandlerFunc(c.Photo.GetPhoto))).Methods("GET")

	return api
}
