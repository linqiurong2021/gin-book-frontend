package logic

import (
	"fmt"
	"linqiurong2021/gin-book-frontend/config"
	"linqiurong2021/gin-book-frontend/dao"
	"linqiurong2021/gin-book-frontend/models"
	"linqiurong2021/gin-book-frontend/services"
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login 登录
func Login(c *gin.Context) string {
	//
	var login = new(dao.Login)
	c.BindJSON(&login)
	// 验证码校验
	if !CheckCode(login.Code) {
		c.JSON(http.StatusBadRequest, utils.BadRequest("code invalidate", ""))
		return ""
	}
	user, err := services.GetUserByNameAndEncryptPassword(login.UserName, MD5Encrypt(login.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return ""
	}
	if user == nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("user name or password invalidate", ""))
		return ""
	}
	//
	singString := JWTToken(user)

	return singString
}

// JWTToken  JSONWebToken
func JWTToken(user *models.User) string {
	//
	mySigningKey := []byte(config.Conf.JWTSignKey)
	// 自定义
	type MyClaims struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		jwt.StandardClaims
	}
	//
	claims := MyClaims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: config.Conf.TokenExpireMinutes * 60, // 设置的分钟之后过期
			Issuer:    "test",
		},
	}
	//
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singString, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", singString, err)
	return singString
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
func CreateUser(c *gin.Context) (err error) {
	//
	var user = new(models.User)
	if err := c.BindJSON(&user); err != nil {
		return err
	}
	// 判断名称是否存在
	exists, err := NameExists(c, user)
	if err != nil {
		return err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("name has exists", ""))
		return nil
	}
	// 判断手机号是否存在
	exists, err = PhoneExists(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("phone has exists", ""))
		return nil
	}
	//
	// 密码加密
	user.Password = MD5Encrypt(user.Password)
	// 新增
	outUser, err := services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return err
	}
	// 创建成功
	c.JSON(http.StatusBadRequest, utils.Success("create success", outUser))
	//
	return nil

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
