package models

import "linqiurong2021/gin-book-frontend/mysql"

// Cart 购物车
type Cart struct {
	Common      `gorm:"embedded"`
	UserID      uint    `json:"user_id"`
	TotalCount  uint    `json:"total_count" gorm:"total_count"`
	TotalAmount float32 `json:"total_amount" gorm:"total_amount"`
	User        User
	CartItem    []CartItem
}

// CartItem 购物车每项
type CartItem struct {
	Common `gorm:"embedded"`
	Count  uint    `json:"count" gorm:"count"`
	Amount float32 `json:"amount" gorm:"amount"`
	CartID uint    `json:"cart_id" gorm:"cart_id"`
	BookID uint    `json:"book_id" gorm:"book_id"`
	Cart   Cart
	Book   Book
}

// CreateCartItem 加入购物车
func CreateCartItem(inItem *CartItem) (item *CartItem, err error) {

	if err := mysql.DB.Create(&inItem).Error; err != nil {
		return nil, err
	}
	item = inItem
	return
}

// GetCartItemByID 获取某项
func GetCartItemByID(itemID uint) (item *CartItem, err error) {
	if err := mysql.DB.Where("id = ?", itemID).Find(&item).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateCartItem 更新购物车项
func UpdateCartItem(item *CartItem) (outItem *CartItem, err error) {
	if err := mysql.DB.Debug().Where("id = ?", item.ID).Save(item).Error; err != nil {
		return nil, err
	}
	outItem = item
	return
}

// DeleteCartItemByID 删除购物车项
func DeleteCartItemByID(itemID int) (user *CartItem, err error) {
	if err := mysql.DB.Debug().Where("id = ?", itemID).Delete(&CartItem{}).Error; err != nil {
		return nil, err
	}
	return
}

// DeleteCartItemsByIDs 删除多项
func DeleteCartItemsByIDs(itemIDs []int) (user *CartItem, err error) {
	if err := mysql.DB.Debug().Delete(&CartItem{}, itemIDs).Error; err != nil {
		return nil, err
	}
	return
}

// GetCartItemListByPage 分页
func GetCartItemListByPage(page int, pageSize int) (itemList []*CartItem, count int64, err error) {
	if err := mysql.DB.Debug().Offset((page - 1) * pageSize).Limit(pageSize).Find(&itemList).Error; err != nil {
		return nil, 0, err
	}
	if err := mysql.DB.Debug().Find(&CartItem{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return itemList, count, nil
}
