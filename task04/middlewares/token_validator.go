package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authorization := context.GetHeader("Authorization")
		if authorization == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "未提供认证令牌",
			})
			context.Abort()
			return
		}
		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "认证令牌格式错误",
			})
			context.Abort()
			return
		}
		claims, err := PareToken(parts[1])
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "无效的认证令牌",
				"error":   err.Error(),
			})
			context.Abort()
			return
		}
		context.Set("userId", claims.UserID)
		context.Set("username", claims.Username)
		context.Next()
	}
}
