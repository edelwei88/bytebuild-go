package app

import (
	"net/http"

	"github.com/edelwei88/bytebuild-go/internal/api"
	"github.com/edelwei88/bytebuild-go/internal/config"
	"github.com/edelwei88/bytebuild-go/internal/middlewares"
	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/redis"
	"github.com/edelwei88/bytebuild-go/internal/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Init()
	token.SetupJWT()
	postgres.OpenPostgresConnection()
	redis.OpenRedisConnection()
}

func Run() {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{config.Config.WebsiteAddress},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsConfig))

	auth := r.Group("/auth")
	auth.POST("/login", api.Login)
	auth.POST("/register", api.Register)

	user := r.Group("/user", middlewares.ForAuthorized([]string{"user", "manager", "admin"}))
	user.GET("/me", api.Me)
	user.GET("/logout", api.Logout)
	user.POST("/compile", api.Compile)
	user.GET("/langs", api.ListLanguages)

	manager := r.Group("/manager", middlewares.ForAuthorized([]string{"manager", "admin"}))
	manager.GET("/compiles", api.ListCompiles)
	manager.GET("/users", api.ListUsers)
	manager.POST("/users/compiles", api.ListUserCompiles)

	admin := r.Group("/admin", middlewares.ForAuthorized([]string{"admin"}))
	admin.GET("/roles", api.ListRoles)
	admin.PATCH("/users/patch", api.PatchUser)
	admin.DELETE("/users/delete/:id", api.DeleteUser)

	r.Run()
}
