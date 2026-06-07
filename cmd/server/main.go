package main

import (
	"ticketing-application/internal/config"
	"ticketing-application/internal/database"
	"ticketing-application/internal/handler"
	"ticketing-application/internal/ws"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	database.InitRedis()

	hub := ws.NewHub()
	go hub.Run()
	handler.SlotHub = hub

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/slots", handler.GetSlots)
		v1.POST("/slots/book", handler.BookSlot)
		v1.GET("/ws", func(c *gin.Context) {
			hub.ServeWS(c.Writer, c.Request)
		})
	}

	port := config.GetEnv("SERVER_PORT", "8080")
	r.Run(":" + port)
}
