package config

import (
	"os"
)

var (
	Port  = os.Getenv("PORT")
	DbUrl = os.Getenv("DB_URL")
)
