package db

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/KaiRibeiro/challenge/internal/logs"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDb() {
	logs.Logger.Info("starting db")
	var err error

	DB, err = sql.Open("postgres", config.DbUrl)
	if err != nil {
		wrappedErr := custom_errors.NewDBError(err, http.StatusInternalServerError)
		logs.Logger.Error("failed to open DB connection",
			"error", wrappedErr,
		)
		os.Exit(1)
	}

	err = DB.Ping()
	if err != nil {
		wrappedErr := custom_errors.NewDBError(err, http.StatusInternalServerError)
		logs.Logger.Error("failed to ping DB",
			"error", wrappedErr,
		)
		os.Exit(1)
	}
	logs.Logger.Info("db started")
}
