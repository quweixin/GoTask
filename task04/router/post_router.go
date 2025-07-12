package router

import (
	"GoTask/task04/api"
	"GoTask/task04/middlewares"
	"github.com/gin-gonic/gin"
)

func InitPostRouter(RouterGroup *gin.RouterGroup) {
	PostRouter := RouterGroup.Group("post").Use(middlewares.JwtAuth())
	{
		PostRouter.POST("", api.CreatePost)
		PostRouter.PUT(":id", api.UpdatePost)
		PostRouter.DELETE(":id", api.DeletePost)
		PostRouter.GET(":id", api.GetPostDetail)
		PostRouter.GET(":id/comments", api.GetCommentList)
		PostRouter.GET("/list", api.GetPostList)
	}

}
