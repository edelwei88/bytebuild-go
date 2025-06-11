package api

import (
	"net/http"
	"strconv"

	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/edelwei88/bytebuild-go/internal/redis"
	"github.com/edelwei88/bytebuild-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to read JWT",
		})
		return
	}

	userID := strconv.Itoa(int(user.(models.User).ID))
	redis.DeleteCache(utils.GenerateRedisKey(userID))

	c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "session cleared",
	})
}
