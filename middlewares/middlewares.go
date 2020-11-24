package middlewares

import (
	"linqiurong2021/gin-book-frontend/cached"
	"linqiurong2021/gin-book-frontend/dao"
	"linqiurong2021/gin-book-frontend/myjwt"
	"linqiurong2021/gin-book-frontend/utils"
	"linqiurong2021/gin-book-frontend/validator"
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
		if token == "" {
			c.JSON(http.StatusBadRequest, utils.BadRequest("token must", ""))
			c.Abort()
			return
		}
		jwtToken, err := myjwt.Parse(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			c.Abort()
			return
		}
		myCliams, ok := myjwt.Check(jwtToken)
		if ok && jwtToken.Valid {
			// 存储当前用户信息
			cached.Save(myCliams)
			// 可以存储到 c.Set("user",user)
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, myCliams)
			c.Abort()
			return
		}
	}
}

// Page 分页
func Page() gin.HandlerFunc {
	// 判断是否存在用户
	return func(c *gin.Context) {
		var page dao.Page
		err := c.BindQuery(&page)
		// 参数校验判断
		ok := validator.Validate(c, err)
		if !ok {
			c.Abort()
			return
		}
		c.Next()
	}
}

// ID 删除
func ID() gin.HandlerFunc {
	// 判断是否存在用户
	return func(c *gin.Context) {
		var id dao.ID
		err := c.BindUri(&id)
		// 参数校验判断
		ok := validator.Validate(c, err)
		if !ok {
			c.Abort()
			return
		}
		c.Next()
	}
}
