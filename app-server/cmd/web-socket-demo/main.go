package main

import (
	"github.com/justyanna/toy-box/src/room"
	"github.com/justyanna/toy-box/src/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	clientManager := ws.NewClientManager()

	router.SetTrustedProxies([]string{})
	router.Use(cors.Default())

	room.RegisterRoutes(router, clientManager)

	router.Run("127.0.0.1:5000")
}
