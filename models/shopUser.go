package models

// ShopUser 店铺管理人
type ShopUser struct {
	Common   `gorm:"embedded"`
	UserName string `json:"user_name" gorm:"user_name;type:varchar(20);" binding:"required,min=6,max=20" label:"用户名" `
	Password string `json:"password" gorm:"password;type:varchar(32);" label:"密码"`
	ShopID   uint   `json:"shop_id" gorm:"shop_id" label:"店铺ID"`
}
