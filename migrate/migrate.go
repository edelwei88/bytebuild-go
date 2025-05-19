package main

import (
	"github.com/edelwei88/bytebuild-go/initialize"
	"github.com/edelwei88/bytebuild-go/models"
)

func init() {
	initialize.LoadEnv()
	initialize.ConnectToDB()
}

func main() {
	initialize.DB.AutoMigrate(&models.Language{}, &models.Compiler{}, &models.Compile{})
}
