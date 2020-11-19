package controller

import (
	"linqiurong2021/gin-book-frontend/logic"
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create 新增用户
func Create(c *gin.Context) {
	ok, err := logic.CreateUser(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusAccepted, utils.CreateFailure(err.Error(), ""))
			return
		}
		return
	}
}

// Update 更新
func Update(c *gin.Context) {
	ok, err := logic.UpdateUser(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// Delete 删除
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Success("Delete success", ""))
}

// ListByPage 删除
func ListByPage(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Success("ListByPage success", ""))
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
	logic.Logout()
}
