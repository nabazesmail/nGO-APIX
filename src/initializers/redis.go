package initializers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"), // Redis server address
		Password: "",                         // No password by default
		DB:       0,                          // Default database
	})

	// Ping the Redis server to check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}

func ResetCache() {
	ctx := context.Background()

	// Clear all cache (flush all databases)
	result, err := RedisClient.FlushAll(ctx).Result()
	if err != nil {
		log.Printf("Error resetting cache: %s", err)
		return
	}

	log.Printf("Cache reset: %s", result)
}
