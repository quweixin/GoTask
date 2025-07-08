package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"success": true,
		"message": "注册成功",
	})
}

func Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"success": true,
		"message": "登录成功",
	})
}
