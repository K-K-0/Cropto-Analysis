package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var mutex sync.Mutex

func Handler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("❌ Error upgrading:", err)
		return
	}
	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	log.Println("✅ Client connected")

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			log.Println("❌ Client disconnected:", err)
			break
		}
	}
}

func Broadcast(msg []byte) {
	mutex.Lock()
	defer mutex.Unlock()

	for conn := range clients {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("❌ Error broadcasting:", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}
