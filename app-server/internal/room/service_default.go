package room

import (
	"app-server/internal/ws"

	"github.com/gorilla/websocket"
)

type DefaultRoomService struct {
	repository    RoomRepository
	clientManager *ws.ClientManager
}

func NewDefaultService(repository RoomRepository, clientManager *ws.ClientManager) *DefaultRoomService {
	return &DefaultRoomService{
		repository:    repository,
		clientManager: clientManager,
	}
}

func (srv *DefaultRoomService) CreateRoom(data RoomCreateDto) (Room, error) {
	room := New(data)

	if _, err := srv.repository.Find(room.ID); err == nil {
		return Room{}, ErrRoomExists
	}

	srv.repository.Save(room)

	return room, nil
}

func (srv *DefaultRoomService) ListRooms() ([]Room, error) {
	return srv.repository.FindAll()
}

// type socketData struct {
// 	Data string `json:"data"`
// }

func (srv *DefaultRoomService) HandleRoomSocket(conn *websocket.Conn) {
	client := ws.NewClient(conn, srv.clientManager)
	srv.clientManager.AddClient(client)
	client.HandleConnection()
}
