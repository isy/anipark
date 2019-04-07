package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers ユーザ一覧
func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping",
	})
}
