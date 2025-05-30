package main

import (
	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/edelwei88/bytebuild-go/internal/setup"
)

func startup() {
	setup.LoadEnv()
	setup.ConnectDB()
}

func main() {
	startup()
	postgres.Postgres.AutoMigrate(&models.Language{}, &models.Compiler{}, &models.Compile{}, &models.Role{}, &models.User{})
}
