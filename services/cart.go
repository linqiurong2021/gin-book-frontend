package services

import "linqiurong2021/gin-book-frontend/models"

// CreateCart 创建商品项
func CreateCart(inCart *models.Cart) (outCart *models.Cart, err error) {
	return models.CreateCart(inCart)
}

// UpdateCart 更新数据
func UpdateCart(cart *models.Cart) (outCart *models.Cart, err error) {

	return models.UpdateCart(cart)
}

// GetCartByUserID 获取购物车
func GetCartByUserID(userID uint) (outCart *models.Cart, err error) {
	return models.GetCartByUserID(userID)
}

// CreateCartItemAndUpdateCart 创建商品项
func CreateCartItemAndUpdateCart(inCartItem *models.CartItem, inCart *models.Cart) (err error) {
	return models.CreateCartItemAndUpdateCart(inCartItem, inCart)
}

// GetCartItemByID 通过商品项ID获取
func GetCartItemByID(ID uint) (outCartItem *models.CartItem, err error) {
	return models.GetCartItemByID(ID)
}

// GetCartItemByBookIDAndCartID 通过商品项ID获取
func GetCartItemByBookIDAndCartID(bookID uint, cartID uint) (outCartItem *models.CartItem, err error) {
	return models.GetCartItemByBookIDAndCartID(bookID, cartID)
}

// UpdateCartItemAndUpdateCart 更新数据
func UpdateCartItemAndUpdateCart(info *models.CartItem, cart *models.Cart) (err error) {

	return models.UpdateCartItemAndUpdateCart(info, cart)
}

// DeleteCartItem 通过ID删除商品项
func DeleteCartItem(userID uint, inItem *models.CartItem, cart *models.Cart) (err error) {
	return models.DeleteCartItem(userID, inItem, cart)
}

// DeleteCartItemByIDs 通过ID删除商品项
func DeleteCartItemByIDs(itemIDs []int) (err error) {
	return models.DeleteCartItemsByIDs(itemIDs)
}

// GetListCartByPage 通过ID删除商品项
func GetListCartByPage(cartID uint, page int, pageSize int) (cartItemList []*models.CartItem, count int64, err error) {
	cartItemList, count, err = models.GetCartItemListByPage(cartID, page, pageSize)
	return
}
