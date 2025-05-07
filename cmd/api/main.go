package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	env "github.com/bielgennaro/v3-challenge-cloud/config"
	"github.com/bielgennaro/v3-challenge-cloud/config/middleware"
	"github.com/bielgennaro/v3-challenge-cloud/internal/db"
	"github.com/bielgennaro/v3-challenge-cloud/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Starting V3 server...")

	_, err := db.GetConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection successful")

	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)

	routes.SetupTelemetryRoutes(router)

	port := env.GetEnv("PORT", "8080")
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := db.RunMigrations(); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	go func() {
		log.Printf("Server listening on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
