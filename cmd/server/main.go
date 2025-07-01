package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/igorlopes88/desafio-v3/configs"
	_ "github.com/igorlopes88/desafio-v3/docs"
	"github.com/igorlopes88/desafio-v3/internal/entity"
	"github.com/igorlopes88/desafio-v3/internal/infra/database"
	"github.com/igorlopes88/desafio-v3/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title API Desafio V3 (Cloud)
// @version 1.0
// @description API + Authentication + Documentation
// @termsOfService http://swagger.io/terms/

// @contact.name Igor Lopes
// @contact.email igorlopes.ilhabela@gmail.com

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	conf, err := configs.Load(".")
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBUser,
		conf.DBPassword,
		conf.DBAddress,
		conf.DBPort,
		conf.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Gyroscope{}, &entity.Gps{}, &entity.Photo{})

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	gyroDB := database.NewGyroscope(db)
	gyroscopeHandler := handlers.NewGyroscopeHandler(gyroDB)

	gpsDB := database.NewGps(db)
	gpsHandler := handlers.NewGpsHandler(gpsDB)

	photoDB := database.NewPhoto(db)
	photoHandler := handlers.NewPhotoHandler(photoDB)

	r := chi.NewRouter()

	// MIDDLEWARE
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", conf.TokenAuth))
	r.Use(middleware.WithValue("jwtExperesIn", conf.JWTExperesIn))

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	r.Route("/users", func(r chi.Router) {
		r.Post("/create", userHandler.Create)
		r.Post("/generate_token", userHandler.GetJWT)
	})

	r.Route("/telemetry", func(r chi.Router) {
		r.Use(jwtauth.Verifier(conf.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/gyroscope", gyroscopeHandler.Register)
		r.Post("/gps", gpsHandler.Register)
		r.Post("/photo", photoHandler.Register)
	})

	http.ListenAndServe(":8000", r)
}
