package api

import (
	"net/http"

	"github.com/edelwei88/bytebuild-go/internal/redis"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}

	redis.DeleteCache(tokenString)

	c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "session cleared",
	})
}
