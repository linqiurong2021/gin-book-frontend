package controller

import (
	"linqiurong2021/gin-book-frontend/logic"
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBook 新增书籍
func CreateBook(c *gin.Context) {
	ok, err := logic.CreateBook(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusAccepted, utils.CreateFailure(err.Error(), ""))
			return
		}
		return
	}
}

// UpdateBook 更新
func UpdateBook(c *gin.Context) {
	ok, err := logic.UpdateBook(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// DeleteBook 删除
func DeleteBook(c *gin.Context) {
	ok, err := logic.DeleteBook(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// ListBookByPage 列表分页
func ListBookByPage(c *gin.Context) {
	logic.ListBookByPage(c)
	return
}
