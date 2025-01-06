package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-websocket-chat/server/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// WebSocket marshruti
	r.GET("/ws", controllers.HandleWebSocket)

	r.GET("/user", controllers.GetUser)

	return r
}
