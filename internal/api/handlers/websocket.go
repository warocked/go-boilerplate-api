package handlers

import (
	"log"
	"sync"

	"github.com/gofiber/websocket/v2"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.RWMutex
)

func HandleWebSocket(c *websocket.Conn) {
	defer func() {
		clientsMu.Lock()
		delete(clients, c)
		clientsMu.Unlock()
		c.Close()
	}()

	clientsMu.Lock()
	clients[c] = true
	clientsMu.Unlock()

	var (
		mt  int
		msg []byte
		err error
	)

	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("websocket read error:", err)
			break
		}

		clientsMu.RLock()
		for client := range clients {
			if err = client.WriteMessage(mt, msg); err != nil {
				log.Println("websocket write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
		clientsMu.RUnlock()
	}
}

func BroadcastMessage(message []byte) {
	clientsMu.RLock()
	defer clientsMu.RUnlock()
	
	for client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("broadcast error:", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func GetConnectedClientsCount() int {
	clientsMu.RLock()
	defer clientsMu.RUnlock()
	return len(clients)
}
