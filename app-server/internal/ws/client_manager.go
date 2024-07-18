package ws

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

type ClientManager struct {
	sync.RWMutex

	clients  clientList
	handlers map[string]EventHandler
}

func NewClientManager() *ClientManager {
	m := &ClientManager{
		clients:  make(clientList),
		handlers: make(map[string]EventHandler),
	}

	m.setUpEventHandlers()

	return m
}

func (m *ClientManager) AddClient(c *Client) {
	m.Lock()
	defer m.Unlock()

	m.clients[c] = true
}

func (m *ClientManager) setUpEventHandlers() {
	m.handlers[EventAuthenticate] = func(e Event, c *Client) error {
		var payload EventPayloadAuthenticate

		if err := json.Unmarshal(e.Payload, &payload); err != nil {
			return errors.New("error marshaling payload of type '" + e.Type + "'")
		}

		log.Printf("connection authenticated as %s:%s", payload.Id, payload.Secret)
		return nil
	}

	m.handlers[EventTest] = func(e Event, c *Client) error {
		var payload EventPayloadText

		if err := json.Unmarshal(e.Payload, &payload); err != nil {
			return errors.New("error marshaling payload of type '" + e.Type + "'")
		}

		log.Printf("test message received: %s", payload)

		response := GenericEvent[string]{Type: "test_response", Payload: string(payload)}
		data, err := json.Marshal(response)

		if err != nil {
			log.Println(err)
			return errors.New("error marshaling response")
		}

		c.egress <- data

		return nil
	}
}

func (m *ClientManager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}

		return nil
	}

	return ErrEventNotSupported
}

func (m *ClientManager) removeClient(c *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[c]; ok {
		c.connection.Close()
		delete(m.clients, c)
	}
}
