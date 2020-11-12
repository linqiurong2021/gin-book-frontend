package models

import "gorm.io/gorm"

// Book 书集
type Book struct {
	gorm.Model
	Title   string  `json:"title" gorm:"title"`
	Auth    string  `json:"auth" gorm:"auth"`
	Price   float32 `json:"price" gorm:"price"`
	Sales   uint    `json:"sales" gorm:"sales"`
	Stock   uint    `json:"stock" gorm:"stock"`
	ImgPath string  `json:"img_path" gorm:"img_path"`
}
