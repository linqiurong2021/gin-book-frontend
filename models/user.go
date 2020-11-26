package models

import (
	"errors"
	"fmt"

	"github.com/linqiurong2021/gin-book-frontend/mysql"

	"gorm.io/gorm"
)

// User 用户
type User struct {
	Common   `gorm:"embedded"`
	Name     string `json:"name" gorm:"name;type:varchar(20);" binding:"required,min=6,max=20" label:"用户名" `
	Password string `json:"password" gorm:"password;type:varchar(32);" binding:"required,min=6,max=20" label:"密码"`
	Phone    string `json:"phone" gorm:"phone;type:char(11);" binding:"len=11" label:"手机号"`
	Cart     *Cart  `json:"cart"`
}

// CreateUser 创建用户
func CreateUser(inUser *User) (user *User, err error) {

	if err := mysql.DB.Create(&inUser).Error; err != nil {
		return nil, err
	}
	user = inUser
	return
}

// GetUserByID 通过ID获取用户信息
func GetUserByID(userID uint) (user *User, err error) {
	if err := mysql.DB.Where("id = ?", userID).Find(&user).Error; err != nil {
		return nil, err
	}
	return
}

// GetUserByNameAndEncryptPassword 通过用户名或密码
func GetUserByNameAndEncryptPassword(userName string, encryptPassword string) (outUser *User, err error) {
	var user = new(User)
	record := mysql.DB.Where("name = ?", userName).Where("password = ?", encryptPassword).Find(&user)
	// 为空或查找数量为0时
	if errors.Is(record.Error, gorm.ErrRecordNotFound) || record.RowsAffected == 0 {
		return nil, nil
	}
	if record.Error != nil {
		// fmt.Print("BBBBB###")
		return nil, record.Error
	}
	return user, nil
}

// UpdateUser 更新数据
func UpdateUser(info *User) (outUser *User, err error) {
	if err := mysql.DB.Where("id = ?", info.ID).Save(info).Error; err != nil {
		return nil, err
	}
	outUser = info
	return
}

// DeleteUserByID 通过ID删除用户
func DeleteUserByID(userID int) (user *User, err error) {
	if err := mysql.DB.Where("id = ?", userID).Delete(&User{}).Error; err != nil {
		return nil, err
	}
	return
}

// GetUserByPhone 通过某个字段获取用户信息
func GetUserByPhone(value string) (outUser *User, rowsAffected int64, err error) {
	return GetUserByFieldValue("phone", value)
}

// GetUserByName 通过某个字段获取用户信息
func GetUserByName(value string) (outUser *User, rowsAffected int64, err error) {
	return GetUserByFieldValue("name", value)
}

// GetUserByFieldValue 通过某个字段获取用户信息
func GetUserByFieldValue(field string, value string) (outUser *User, rowsAffected int64, err error) {
	var user = new(User)
	var where string = fmt.Sprintf("%s = ?", field)
	//
	record := mysql.DB.Where(where, value).First(&user)
	// 查不到数据
	if errors.Is(record.Error, gorm.ErrRecordNotFound) {
		return nil, 0, record.Error
	}
	// 异常
	if record.Error != nil {
		return nil, 1, record.Error
	}

	return user, 1, nil
}

// GetListUserByPage 获取列表 分页 (暂用不到)
func GetListUserByPage(page int, pageSize int) (userList []*User, count int64, err error) {
	if err := mysql.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&userList).Error; err != nil {
		return nil, 0, err
	}
	if err := mysql.DB.Find(&User{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return userList, count, nil
}

// GetListUser 获取列表 不分页  (暂用不到)
func GetListUser() (userList []*User, err error) {
	if err := mysql.DB.Find(&userList).Error; err != nil {
		return nil, err
	}

	return userList, nil
}
