package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ricardoraposo/challenge/internal/repository"
)

type Database struct {
	Pool  *pgxpool.Pool
	Query *repository.Queries
}

func NewDatabase() *Database {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	query := repository.New(pool)

	database := Database{
		Pool:  pool,
		Query: query,
	}

	return &database
}
