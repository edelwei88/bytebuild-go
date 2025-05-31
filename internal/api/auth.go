package api

import (
	"net/http"

	"github.com/edelwei88/bytebuild-go/internal/token"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}

	user, err := token.GetUserByJWT(tokenString)
	if err != nil {
		c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, user)
}
