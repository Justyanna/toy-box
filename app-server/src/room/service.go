package room

import "github.com/gorilla/websocket"

type RoomService interface {
	CreateRoom(data RoomCreateDto) (Room, error)
	ListRooms() ([]Room, error)
	HandleRoomSocket(conn *websocket.Conn)
}
