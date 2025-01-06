package main

import (
	"fmt"
	"log"

	"github.com/xarick/golang-websocket-chat/server/routes"
)

func main() {
	r := routes.SetupRouter()

	fmt.Println("WebSocket serveri 8060 portda ishga tushdi...")

	if err := r.Run(":8060"); err != nil {
		log.Fatal("Serverda xatolik:", err)
	}
}
