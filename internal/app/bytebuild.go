package app

import (
	"net/http"

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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/login", api.Login)
	r.POST("/register", api.Register)
	r.GET("/langs", api.ListLanguages)
	r.POST("/compile", middlewares.ForAuthorized([]string{"user"}), api.Compile)
	r.GET("/auth", api.Auth)
	r.GET("/logout", api.Logout)
	r.GET("/users", middlewares.ForAuthorized([]string{"manager", "admin"}), api.ListUsers)
	r.POST("/usercompiles", middlewares.ForAuthorized([]string{"manager", "admin"}), api.ListUserCompiles)
	r.GET("/compiles", middlewares.ForAuthorized([]string{"manager", "admin"}), api.ListCompiles)

	r.Run()
}
