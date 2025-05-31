package api

import (
	"net/http"

	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func ListUsers(c *gin.Context) {
	var users []models.User
	postgres.Postgres.Preload(clause.Associations).Find(&users)

	c.JSON(http.StatusOK, users)
}
