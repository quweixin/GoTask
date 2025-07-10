package router

import (
	"GoTask/task04/api"
	"GoTask/task04/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(RouterGroup *gin.RouterGroup) {
	UserRouter := RouterGroup.Group("user")
	{
		UserRouter.POST("register", api.Register)
		UserRouter.POST("login", api.Login)
		UserRouter.Use(middlewares.JwtAuth()).GET("detail", api.UserDetail)
	}
}
