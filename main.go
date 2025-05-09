package main

import (
	"v3-backend-challenge/src/api"
	"v3-backend-challenge/src/db"
	"v3-backend-challenge/src/env"
)

func main() {
	env.Init()
	db.Init()
	api.Init()
}
