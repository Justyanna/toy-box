package room

import (
	"app-server/internal/ws"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, clientManager *ws.ClientManager) {
	repository := NewMemRepository()
	service := NewDefaultService(repository, clientManager)
	controller := NewRoomController(service)

	// Websocket related endpoints
	router.GET("/rooms", controller.GetRooms)
	router.POST("/rooms/", controller.CreateRooms)
	router.POST("/rooms/:room_id/join", controller.JoinRoom)
	router.DELETE("/rooms/:room_id", controller.RemoveRoom)
	router.POST("/rooms/:room_id/setup", controller.ChooseGame)
	router.POST("/rooms/:room_id/start", controller.StartGame)
	router.POST("/rooms/:room_id/connect", controller.ConnectWebsocket)
	router.POST("/rooms/:room_id/finish", controller.EndGame)
}
