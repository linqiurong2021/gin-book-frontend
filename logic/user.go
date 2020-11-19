package logic

import (
	"fmt"
	"linqiurong2021/gin-book-frontend/cached"
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
func NameExists(c *gin.Context, user *models.User, ID uint) (exists bool, err error) {
	// 判断名称是否已存在
	outUser, err := services.GetUserByName(user.Name)
	if err != nil {
		return false, err
	}

	fmt.Sprintf("%#v\n", outUser)
	//
	if outUser != nil {
		//
		if outUser.ID == ID {
			return false, nil
		}
		return true, nil
	}
	return false, nil
}

// PhoneExists 校验手机号是否存在
func PhoneExists(c *gin.Context, user *models.User, ID uint) (exists bool, err error) {
	// 判断手机号是否已存在
	outUser, err := services.GetUserByPhone(user.Phone)
	if err != nil {
		return false, err
	}
	if outUser != nil {
		if outUser.ID == ID {
			return false, nil
		}
		return true, nil
	}
	return false, nil
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) (ok bool, err error) {
	var user models.User
	err = c.ShouldBindJSON(&user) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	// 判断名称是否存在
	exists, err := NameExists(c, &user, cached.User.ID)
	if err != nil {
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("name has exists", ""))
		return false, nil
	}
	// 判断手机号是否存在
	exists, err = PhoneExists(c, &user, cached.User.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("phone has exists", ""))
	}
	//
	// 密码加密
	user.Password = MD5Encrypt(user.Password)
	// 新增
	outUser, err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusOK, utils.Success("create success", outUser))
	//
	return true, nil

}

// NameAndPhoneExists 名称与手机号校验
func NameAndPhoneExists(c *gin.Context, user *models.User, ID uint) (ok bool, err error) {
	// 判断名称是否存在
	exists, err := NameExists(c, user, ID)
	if err != nil {
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("name has exists", ""))
		return false, nil
	}
	// 判断手机号是否存在
	exists, err = PhoneExists(c, user, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("phone has exists", ""))
		return false, nil
	}
	return
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) (ok bool, err error) {

	var user models.User
	err = c.BindJSON(&user) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	// 判断名称是否存在
	exists, err := NameExists(c, &user, cached.User.ID)
	if err != nil {
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("name has exists", ""))
		return false, nil
	}
	// 判断手机号是否存在
	exists, err = PhoneExists(c, &user, cached.User.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("phone has exists", ""))
		return false, nil
	}
	user.ID = cached.User.ID
	//
	user.Password = MD5Encrypt(user.Password)
	outUser, err := models.UpdateUser(&user)
	if err != nil {
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusOK, utils.Success("update success", outUser))
	// 获取当前
	return true, nil
}

// Logout 退出登录
func Logout(c *gin.Context) {
	// 直接清除
	// 数据清除
	// 退出成功
	c.JSON(http.StatusOK, utils.Success("logout success", ""))
}

// ListUserByPage 用户列表分页
func ListUserByPage(c *gin.Context) {
	//
	var page dao.Page
	c.BindQuery(&page)
	//
	list, total := services.GetListUserByPage(page.Page, page.PageSize)
	listPage := &dao.ListPage{
		Total: total,
		List:  list,
	}

	c.JSON(http.StatusOK, utils.Success("get success", listPage))
	return
}

// ListUser 用户列表 不分页
func ListUser(c *gin.Context) {
	list := services.GetList()
	c.JSON(http.StatusOK, utils.Success("get list success", list))
	return
}
