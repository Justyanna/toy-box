package room

import (
	"example/qrka/src/ws"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, clientManager *ws.ClientManager) {
	repository := NewMemRepository()
	service := NewDefaultService(repository, clientManager)
	controller := NewRoomController(service)

	router.POST("/rooms", controller.PostRooms)
	router.GET("/rooms", controller.GetRooms)
	router.GET("/rooms/ws", controller.JoinRoom)
}
