package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type dockerConfig struct {
	APIVersion string
}

type postgresConfig struct {
	Port     string
	User     string
	Password string
	DB       string
}

type redisConfig struct {
	Port     string
	Password string
}

type jwtConfig struct {
	Secret string
}

type appConfig struct {
	Port           string
	WebsiteAddress string
	Docker         dockerConfig
	Postgres       postgresConfig
	Redis          redisConfig
	JWT            jwtConfig
}

var Config appConfig

func Init() {
	readFromEnv()

	Config.Port = envOrDefault("PORT", "3001")
	Config.WebsiteAddress = envOrDefault("WEBSITE_ADDRESS", "http://localhost:3000")

	Config.Docker.APIVersion = envOrDefault("DOCKER_API_VERSION", "1.47")

	Config.Postgres.Port = envOrDefault("POSTGRES_PORT", "8080")
	Config.Postgres.User = envOrDefault("POSTGRES_USER", "bytebuild_user")
	Config.Postgres.Password = envOrDefault("POSTGRES_PASSWORD", "bytebuild_password")
	Config.Postgres.DB = envOrDefault("POSTGRES_DB", "bytebuild")

	Config.Redis.Port = envOrDefault("REDIS_PORT", "8081")
	Config.Redis.Password = envOrDefault("REDIS_PASSWORD", "bytebuild_password")

	Config.JWT.Secret = envOrDefault("JWT_SECRET_KEY", "jwt_secret")
}

func readFromEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found, loading default config")
	}
}

func envOrDefault(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
