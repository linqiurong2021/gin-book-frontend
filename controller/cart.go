package controller

import (
	"linqiurong2021/gin-book-frontend/logic"
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddToCart 新增购物项
func AddToCart(c *gin.Context) {
	ok, err := logic.CreateCartItem(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusAccepted, utils.CreateFailure(err.Error(), ""))
			return
		}
		return
	}
}

// UpdateCartItem 更新
func UpdateCartItem(c *gin.Context) {
	ok, err := logic.UpdateCartItem(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// DeleteCartItem 删除
func DeleteCartItem(c *gin.Context) {
	ok, err := logic.DeleteCartItemByID(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// DeleteCartItems 删除
func DeleteCartItems(c *gin.Context) {
	ok, err := logic.DeleteCartItemByIDs(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// ListCartItemByPage 列表分页
func ListCartItemByPage(c *gin.Context) {
	logic.ListCartItemByPage(c)
	return
}
