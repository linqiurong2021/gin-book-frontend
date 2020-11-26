package services

import (
	"github.com/linqiurong2021/gin-book-frontend/models"
)

// CreateUser 创建用户
func CreateUser(inUser *models.User) (outUser *models.User, err error) {
	return models.CreateUser(inUser)
}

// GetUserByID 通过用户ID获取
func GetUserByID(userID uint) (user *models.User, err error) {
	return models.GetUserByID(userID)
}

// GetUserByPhone 通过手机号获取用户信息
func GetUserByPhone(phone string) (user *models.User, rowsAffected int64, err error) {
	return models.GetUserByPhone(phone)
}

// GetUserByName 通过名称获取用户信息
func GetUserByName(name string) (user *models.User, rowsAffected int64, err error) {
	return models.GetUserByName(name)
}

// GetUserByFieldValue 通过名称获取用户信息
func GetUserByFieldValue(field string, value string) (user *models.User, rowsAffected int64, err error) {
	return models.GetUserByFieldValue(field, value)
}

// GetUserByNameAndEncryptPassword 通过用户名或密码
func GetUserByNameAndEncryptPassword(userName string, encryptPassword string) (user *models.User, err error) {
	return models.GetUserByNameAndEncryptPassword(userName, encryptPassword)
}

// UpdateUser 更新数据
func UpdateUser(info *models.User) (user *models.User, err error) {

	return models.UpdateUser(info)
}

// DeleteUserByID 通过ID删除用户
func DeleteUserByID(userID int) (user *models.User, err error) {

	return models.DeleteUserByID(userID)
}

// GetListUserByPage 通过ID删除用户
func GetListUserByPage(page int, pageSize int) (userList []*models.User, count int64, err error) {
	userList, count, err = models.GetListUserByPage(page, pageSize)
	return
}

// GetListUser 通过ID删除用户
func GetListUser() (userList []*models.User, err error) {

	userList, err = models.GetListUser()
	return
}
