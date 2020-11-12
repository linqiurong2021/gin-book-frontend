package logic

import (
	"linqiurong2021/gin-book-frontend/dao"
	"linqiurong2021/gin-book-frontend/models"
	"linqiurong2021/gin-book-frontend/services"
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 登录
func Login(c *gin.Context) (user *models.User) {
	//
	var login = new(dao.Login)
	c.BindJSON(&login)
	// 验证码校验
	if !CheckCode(login.Code) {
		c.JSON(http.StatusBadRequest, utils.BadRequest("code invalidate", ""))
	}
	// 用户校验
	var err error = nil
	//
	user, err = services.GetUserByNameAndEncryptPassword(login.UserName, MD5Encrypt(login.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
	}
	return user
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) bool {
	//
	
	return true
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) bool {
	return true
}

// Logout 退出登录
func Logout() {
	// 直接清除
}

// GetUserByID 获取用户信息
func GetUserByID(userID uint) bool {
	return true
}
