package controllers

import (
	"net/http"

	"github.com/edelwei88/bytebuild-go/initialize"
	"github.com/edelwei88/bytebuild-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func LanguagesGET(c *gin.Context) {
	var languages []models.Language
	result := initialize.DB.Preload(clause.Associations).Find(&languages)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(200, languages)
}
