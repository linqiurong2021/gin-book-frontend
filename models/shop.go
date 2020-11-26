package models

// Shop 店铺
type Shop struct {
	Common   `gorm:"embedded"`
	Name     string    `json:"name" gorm:"name;type:varchar(20);" binding:"required,min=6,max=20" label:"用户名" `
	Contact  string    `json:"contact" gorm:"contact;type:char(11);" binding:"len=11" label:"手机号"`
	ShopUser *ShopUser `json:"shop_user"`
	Books    []*Book   `json:"books"`
}
