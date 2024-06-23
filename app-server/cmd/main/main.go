package main

import (
	"app-server/internal/repository"
	"app-server/internal/utilities"
)

func main() {
	database := repository.InitDB()
	// repository.PerformMigration(database)

	r := utilities.ServerRouter(database.GetDB())
	r.Run()
}
