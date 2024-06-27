package main

import (
	"app-server/internal/repository"
	"app-server/internal/utilities"
	"app-server/internal/ws"
)

func main() {
	database := repository.InitDB()
	// repository.PerformMigration(database)
	clientManager := ws.NewClientManager()

	r := utilities.ServerRouter(database.GetDB(), clientManager)
	r.Run()
}
