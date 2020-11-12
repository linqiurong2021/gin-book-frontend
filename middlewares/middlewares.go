package middlewares

import "github.com/gin-gonic/gin"

// AuthCheck 权限校验
func AuthCheck() gin.HandlerFunc {
	// 判断是否存在用户
	return func(c *gin.Context) {
		c.Next()
	}
}
