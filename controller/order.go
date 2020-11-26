package controller

import (
	"net/http"

	"github.com/linqiurong2021/gin-book-frontend/logic"
	"github.com/linqiurong2021/gin-book-frontend/utils"

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

// OrderStatus 修改状态
func OrderStatus(c *gin.Context) {
	ok, err := logic.OrderStatus(c)
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
