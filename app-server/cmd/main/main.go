package main

import (
	"app-server/internal/routes"
)

func main() {
	r := routes.ServerRouter()
	r.Run()
}
