package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/edelwei88/bytebuild-go/internal/redis"
	"github.com/edelwei88/bytebuild-go/internal/token"
	"github.com/edelwei88/bytebuild-go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func Login(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
	var crds struct {
		Email    string `binding:"required" json:"email"`
		Password string `binding:"required" json:"password"`
	}

	err := c.ShouldBind(&crds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	var users []models.User
	result := postgres.Postgres.Where(&models.User{
		Email:    crds.Email,
		Password: utils.HashString(crds.Password),
	}).Preload(clause.Associations).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	if len(users) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong email or password",
		})
		return
	}

	jwt, err := token.GenerateJWT(users[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to generate JWT",
		})
		return
	}

	idStr := strconv.Itoa(int(users[0].ID))
	err = redis.SetCache(utils.GenerateRedisKey(idStr), jwt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to set cache to redis",
		})
		return
	}

	c.SetCookie("Authorization", jwt, int(time.Hour.Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"user": users[0],
	})
}
