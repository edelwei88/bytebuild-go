package api

import (
	"net/http"
	"strconv"

	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	result := postgres.Postgres.Delete(models.User{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to delete user in db",
		})
		return
	}

	c.Status(http.StatusOK)
}
