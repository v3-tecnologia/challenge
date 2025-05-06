package database_test

import (
	"path"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/infra/database"
)

func init() {
	godotenv.Load(path.Join("..", "..", "..", ".env"))
}

func TestDatabaseConnection(t *testing.T) {
	db := database.NewDatabase(
		"test",
		"test",
		"test",
		"test",
		true,
	)

	err := db.Connect()
	require.NoError(t, err, "Failed to connect to the database")
	defer db.Close()

	err = db.Ping()
	require.NoError(t, err, "Failed to ping the database")
}

func TestDatabaseCloseConnection(t *testing.T) {
	db := database.NewDatabase(
		"test",
		"test",
		"test",
		"test",
		true,
	)

	err := db.Connect()
	require.NoError(t, err, "Failed to connect to the database")

	err = db.Ping()
	require.NoError(t, err, "Failed to ping the database")

	err = db.Close()
	require.NoError(t, err, "Failed to close the database connection")

	err = db.Ping()
	require.Error(t, err, "Expected error when pinging closed database connection")
}
