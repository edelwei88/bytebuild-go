package middlewares

import (
	"net/http"
	"slices"

	"github.com/edelwei88/bytebuild-go/internal/token"
	"github.com/gin-gonic/gin"
)

func ForAuthorized(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
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
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		if !slices.Contains(roles, user.Role.Name) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "your role has insufficent rights",
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("token", tokenString)
		c.Next()
	}
}
