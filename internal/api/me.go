package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to read JWT",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
