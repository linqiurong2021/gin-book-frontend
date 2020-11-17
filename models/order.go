package models

// Order  订单
type Order struct {
	Common      `gorm:"embedded"`
	TotalCount  uint    `json:"total_count" gorm:"total_count"`
	TotalAmount float32 `json:"total_amount" gorm:"total_amount"`
	State       uint    `json:"state" gorm:"state"`
	UserID      uint
	User        User
	OrderItem   []OrderItem
}

// OrderItem 购物车每项
type OrderItem struct {
	Common  `gorm:"embedded"`
	Count   uint    `json:"count" gorm:"count"`
	Amount  float32 `json:"amount" gorm:"amount"`
	Title   string  `json:"title" gorm:"title"`
	Auth    string  `json:"auth" gorm:"auth"`
	Price   float32 `json:"price" gorm:"price"`
	ImgPath string  `json:"img_path" gorm:"img_path"`
	OrderID string  `json:"order_id" gorm:"order_id"`
	Order   Order
}
