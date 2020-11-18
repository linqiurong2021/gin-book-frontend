package logic

import (
	"linqiurong2021/gin-book-frontend/dao"
	"linqiurong2021/gin-book-frontend/models"
	"linqiurong2021/gin-book-frontend/myjwt"
	"linqiurong2021/gin-book-frontend/services"
	"linqiurong2021/gin-book-frontend/utils"
	"linqiurong2021/gin-book-frontend/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 登录
func Login(c *gin.Context) (string, bool) {
	//
	// var login = new(dao.Login) // binding 校验无效
	var login dao.Login       // binding 校验有效
	err := c.BindJSON(&login) // 绑定并校验
	// 参数校验判断
	ok := validator.Validate(c, err)
	if !ok {
		return "", false
	}
	// 验证码校验
	if !CheckCode(login.Code) {
		c.JSON(http.StatusBadRequest, utils.BadRequest("code invalidate", ""))
		return "", false
	}
	user, err := services.GetUserByNameAndEncryptPassword(login.UserName, MD5Encrypt(login.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return "", false
	}
	if user == nil {
		c.JSON(http.StatusForbidden, utils.Forbidden("user name or password invalidate", ""))
		return "", false
	}
	//
	singString, err := JWTToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return "", false
	}

	return singString, true
}

// JWTToken  JSONWebToken
func JWTToken(user *models.User) (string, error) {
	//
	return myjwt.Create(user)
}

// NameExists 校验用户名是否存在
func NameExists(c *gin.Context, user *models.User) (exists bool, err error) {
	// 判断名称是否已存在
	outUser, err := services.GetUserByName(user.Name)
	if err != nil {
		return false, err
	}
	if outUser != nil {
		return true, nil
	}
	return false, nil
}

// PhoneExists 校验手机号是否存在
func PhoneExists(c *gin.Context, user *models.User) (exists bool, err error) {
	// 判断手机号是否已存在
	outUser, err := services.GetUserByPhone(user.Phone)
	if err != nil {
		return false, err
	}
	if outUser != nil {
		return true, nil
	}
	return false, nil
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) (ok bool, err error) {
	var userCreate dao.UserCreate
	err = c.ShouldBindJSON(&userCreate) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	var user = new(models.User)
	// 判断名称是否存在
	exists, err := NameExists(c, user)
	if err != nil {
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("name has exists", ""))
		return false, nil
	}
	// 判断手机号是否存在
	exists, err = PhoneExists(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("phone has exists", ""))
		return false, nil
	}
	//
	// 密码加密
	user.Password = MD5Encrypt(user.Password)
	// 新增
	outUser, err := services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusBadRequest, utils.Success("create success", outUser))
	//
	return true, nil

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
