package models

import "gorm.io/gorm"

// User 用户
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"name"`
	Password string `json:"password" gorm:"password"`
	Phone    string `json:"phone" gorm:"phone"`
	Cart     *Cart
}
