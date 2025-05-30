package setup

import (
	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/redis"
	"github.com/edelwei88/bytebuild-go/internal/token"
)

func ConnectDB() {
	postgres.InitPostgresConnection()
	redis.InitRedisConnection()
	token.InitJWT()
}
