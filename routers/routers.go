package routers

import (
	"linqiurong2021/gin-book-frontend/middlewares"
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		userGroup(v1)
		bookGroup(v1)
		cartGroup(v1)
		// 默认路由
		defaultRouter(r)
	}
}

func notMethod(c *gin.Context) {
	c.JSON(http.StatusBadRequest, utils.BadRequest("bad request", ""))
	c.Abort()
}

//
func defaultRouter(r *gin.Engine) {

	r.NoMethod(notMethod)
}

// UserGroup User路由
func userGroup(g *gin.RouterGroup) {
	g.Group("/user").Use(middlewares.AuthCheck())
	{

	}
}

// BookGroup Book路由
func bookGroup(g *gin.RouterGroup) {
	g.Group("/book")
	{

	}
}

// CartGroup 用户路由
func cartGroup(g *gin.RouterGroup) {
	g.Group("/cart")
	{

	}
}
