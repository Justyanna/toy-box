package ws

import "encoding/json"

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

const (
	EventAuthenticate = "authenticate"
	EventTest         = "test"
)

type EventPayloadAuthenticate struct {
	Id     string `json:"id"`
	Secret string `json:"secret"`
}

type EventPayloadText string

type GenericEvent[P any] struct {
	Type    string `json:"type"`
	Payload P      `json:"payload"`
}
