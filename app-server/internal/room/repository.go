package room

type RoomRepository interface {
	Save(game Room) error
	FindAll() ([]Room, error)
	Find(id string) (Room, error)
	Delete(id string) error
}
