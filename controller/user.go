package controller

import (
	"linqiurong2021/gin-book-frontend/logic"
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create 新增用户
func Create(c *gin.Context) {
	err := logic.CreateUser(c)
	if err != nil {
		c.JSON(http.StatusAccepted, utils.CreateFailure(err.Error(), ""))
		return
	}
}

// Login 登录
func Login(c *gin.Context) {
	singString := logic.Login(c)
	if singString != "" {
		c.JSON(http.StatusOK, utils.Success("login success", singString))
	} else {
		c.JSON(http.StatusOK, utils.BadRequest("create token failure", ""))
	}
}

// Token 校验测试
func Token(c *gin.Context) {
	c.JSON(http.StatusOK, utils.BadRequest("token validate", ""))
}
