package services

import (
	"linqiurong2021/gin-book-frontend/models"
)

// GetUserByID 通过用户ID获取
func GetUserByID(userID uint) (user *models.User, err error) {
	return models.GetUserByID(userID)
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
