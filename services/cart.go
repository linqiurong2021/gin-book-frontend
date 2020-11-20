package services

import "linqiurong2021/gin-book-frontend/models"

// CreateCartItem 创建商品项
func CreateCartItem(inCart *models.CartItem) (outCartItem *models.CartItem, err error) {
	return models.CreateCartItem(inCart)
}

// GetCartItemByID 通过商品项ID获取
func GetCartItemByID(itemID uint) (outCartItem *models.CartItem, err error) {
	return models.GetCartItemByID(itemID)
}

// UpdateCartItem 更新数据
func UpdateCartItem(info *models.CartItem) (outCartItem *models.CartItem, err error) {

	return models.UpdateCartItem(info)
}

// DeleteCarItemtByID 通过ID删除商品项
func DeleteCarItemtByID(itemID int) (outCartItem *models.CartItem, err error) {
	return models.DeleteCartItemByID(itemID)
}

// DeleteCartItemByIDs 通过ID删除商品项
func DeleteCartItemByIDs(itemIDs []int) (outCartItem *models.CartItem, err error) {
	return models.DeleteCartItemsByIDs(itemIDs)
}

// GetListCartByPage 通过ID删除商品项
func GetListCartByPage(page int, pageSize int) (cartItemList []*models.CartItem, count int64, err error) {
	cartItemList, count, err = models.GetCartItemListByPage(page, pageSize)
	return
}
