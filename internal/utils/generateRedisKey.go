package utils

import "fmt"

func GenerateRedisKey(id string) string {
	return fmt.Sprintf("user:%s", id)
}
