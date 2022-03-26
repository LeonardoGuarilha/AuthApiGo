package main

import (
	"auth-api/infra/database"
	"auth-api/server"
)

func main() {
	database.StartDB()

	runServer := server.NewServer()

	runServer.Run()
}
