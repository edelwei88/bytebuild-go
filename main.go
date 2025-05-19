package main

import (
	"github.com/edelwei88/bytebuild-go/controllers"
	"github.com/edelwei88/bytebuild-go/initialize"
	"github.com/gin-gonic/gin"
)

func setup() {
	initialize.LoadEnv()
	initialize.ConnectToDB()
}

func main() {
	setup()
	router := gin.Default()
	router.GET("/languages", controllers.LanguagesGET)
	router.POST("/compile", controllers.CompilePOST)

	router.Run()
}
