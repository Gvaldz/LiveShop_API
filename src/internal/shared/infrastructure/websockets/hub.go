package websockets

import (
	"sync"
	"github.com/gorilla/websocket"
)

type client struct {
	conn *websocket.Conn
	mu   sync.Mutex 
}

type Hub struct {
	clients map[int32]*client 
	mu      sync.RWMutex      
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[int32]*client),
	}
}

func (h *Hub) AddClient(userID int32, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[userID] = &client{conn: conn}
}

func (h *Hub) RemoveClient(userID int32) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, userID)
}

func (h *Hub) NotifyUser(userID int32, message interface{}) error {
	h.mu.RLock()
	c, exists := h.clients[userID]
	h.mu.RUnlock()

	if !exists {
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.WriteJSON(message)
}

func (h *Hub) Broadcast(message interface{}) error {
	h.mu.RLock()
	clientsToNotify := make([]*client, 0, len(h.clients))
	for _, c := range h.clients {
		clientsToNotify = append(clientsToNotify, c)
	}
	h.mu.RUnlock()

	for _, c := range clientsToNotify {
		go func(cl *client) {
			cl.mu.Lock()
			defer cl.mu.Unlock()
			_ = cl.conn.WriteJSON(message)
		}(c)
	}
	
	return nil
}