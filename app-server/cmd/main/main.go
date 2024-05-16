package main

import (
	"app-server/internal/repository"
	"app-server/internal/routes"
)

func main() {
	database := repository.InitDB()
	// repository.PerformMigration(database)

	r := routes.ServerRouter(database)
	r.Run()
}
