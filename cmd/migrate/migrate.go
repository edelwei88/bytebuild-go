package main

import (
	"github.com/edelwei88/bytebuild-go/internal/config"
	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
)

func startup() {
	config.Init()
	postgres.OpenPostgresConnection()
}

func main() {
	startup()
	postgres.Postgres.AutoMigrate(&models.Language{},
		&models.Compiler{},
		&models.Compile{},
		&models.Role{},
		&models.User{})
}
