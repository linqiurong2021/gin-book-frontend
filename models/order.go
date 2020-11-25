package models

import "linqiurong2021/gin-book-frontend/mysql"

// 区别在于定义模型时，has one 关系，把附属者 CreditCard 作为所有者 User 的一个字段，更自然；而 belongs to 关系，把所有者 User 作为附属者 CreditCard 的一个字段。

// Order  订单
type Order struct {
	Common      `gorm:"embedded"`
	TotalCount  uint        `json:"total_count" gorm:"total_count" binding:"required,numeric,gte=0;" label:"总量"`
	TotalAmount float32     `json:"total_amount" gorm:"total_amount" binding:"required,numeric,gte=0;" label:"总价"`
	State       uint        `json:"state" gorm:"state"`
	UserID      uint        `json:"user_id"` // Belongs To
	OrderItem   []OrderItem `json:"order_items"`
}

// OrderItem 购物车每项
type OrderItem struct {
	Common  `gorm:"embedded"`
	Count   uint    `json:"count" gorm:"count" binding:"required,numeric,gte=0;"`
	Amount  float32 `json:"amount" gorm:"amount" binding:"required,numeric,gte=0;"`
	Title   string  `json:"title" gorm:"title" binding:"required,string;"`
	Author  string  `json:"auth" gorm:"auth" binding:"required"`
	Price   float32 `json:"price" gorm:"price" binding:"required,numeric,gte=0;"`
	ImgPath string  `json:"img_path" gorm:"img_path"`
	OrderID uint    `json:"order_id" gorm:"order_id"` // Belongs To
}

// CreateOrder 创建订单
func CreateOrder(inOrder *Order) (outOrder *Order, err error) {
	if err := mysql.DB.Create(&inOrder).Error; err != nil {
		return nil, err
	}
	outOrder = inOrder
	return
}

// UpdateOrder 更新订单
func UpdateOrder(inOrder *Order) (outOrder *Order, err error) {
	if err := mysql.DB.Create(&inOrder).Error; err != nil {
		return nil, err
	}
	outOrder = inOrder
	return
}

// GetOrderByID 通过ID获取用户信息
func GetOrderByID(orderID uint) (outOrder *Order, err error) {
	if err := mysql.DB.Where("id = ?", orderID).Find(&outOrder).Error; err != nil {
		return nil, err
	}
	return
}

// DeleteOrderByID 通过ID删除用户
func DeleteOrderByID(orderID int) (outOrder *Order, err error) {
	if err := mysql.DB.Debug().Where("id = ?", orderID).Delete(&Order{}).Error; err != nil {
		return nil, err
	}
	return
}

// GetListOrderByPageAndUserID 通过用户ID分页
func GetListOrderByPageAndUserID(userID uint, page int, pageSize int) (outOrderList []*Order, count int64, err error) {
	// 加载数据项
	if err := mysql.DB.Debug().Preload("OrderItem").Where("user_id = ?", userID).Offset((page - 1) * pageSize).Limit(pageSize).Find(&outOrderList).Error; err != nil {
		return nil, 0, err
	}
	if err := mysql.DB.Debug().Where("user_id = ?", userID).Find(&Order{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return outOrderList, count, nil
}
