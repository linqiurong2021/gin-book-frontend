package routers

import (
	"linqiurong2021/gin-book-frontend/controller"
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

// 未匹配到路由时
func noRoute(c *gin.Context) {
	c.JSON(http.StatusBadRequest, utils.BadRequest("not found route", ""))
	c.Abort()
}

// 系统默认的路由
func defaultRouter(r *gin.Engine) {

	r.NoMethod(notMethod)
	// 未匹配到路由时
	r.NoRoute(noRoute)
	// 心跳检测
	r.GET("/ping", controller.Ping)
}

// UserGroup User路由
func userGroup(g *gin.RouterGroup) {
	user := g.Group("/user").Use(middlewares.AuthCheck())
	{
		user.POST("/", controller.Create)
		user.POST("/token", middlewares.JWTTokenCheck(), controller.Token)
		user.POST("/login", controller.Login)
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
