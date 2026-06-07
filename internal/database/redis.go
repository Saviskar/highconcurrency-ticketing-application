package database

import (
	"context"
	"log"

	"ticketing-application/internal/config"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	addr := config.GetEnv("REDIS_ADDR", "localhost:6379")
	password := config.GetEnv("REDIS_PASSWORD", "")

	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis: ", err)
	}

	log.Println("Successfully connected to Redis!")
}
