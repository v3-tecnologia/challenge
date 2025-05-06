package config

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func SetupPostgres() *pgxpool.Pool {
	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL not set")
	}

	connectionPool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	return connectionPool
}
