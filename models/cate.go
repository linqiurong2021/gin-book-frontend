package models

// Cate 分类
type Cate struct {
	Common `gorm:"embedded"`
	Name   string `json:"name" gorm:"name;type:varchar(20);" binding:"required,min=6,max=20" label:"用户名" `
	Order  uint   `json:"order" gorm:"order" label:"排序"`
	Note   string `json:"note" gorm:"note;type:varchar(100)" binding:"required,max=100" label:"备注"`
}
