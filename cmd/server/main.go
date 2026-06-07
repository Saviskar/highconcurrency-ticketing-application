package main

import (
	"ticketing-application/internal/config"
	"ticketing-application/internal/database"
	"ticketing-application/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	database.InitRedis()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/slots", handler.GetSlots)
		v1.POST("/slots/book", handler.BookSlot)
	}

	port := config.GetEnv("SERVER_PORT", "8080")
	r.Run(":" + port)
}
