package api

import (
	"net/http"

	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/edelwei88/bytebuild-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func PatchUser(c *gin.Context) {
	var req struct {
		ID       uint   `json:"id" binding:"required"`
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password"`
		RoleID   uint   `json:"role_id" binding:"required"`
	}

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	var user models.User
	result := postgres.Postgres.First(&user, req.ID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user",
		})
		return
	}
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no user with this id",
		})
		return
	}

	user.Username = req.Username
	user.Email = req.Email
	if len(req.Password) > 0 {
		user.Password = utils.HashString(req.Password)
	}
	user.RoleID = req.RoleID

	result = postgres.Postgres.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save user to db",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
