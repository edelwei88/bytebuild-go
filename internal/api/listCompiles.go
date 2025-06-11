package api

import (
	"net/http"

	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func ListCompiles(c *gin.Context) {
	var cmp []models.Compile
	postgres.Postgres.Preload(clause.Associations).Find(&cmp)

	c.JSON(http.StatusOK, cmp)
}
