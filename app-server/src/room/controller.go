package room

import (
	"example/qrka/src/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	service RoomService
}

func NewRoomController(service RoomService) *RoomController {
	return &RoomController{
		service: service,
	}
}

func (ctr RoomController) PostRooms(ctx *gin.Context) {
	var dto RoomCreateDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := ctr.service.CreateRoom(dto)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"content": room})
}

func (ctr RoomController) GetRooms(ctx *gin.Context) {
	data, err := ctr.service.ListRooms()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (ctr RoomController) JoinRoom(ctx *gin.Context) {
	ws.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := ws.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctr.service.HandleRoomSocket(conn)
}
