package logic

import (
	"linqiurong2021/gin-book-frontend/dao"
	"linqiurong2021/gin-book-frontend/models"
	"linqiurong2021/gin-book-frontend/services"
	"linqiurong2021/gin-book-frontend/utils"
	"linqiurong2021/gin-book-frontend/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBook 创建书籍
func CreateBook(c *gin.Context) (ok bool, err error) {
	var book models.Book
	err = c.ShouldBindJSON(&book) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	// 新增
	outBook, err := services.CreateBook(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusOK, utils.Success("create success", outBook))
	//
	return true, nil

}

// UpdateBook 更新书籍
func UpdateBook(c *gin.Context) (ok bool, err error) {

	var book models.Book
	err = c.BindJSON(&book) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}

	outBook, err := models.UpdateBook(&book)
	if err != nil {
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusOK, utils.Success("update success", outBook))
	// 获取当前
	return true, nil
}

// DeleteBook 删除书籍
func DeleteBook(c *gin.Context) (ok bool, err error) {
	var delete dao.ID
	c.BindUri(&delete)
	err = services.DeleteBookByID(delete.ID)
	if err != nil {
		// 删除失败
		c.JSON(http.StatusOK, utils.Success(err.Error(), ""))
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusOK, utils.Success("delete success", ""))
	// 获取当前
	return true, nil
}

// ListBookByPage 书籍列表分页
func ListBookByPage(c *gin.Context) {
	//
	var page dao.Page
	c.BindQuery(&page)
	//
	list, total, err := services.GetListBookByPage(page.Page, page.PageSize)
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
