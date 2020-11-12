package models

import "gorm.io/gorm"

// Cart 购物车
type Cart struct {
	gorm.Model
	UserID      uint    `json:"user_id"`
	TotalCount  uint    `json:"total_count" gorm:"total_count"`
	TotalAmount float32 `json:"total_amount" gorm:"total_amount"`
	User        User
	CartItem    []CartItem
}

// CartItem 购物车每项
type CartItem struct {
	gorm.Model
	Count  uint    `json:"count" gorm:"count"`
	Amount float32 `json:"amount" gorm:"amount"`
	CartID uint    `json:"cart_id" gorm:"cart_id"`
	BookID uint    `json:"book_id" gorm:"book_id"`
	Cart   Cart
	Book   Book
}
