package router

import (
	"GoTask/task04/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(RouterGroup *gin.RouterGroup) {
	UserRouter := RouterGroup.Group("user")
	{
		UserRouter.POST("register", api.Register)
		UserRouter.POST("login", api.Login)
	}
}
