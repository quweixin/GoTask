package router

import (
	"GoTask/task04/api"
	"GoTask/task04/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCommentRouter(RouterGroup *gin.RouterGroup) {
	CommentRouter := RouterGroup.Group("comment").Use(middlewares.JwtAuth())
	{
		CommentRouter.POST("", api.CreateComment)
		CommentRouter.GET(":id", api.GetCommentDetail)
		CommentRouter.DELETE(":id", api.DeleteComment)
	}
}
