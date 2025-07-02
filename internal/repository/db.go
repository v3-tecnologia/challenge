package repository

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

var db *sql.DB

func Connect(databaseURL string) {
    var err error
    db, err = sql.Open("postgres", databaseURL)
    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }
    log.Println("Conectado ao banco de dados")
}

func Close() {
    if db != nil {
        db.Close()
    }
}