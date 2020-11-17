package models

// AdminUser 管理员
type AdminUser struct {
	Common   `gorm:"embedded"`
	UserName string `json:"user_name" gorm:"user_name"`
	Password string `json:"password" gorm:"password"`
	Email    string `json:"email" gorm:"email"`
	Phone    string `json:"phone" gorm:"phone"`
}
