package middlewares

import (
	"linqiurong2021/gin-book-frontend/myjwt"
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthCheck 权限校验
func AuthCheck() gin.HandlerFunc {
	// 判断是否存在用户
	return func(c *gin.Context) {
		c.Next()
	}
}

// JWTTokenCheck JWTToken 校验
func JWTTokenCheck() gin.HandlerFunc {
	// 判断是否存在用户
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		jwtToken, err := myjwt.Parse(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			c.Abort()
			return
		}
		myCliams, ok := myjwt.Check(jwtToken)
		if ok && jwtToken.Valid {
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, myCliams)
			c.Abort()
		}
	}
}
