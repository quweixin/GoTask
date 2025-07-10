package initialize

import (
	"GoTask/task04/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	ApiGroup := Router.Group("/api/v1")
	router.InitUserRouter(ApiGroup)
	router.InitPostRouter(ApiGroup)
	router.InitCommentRouter(ApiGroup)
	return Router
}
