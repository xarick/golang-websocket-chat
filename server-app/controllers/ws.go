package controllers

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket sozlash
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Ulanishlar ro‘yxati
var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex

// Xabar yuborish
func broadcastMessage(message []byte, sender *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	for client := range clients {
		if client != sender {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Xabar yuborishda xatolik:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// WebSocket ulanishini boshqarish
func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Ulanishda xatolik:", err)
		return
	}
	defer conn.Close()

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	fmt.Println("Yangi foydalanuvchi ulandi")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Klientdan uzilish:", err)
			break
		}
		broadcastMessage(msg, conn)
	}
}
