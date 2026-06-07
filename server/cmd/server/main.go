package main

import (
	deliveryhttp "ticketing-application/internal/delivery/http"
	"ticketing-application/internal/infrastructure"
	"ticketing-application/internal/infrastructure/ws"
	"ticketing-application/internal/repository/postgres"
	"ticketing-application/internal/repository/redis"
	"ticketing-application/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := infrastructure.LoadConfig()

	db := infrastructure.NewPostgresDB(cfg)
	postgres.AutoMigrate(db)
	postgres.SeedSlots(db)

	rdb := infrastructure.NewRedisClient(cfg)

	slotRepo := postgres.NewSlotRepository(db)
	lockRepo := redis.NewDistributedLock(rdb)
	hub := ws.NewHub()
	go hub.Run()

	slotUC := usecase.NewSlotUseCase(slotRepo, lockRepo, hub)
	slotHandler := deliveryhttp.NewSlotHandler(slotUC)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/slots", slotHandler.GetSlots)
		v1.POST("/slots/book", slotHandler.BookSlot)
		v1.GET("/ws", func(c *gin.Context) {
			hub.ServeWS(c.Writer, c.Request)
		})
	}

	r.Run(":" + cfg.ServerPort)
}
