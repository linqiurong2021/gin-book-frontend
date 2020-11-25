package controller

import (
	"linqiurong2021/gin-book-frontend/logic"
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrder 新增订单
func CreateOrder(c *gin.Context) {
	ok, err := logic.CreateOrder(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusAccepted, utils.CreateFailure(err.Error(), ""))
			return
		}
		return
	}
}

// OrderIncrease 添加数量
func OrderIncrease(c *gin.Context) {
	ok, err := logic.OrderIncrease(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// OrderDecrease 减少数量
func OrderDecrease(c *gin.Context) {
	ok, err := logic.OrderDecrease(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// ListOrderByPage 列表分页
func ListOrderByPage(c *gin.Context) {
	logic.ListOrderByPageAndUserID(c)
	return
}
