package api

import (
	"net/http"

	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func ListLanguages(c *gin.Context) {
	var langs []models.Language
	postgres.Postgres.Preload(clause.Associations).Find(&langs)

	c.JSON(http.StatusOK, langs)
}
