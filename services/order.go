package services

import "linqiurong2021/gin-book-frontend/models"

// CreateOrder 创建订单
func CreateOrder(inOrder *models.Order) (outOrder *models.Order, err error) {
	return models.CreateOrder(inOrder)
}

// GetOrderByID 通过订单ID获取
func GetOrderByID(OrderID uint) (Order *models.Order, err error) {
	return models.GetOrderByID(OrderID)
}

// UpdateOrder 更新数据
func UpdateOrder(info *models.Order) (Order *models.Order, err error) {

	return models.UpdateOrder(info)
}

// DeleteOrderByID 通过ID删除订单
func DeleteOrderByID(OrderID int) (Order *models.Order, err error) {

	return models.DeleteOrderByID(OrderID)
}

// GetListOrderByPageAndUserID 通过ID删除订单
func GetListOrderByPageAndUserID(userID uint, page int, pageSize int) (OrderList []*models.Order, count int64, err error) {
	OrderList, count, err = models.GetListOrderByPageAndUserID(userID, page, pageSize)
	return
}
