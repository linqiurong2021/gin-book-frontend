package logic

import (
	"linqiurong2021/gin-book-frontend/dao"
	"linqiurong2021/gin-book-frontend/services"
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCartItem 创建用户
func CreateCartItem(c *gin.Context) (ok bool, err error) {

	//
	return true, nil

}

// UpdateCartItem 更新用户
func UpdateCartItem(c *gin.Context) (ok bool, err error) {
	// 获取当前
	return true, nil
}

// DeleteCartItemByID 删除单项
func DeleteCartItemByID(c *gin.Context) (ok bool, err error) {

	return true, nil
}

// DeleteCartItemByIDs 删除多项
func DeleteCartItemByIDs(c *gin.Context) (ok bool, err error) {

	return true, nil
}

// ListCartItemByPage 用户列表分页
func ListCartItemByPage(c *gin.Context) {
	//
	var page dao.Page
	c.BindQuery(&page)
	//
	list, total, err := services.GetListCartByPage(page.Page, page.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
	listPage := &dao.ListPage{
		Total: total,
		List:  list,
	}

	c.JSON(http.StatusOK, utils.Success("get success", listPage))
	return
}
