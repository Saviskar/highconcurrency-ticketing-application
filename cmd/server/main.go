package main

import (
	"log"

	"ticketing-application/internal/config"
	"ticketing-application/internal/database"
	"ticketing-application/internal/handler"
	"ticketing-application/internal/ws"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using OS environment variables")
	} else {
		log.Println("Loaded .env file:")
		log.Printf("  DB:     %s:%s/%s (user=%s)", config.GetEnv("DB_HOST", ""), config.GetEnv("DB_PORT", ""), config.GetEnv("DB_NAME", ""), config.GetEnv("DB_USER", ""))
		log.Printf("  Redis:  %s", config.GetEnv("REDIS_ADDR", ""))
		log.Printf("  Server: :%s", config.GetEnv("SERVER_PORT", ""))
	}
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
