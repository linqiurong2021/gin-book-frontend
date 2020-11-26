package logic

import (
	"net/http"

	"github.com/linqiurong2021/gin-book-frontend/cached"
	"github.com/linqiurong2021/gin-book-frontend/consts/order"
	"github.com/linqiurong2021/gin-book-frontend/dao"
	"github.com/linqiurong2021/gin-book-frontend/models"
	"github.com/linqiurong2021/gin-book-frontend/services"
	"github.com/linqiurong2021/gin-book-frontend/utils"
	"github.com/linqiurong2021/gin-book-frontend/validator"

	"github.com/gin-gonic/gin"
)

// CreateOrder 创建订单
func CreateOrder(c *gin.Context) (ok bool, err error) {
	var orderItems dao.OrderItems
	//
	err = c.ShouldBindJSON(&orderItems) // 绑定并校验

	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	// 保存书籍ID
	var bookIDs = make([]uint, len(orderItems.OrderItem))
	// 购买的数量
	var numbers = make(map[uint]uint)
	// 总数量
	var totalNumber uint
	// 总价格
	var totalPrice float32
	// 循环
	for key, item := range orderItems.OrderItem {
		//
		bookIDs[key] = item.BookID
		numbers[item.BookID] = item.Number
		totalNumber = totalNumber + item.Number
	}
	// 获取当前书籍数据
	bookList, err := getBooksByIDs(bookIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
	}

	// // 订单项
	var orderItem = make([]models.OrderItem, len(bookList))
	for index, item := range bookList {
		var number = numbers[item.ID]
		var amount = item.Price * float32(number)
		orderItem[index].Count = number
		orderItem[index].Amount = amount
		orderItem[index].Title = item.Title
		orderItem[index].Author = item.Author
		orderItem[index].Price = item.Price
		orderItem[index].ImgPath = item.ImgPath
		totalPrice = totalPrice + amount
	}
	// var orderInfo = new(models.Order)
	// orderInfo.UserID = cached.User.ID
	// orderInfo.State = order.Init // 订单初始化
	// orderInfo.OrderItem = orderItem
	// orderInfo.TotalCount = totalNumber
	// orderInfo.TotalAmount = totalPrice
	var orderInfo = &models.Order{
		UserID:      cached.User.ID,
		State:       services.GetOrderStatusByString(order.Create),
		OrderItem:   orderItem,
		TotalAmount: totalPrice,
		TotalCount:  totalNumber,
	}

	_, err = services.CreateOrder(orderInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
	}
	c.JSON(http.StatusOK, utils.Success("create order success", orderInfo))
	return true, nil
}

// 通过书籍ID获取书本信息
func getBooksByIDs(bookIDs []uint) (bookList []*models.Book, err error) {
	bookList, err = services.GetBookListByIDs(bookIDs)
	if err != nil {
		return nil, err
	}
	return
}

// ListOrderByPageAndUserID 创建订单
func ListOrderByPageAndUserID(c *gin.Context) {

	var page dao.Page

	c.BindQuery(&page) // 绑定并校验

	list, total, err := services.GetListOrderByPageAndUserID(cached.User.ID, page.Page, page.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
	// order.OrderItem = []
	listPage := &dao.ListPage{
		Total: total,
		List:  list,
	}
	c.JSON(http.StatusOK, utils.Success("get success", listPage))
	return
}

// OrderStatus 订单状态
func OrderStatus(c *gin.Context) (ok bool, err error) {
	//
	var orderStatus dao.OrderStatus
	//
	err = c.ShouldBindJSON(&orderStatus) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	// 获取当前状态ID
	order, err := services.GetOrderByUserIDAndID(cached.User.ID, orderStatus.OrderID)
	if err != nil {
		return false, err
	}
	orderState := services.GetOrderStatusByString(orderStatus.Status)
	if orderState <= order.State {
		c.JSON(http.StatusBadRequest, utils.BadRequest("order status can not rollback", ""))
		return
	}

	ok, err = services.UpdateOrderByIDAndState(cached.User.ID, orderStatus.OrderID, orderState)
	if err != nil {
		return false, err
	}
	c.JSON(http.StatusOK, utils.Success("update status success", ""))
	return ok, nil
}
