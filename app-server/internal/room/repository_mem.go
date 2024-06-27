package room

import (
	"errors"
)

var (
	ErrRoomNotFound = errors.New("room not found")
	ErrRoomExists   = errors.New("room exists")
)

type GameMemoryRepository struct {
	data  map[string]Room
	index []string
}

func NewMemRepository() *GameMemoryRepository {
	list := make(map[string]Room)
	var index []string

	return &GameMemoryRepository{
		list,
		index,
	}
}

func (repository *GameMemoryRepository) Save(game Room) error {
	// -- Update index when adding new item
	if _, ok := repository.data[game.ID]; !ok {
		repository.index = append(repository.index, game.ID)
	}

	repository.data[game.ID] = game

	return nil
}

func (repository *GameMemoryRepository) FindAll() ([]Room, error) {
	data := make([]Room, len(repository.index))

	for index, id := range repository.index {
		data[index] = repository.data[id]
	}

	return data, nil
}

func (repository *GameMemoryRepository) Find(id string) (Room, error) {
	if val, ok := repository.data[id]; ok {
		return val, nil
	}

	return Room{}, ErrRoomNotFound
}

func (repository *GameMemoryRepository) Delete(id string) error {
	// -- Update index when removing item
	for index, value := range repository.index {
		if value == id {
			repository.index = append(repository.index[:index], repository.index[index+1:]...)
		}
	}

	delete(repository.data, id)

	return nil
}
