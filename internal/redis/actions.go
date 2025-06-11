package redis

import (
	"context"
	"time"
)

func SetCache(key string, value string) error {
	ctx := context.Background()
	status := Redis.Set(ctx, key, value, time.Hour)
	if status.Err() != nil {
		return status.Err()
	}

	return nil
}

func GetCache(key string) (string, error) {
	ctx := context.Background()
	value, err := Redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func DeleteCache(key string) {
	ctx := context.Background()
	Redis.Del(ctx, key)
}
