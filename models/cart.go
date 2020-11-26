package models

import (
	"linqiurong2021/gin-book-frontend/mysql"

	"gorm.io/gorm"
)

// Cart 购物车
type Cart struct {
	Common      `gorm:"embedded"`
	UserID      uint       `json:"user_id;not null" binding:"required" label:"用户ID"`
	TotalCount  uint       `json:"total_count" gorm:"total_count;default:0" binding:"" label:"总量"`
	TotalAmount float32    `json:"total_amount" gorm:"total_amount;default:0" binding:"" label:"总价"`
	User        User       `json:"user"`
	CartItem    []CartItem `json:"cart_items"`
}

// CartItem 购物车每项
type CartItem struct {
	Common `gorm:"embedded"`
	Count  uint    `json:"count" gorm:"count;default:1" binding:"required" label:"数量"`
	Amount float32 `json:"amount" gorm:"amount;type:float(9,2)" binding:"required" label:"价格"`
	CartID uint    `json:"cart_id" gorm:"cart_id;not null" binding:"required" label:"购物车ID"`
	BookID uint    `json:"book_id" gorm:"book_id;not null" binding:"required" label:"已知ID"`
}

// CreateCart 创建购物车
func CreateCart(inCart *Cart) (outCart *Cart, err error) {
	if err := mysql.DB.Create(&inCart).Error; err != nil {
		return nil, err
	}
	return inCart, nil
}

// GetCartByUserID 获取某项
func GetCartByUserID(userID uint) (outCart *Cart, err error) {
	var cart = new(Cart)
	if err := mysql.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

// UpdateCart 更新购物车
func UpdateCart(cart *Cart) (outCart *Cart, err error) {
	if err := mysql.DB.Debug().Where("id = ?", cart.ID).Save(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

// CreateCartItemAndUpdateCart 加入购物车
func CreateCartItemAndUpdateCart(inItem *CartItem, inCart *Cart) (err error) {
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := mysql.DB.Create(&inItem).Error; err != nil {
			return err
		}
		if err := mysql.DB.Debug().Where("id = ?", inCart.ID).Save(inCart).Error; err != nil {
			return err
		}
		return nil
	})
	return
}

// GetCartItemByBookIDAndCartID 获取某项
func GetCartItemByBookIDAndCartID(bookID uint, cartID uint) (outItem *CartItem, err error) {
	var item = new(CartItem)
	if err := mysql.DB.Where("cart_id = ?", cartID).Where("book_id = ?", bookID).First(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

// GetCartItemByID 获取某项
func GetCartItemByID(ID uint) (outItem *CartItem, err error) {
	var item = new(CartItem)
	if err := mysql.DB.Where("id = ?", ID).First(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

// UpdateCartItemAndUpdateCart 更新购物车项
func UpdateCartItemAndUpdateCart(item *CartItem, cart *Cart) (err error) {

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := mysql.DB.Debug().Where("id = ?", item.ID).Save(item).Error; err != nil {
			return err
		}
		if err := mysql.DB.Debug().Where("id = ?", cart.ID).Save(cart).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteCartItem 删除购物车项
func DeleteCartItem(userID uint, inItem *CartItem, cart *Cart) (err error) {
	//
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		//
		if err := mysql.DB.Debug().Where("user_id = ?", userID).Delete(inItem).Error; err != nil {
			return err
		}
		if err := mysql.DB.Debug().Where("user_id = ?", userID).Where("id = ?", cart.ID).Save(cart).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

// DeleteCartItemsByIDs 删除多项
func DeleteCartItemsByIDs(itemIDs []int) (err error) {
	if err := mysql.DB.Debug().Delete(&CartItem{}, itemIDs).Error; err != nil {
		return err
	}
	return nil
}

// GetCartItemListByPage 分页
func GetCartItemListByPage(cartID uint, page int, pageSize int) (itemList []*CartItem, count int64, err error) {
	if err := mysql.DB.Debug().Where("cart_id = ?", cartID).Offset((page - 1) * pageSize).Limit(pageSize).Find(&itemList).Error; err != nil {
		return nil, 0, err
	}
	if err := mysql.DB.Debug().Where("cart_id = ?", cartID).Find(&CartItem{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return itemList, count, nil
}
