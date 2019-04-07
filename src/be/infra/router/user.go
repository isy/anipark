package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isy/anipark/src/be/interface/controller"
)

func userRouter(r *gin.RouterGroup) {
	r.GET("/", controller.GetAllUsers)
}
