package api

import (
	"net/http"
	"strconv"

	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func ListUserCompiles(c *gin.Context) {
	var req struct {
		ID string `json:"id" binding:"required"`
	}

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	parsed, _ := strconv.Atoi(req.ID)
	user := models.User{
		ID: uint(parsed),
	}

	postgres.Postgres.Preload(clause.Associations).Preload("Compiles.Compiler").Find(&user)

	c.JSON(http.StatusOK, user.Compiles)
}
