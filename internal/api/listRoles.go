package api

import (
	"net/http"

	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func ListRoles(c *gin.Context) {
	var roles []models.Role
	postgres.Postgres.Preload(clause.Associations).Find(&roles)

	c.JSON(http.StatusOK, roles)
}
