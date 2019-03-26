package router

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	users := api.Group("/users")
	tweets := api.Group("/tweets")

	userRouter(users)
	tweetRouter(tweets)

	return r
}
