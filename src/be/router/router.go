package router

import(
	"github.com/gin-gonic/gin"
	"github.com/isy/ownedia/controller"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api") {
		users := api.Group("/users")
		userRouter(users)
	}
}