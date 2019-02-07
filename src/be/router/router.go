package router

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	users := api.Group("/users")
	userRouter(users)

	return r
}
