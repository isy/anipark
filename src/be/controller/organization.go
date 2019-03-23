package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllOrganizations ユーザ一覧
func GetAllOrganizations(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping",
	})
}
