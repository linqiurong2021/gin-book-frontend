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
		// 需要校验的分组
		authGroup(v1)
		// 不需要校验
		noAuthGroup(v1)
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

// noAuthGroup 不需要登录校验
func noAuthGroup(version *gin.RouterGroup) {
	version.POST("/login", controller.Login)
	version.GET("/logout", controller.Logout)
	// 新增用户(注册)
	version.POST("/user", controller.CreateUser)
	//
	bookGroup(version)
}

// authGroup 需要登录校验
func authGroup(version *gin.RouterGroup) {
	// 用户
	userGroup(version)
	// 购物车
	cartGroup(version)
}

// UserGroup User路由
func userGroup(g *gin.RouterGroup) {
	// 中间件
	user := g.Group("/user").Use(middlewares.JWTTokenCheck())
	{
		// 修改
		user.PUT("", middlewares.ID(), controller.UpdateUser)
		// user.DELETE("", controller.Delete)
		// 分页校验 middlewares.Page()
		user.GET("", middlewares.Page(), controller.ListUserByPage)
		// 无分页列表
		user.GET("/list", controller.ListUser)
	}
}

// BookGroup Book路由
func bookGroup(g *gin.RouterGroup) {
	book := g.Group("/book")
	{
		// 新增
		book.POST("", controller.CreateBook)
		// 修改
		book.PUT("", controller.UpdateBook)
		// 删除
		book.DELETE("/:id", middlewares.ID(), controller.DeleteBook)
		// 分页校验 middlewares.Page()
		book.GET("", middlewares.Page(), controller.ListBookByPage)
	}
}

// CartGroup 用户路由
func cartGroup(g *gin.RouterGroup) {
	g.Group("/cart")
	{

	}
}
