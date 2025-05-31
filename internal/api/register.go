package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/edelwei88/bytebuild-go/internal/redis"
	"github.com/edelwei88/bytebuild-go/internal/token"
	"github.com/edelwei88/bytebuild-go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var crds struct {
		Username string `binding:"required" json:"username"`
		Email    string `binding:"required" json:"email"`
		Password string `binding:"required" json:"password"`
	}

	err := c.ShouldBind(&crds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
		return
	}

	var users []models.User
	result := postgres.Postgres.Where(&models.User{
		Email: crds.Email,
	}).Find(&users)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
		return
	}
	if len(users) != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "this email is already used",
		})
		c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
		return
	}

	var role models.Role
	result = postgres.Postgres.Where(&models.Role{
		Name: "user",
	}).First(&role)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get role user",
		})
		c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
		return
	}

	user := models.User{
		Username: crds.Username,
		Email:    crds.Email,
		Password: utils.HashString(crds.Password),
		Role:     role,
		Compiles: []models.Compile{},
	}
	result = postgres.Postgres.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create DB row",
		})
		c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
		return
	}

	jwt, err := token.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to generate JWT",
		})
		c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
		return
	}

	idStr := strconv.Itoa(int(user.ID))
	err = redis.SetCache(utils.GenerateRedisKey(idStr), jwt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to set cache to redis",
		})
		c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
		return
	}

	c.SetCookie("Authorization", jwt, 0, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
