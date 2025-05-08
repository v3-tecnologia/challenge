package main

import (
	"v3-backend-challenge/api"
	"v3-backend-challenge/db"
	"v3-backend-challenge/env"
)

func main() {
	env.Init()
	db.Init()
	api.RegisterRoutes()
}
