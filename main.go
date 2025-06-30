package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/kaptinlin/jsonschema"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nats-io/nats.go"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	_ "v3challenge/docs"
)

type AppContext struct {
	DB          *sql.DB
	NC          *nats.Conn
	Compiler    *jsonschema.Compiler
	GpsSchema   *jsonschema.Schema
	GyroSchema  *jsonschema.Schema
	PhotoSchema *jsonschema.Schema
}

func NewAppContext() (*AppContext, error) {
	ctx := &AppContext{}

	err := godotenv.Load()
	if err == nil {
		log.Print("loaded .env file")
	}

	ctx.DB, err = openDatabase()
	if err != nil {
		log.Print("error opening database")
		log.Fatal(err)
	}

	ctx.Compiler = jsonschema.NewCompiler()

	buildSchemas(ctx)

	ctx.NC, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Print("failed to connect to nats")
		log.Fatal(err)
	}
	log.Print("connected to nats")

	return ctx, nil
}

func openDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return nil, err
	}

	log.Print("connected to database")

	var sqliteVersion string
	err = db.QueryRow("SELECT sqlite_version();").Scan(&sqliteVersion)
	if err != nil {
		log.Print("could not connect to database")
		return nil, err
	}

	log.Print("sqlite version: ", sqliteVersion)

	createTables(db)

	return db, nil
}

func buildSchemas(ctx *AppContext) {
	log.Print("building validation schemas")

	ctx.buildGyroSchema()
	ctx.buildGpsSchema()
	ctx.buildPhotoSchema()
}

func createTables(db *sql.DB) {
	log.Print("creating tables")

	createGyroTable(db)
	createGpsTable(db)
	createPhotoTable(db)
}

func startServer(ctx *AppContext) error {
	go ctx.gyroConsumer()
	go ctx.gpsConsumer()
	go ctx.photoConsumer()

	http.HandleFunc("/telemetry/gyroscope", ctx.gyroHandler)
	http.HandleFunc("/telemetry/gps", ctx.gpsHandler)
	http.HandleFunc("/telemetry/photo", ctx.photoHandler)

	http.Handle("/swagger/", httpSwagger.WrapHandler)
	log.Print("serving swagger")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("listening on port %s", port)
	return http.ListenAndServe(":"+port, nil)
}

// @title			api de telemetria
// @version		0.0.1
// @description	recebe dados de sensores e publica via nats
// @host			localhost:3000
// @BasePath		/
func main() {
	ctx, err := NewAppContext()
	if err != nil {
		log.Print("error creating app context")
		log.Fatal(err)
	}

	err = startServer(ctx)
	if err != nil {
		log.Print("failed to start server")
		log.Fatal(err)
	}
}
