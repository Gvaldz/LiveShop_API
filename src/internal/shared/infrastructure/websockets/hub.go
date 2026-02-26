package websockets

import (
	"log"
	"sync"
	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[int32]*websocket.Conn
	mu      sync.RWMutex 
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[int32]*websocket.Conn),
	}
}

func (h *Hub) AddClient(userID int32, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[userID] = conn
	log.Printf("Usuario %d conectado al WebSocket", userID)
}

func (h *Hub) RemoveClient(userID int32) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, userID)
	log.Printf("Usuario %d desconectado del WebSocket", userID)
}

func (h *Hub) NotifyUser(userID int32, message interface{}) error {
	h.mu.RLock()
	conn, exists := h.clients[userID]
	h.mu.RUnlock()

	if !exists {
		return nil
	}

	return conn.WriteJSON(message)
}