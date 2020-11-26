package services

import (
	"linqiurong2021/gin-book-frontend/models"
)

// CreateOrder 创建订单
func CreateOrder(inOrder *models.Order) (outOrder *models.Order, err error) {
	return models.CreateOrder(inOrder)
}

// GetOrderByUserIDAndID 通过订单ID获取
func GetOrderByUserIDAndID(userID uint, OrderID uint) (Order *models.Order, err error) {
	return models.GetOrderByUserIDAndID(userID, OrderID)
}

// UpdateOrderByIDAndState 更新订单状态
func UpdateOrderByIDAndState(userID uint, orderID uint, status uint) (bool, error) {
	return models.UpdateOrderByIDAndState(userID, orderID, status)
}

// DeleteOrderByID 通过ID删除订单
func DeleteOrderByID(userID uint, OrderID int) (Order *models.Order, err error) {

	return models.DeleteOrderByID(userID, OrderID)
}

// GetListOrderByPageAndUserID 通过ID删除订单
func GetListOrderByPageAndUserID(userID uint, page int, pageSize int) (OrderList []*models.Order, count int64, err error) {
	OrderList, count, err = models.GetListOrderByPageAndUserID(userID, page, pageSize)
	return
}

// GetOrderStatusByString 通过字符串获取订单数值
func GetOrderStatusByString(status string) uint {
	switch status {
	case "create":
		return 0
	case "paid":
		return 1
	case "cancel":
		return 2
	case "refund":
		return 3
	case "refunded":
		return 4
	default:
		return 0
	}
}
