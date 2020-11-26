package controller

import (
	"net/http"

	"github.com/linqiurong2021/gin-book-frontend/logic"
	"github.com/linqiurong2021/gin-book-frontend/utils"

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

// CartIncrease 添加数量
func CartIncrease(c *gin.Context) {
	ok, err := logic.CartIncrease(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// CartDecrease 减少数量
func CartDecrease(c *gin.Context) {
	ok, err := logic.CartDecrease(c)
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
