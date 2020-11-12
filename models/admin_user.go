package models

import "gorm.io/gorm"

// AdminUser 管理员
type AdminUser struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"user_name"`
	Password string `json:"password" gorm:"password"`
	Email    string `json:"email" gorm:"email"`
	Phone    string `json:"phone" gorm:"phone"`
}
