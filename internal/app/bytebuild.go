package app

import (
	"github.com/edelwei88/bytebuild-go/internal/api"
	"github.com/edelwei88/bytebuild-go/internal/middlewares"
	"github.com/edelwei88/bytebuild-go/internal/setup"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startup() {
	setup.LoadEnv()
	setup.ConnectDB()
	setup.DownloadImages()
}

func Run() {
	startup()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.POST("/login", api.Login)
	r.POST("/register", api.Register)
	r.POST("/compile", middlewares.ForAuthorized([]string{"user"}), api.Compile)
	r.GET("/langs", api.ListLanguages)

	r.Run()
}
