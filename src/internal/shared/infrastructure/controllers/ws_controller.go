package controllers

import (
	"net/http"
	"liveshop_api/src/internal/shared/infrastructure/websockets"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}

type WSController struct {
	hub *websockets.Hub
}

func NewWSController(hub *websockets.Hub) *WSController {
	return &WSController{hub: hub}
}

func (ws *WSController) Connect(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
		return
	}
	userID := userIDInterface.(int32)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	ws.hub.AddClient(userID, conn)

	defer func() {
		ws.hub.RemoveClient(userID)
		conn.Close()
	}()

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}