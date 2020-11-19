package controller

import (
	"linqiurong2021/gin-book-frontend/logic"
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser 新增用户
func CreateUser(c *gin.Context) {
	ok, err := logic.CreateUser(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusAccepted, utils.CreateFailure(err.Error(), ""))
			return
		}
		return
	}
}

// UpdateUser 更新
func UpdateUser(c *gin.Context) {
	ok, err := logic.UpdateUser(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// DeleteUser 删除
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Success("Delete success", ""))
}

// ListUserByPage 列表分页
func ListUserByPage(c *gin.Context) {
	logic.ListUserByPage(c)
	return
}

// ListUser 列表
func ListUser(c *gin.Context) {
	logic.ListUser(c)
	return
}

// Login 登录
func Login(c *gin.Context) {
	singString, ok := logic.Login(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, utils.Success("login success", singString))
}

// Logout 退出登录
func Logout(c *gin.Context) {
	logic.Logout(c)
}
