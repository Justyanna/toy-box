package room

import (
	"time"

	"github.com/gosimple/slug"
)

type RoomCreateDto struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type Room struct {
	ID        string `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Type      string `json:"type" binding:"required"`
	Players   uint8  `json:"players" binding:"required"`
	Capacity  uint8  `json:"capacity" binding:"required"`
	CreatedAt string `json:"createdAt" binding:"required"`
}

func New(data RoomCreateDto) Room {
	id := slug.Make(data.Name)
	now := time.Now().Format(time.RFC3339)

	return Room{
		ID:        id,
		Name:      data.Name,
		Type:      data.Type,
		Players:   0,
		Capacity:  4,
		CreatedAt: now,
	}
}
