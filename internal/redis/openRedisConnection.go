package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/edelwei88/bytebuild-go/internal/config"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func OpenRedisConnection() {
	port := config.Config.Redis.Port
	password := config.Config.Redis.Password
	addr := fmt.Sprintf("localhost:%s", port)

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	if err := Redis.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
}
