package models

// Admin 用户
type Admin struct {
	Common   `gorm:"embedded"`
	UserName string `json:"user_name" gorm:"user_name;type:varchar(20);" binding:"required,min=6,max=20" label:"用户名" `
	Password string `json:"password" gorm:"password;type:varchar(32);" binding:"required,min=6,max=20" label:"密码"`
	Phone    string `json:"phone" gorm:"phone;type:char(11);" binding:"len=11" label:"手机号"`
}
