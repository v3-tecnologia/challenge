package db

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDb() {
	var err error

	DB, err = sql.Open("postgres", config.DbUrl)
	if err != nil {
		log.Fatal(custom_errors.NewDBError(err, http.StatusInternalServerError))
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(custom_errors.NewDBError(err, http.StatusInternalServerError))
	}
	log.Println("Connected to database")
}
