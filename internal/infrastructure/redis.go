package infrastructure

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatal("Failed to connect to Redis: ", err)
	}

	log.Println("Successfully connected to Redis!")
	return rdb
}
