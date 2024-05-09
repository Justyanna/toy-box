package room

import (
	"github.com/gin-gonic/gin"
	"github.com/justyanna/toy-box/src/ws"
)

func RegisterRoutes(router *gin.Engine, clientManager *ws.ClientManager) {
	repository := NewMemRepository()
	service := NewDefaultService(repository, clientManager)
	controller := NewRoomController(service)

	router.POST("/rooms", controller.PostRooms)
	router.GET("/rooms", controller.GetRooms)
	router.GET("/rooms/ws", controller.JoinRoom)
}
