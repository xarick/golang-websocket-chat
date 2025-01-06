package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8060/ws", nil)
	if err != nil {
		log.Fatal("Ulanishda xatolik:", err)
	}
	defer conn.Close()

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Xabarni o‘qishda xatolik:", err)
				return
			}
			fmt.Printf("Boshqa foydalanuvchi: %s\n", msg)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Klient 2 (Siz): ")
		scanner.Scan()
		message := scanner.Text()

		err := conn.WriteMessage(websocket.TextMessage, []byte("Klient 2: "+message))
		if err != nil {
			log.Println("Xabar yuborishda xatolik:", err)
			return
		}
	}
}
