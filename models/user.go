package models

import (
	"linqiurong2021/gin-book-frontend/mysql"

	"gorm.io/gorm"
)

// User 用户
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"name"`
	Password string `json:"password" gorm:"password"`
	Phone    string `json:"phone" gorm:"phone"`
	Cart     *Cart
}

// 通过ID获取用户信息
func GetUserByID(userID uint) (user *User, err error) {
	if err := mysql.DB.Where("id=?", userID).Find(&user); err != nil {
		return nil, err.Error
	}
	return
}

// 通过用户名或密码
func GetUserByNameAndEncryptPassword(userName string, encryptPassword string) (user *User, err error) {
	if err := mysql.DB.Debug().Where("user_name=?", userName).Where("password=?", encryptPassword).Find(&user); err != nil {
		return nil, err.Error
	}
	return
}

// 更新数据
func UpdateUser(info *User) (user *User, err error) {
	if err := mysql.DB.Debug().Where("id=?", info.ID).Save(info); err != nil {
		return nil, err.Error
	}
	return
}

// 通过ID删除用户
func DeleteUserByID(userID int) (user *User, err error) {
	if err := mysql.DB.Debug().Where("id=?", userID).Delete(&User{}); err != nil {
		return nil, err.Error
	}
	return
}
