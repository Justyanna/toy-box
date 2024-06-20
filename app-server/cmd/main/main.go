package main

import (
	"app-server/internal/repository"
	"app-server/internal/utilis"
)

func main() {
	database := repository.InitDB()
	// repository.PerformMigration(database)

	r := utilis.ServerRouter(database.GetDB())
	r.Run()
}
