package redis

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedisConnection() {
	port := os.Getenv("REDIS_PORT")
	addr := fmt.Sprintf("localhost:%s", port)
	password := os.Getenv("REDIS_PASSWORD")

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	if err := Redis.Ping(context.Background()).Err(); err != nil {
		log.Fatal("failed to connect to postgres instance")
	}
}
