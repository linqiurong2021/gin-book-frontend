package logic

import (
	"errors"
	"fmt"
	"linqiurong2021/gin-book-frontend/cached"
	"linqiurong2021/gin-book-frontend/dao"
	"linqiurong2021/gin-book-frontend/models"
	"linqiurong2021/gin-book-frontend/services"
	"linqiurong2021/gin-book-frontend/utils"
	"linqiurong2021/gin-book-frontend/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateCartItem 创建用户
func CreateCartItem(c *gin.Context) (ok bool, err error) {
	// 获取书籍ID
	var cartItem models.CartItem
	c.BindJSON(&cartItem)
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	// 通过UserID获取购物车ID
	cart, err := services.GetCartByUserID(cached.User.ID)
	if err != nil {
		// 判断是否是空值
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		// 创建购物车
		cart, err = services.CreateCart(&models.Cart{UserID: cached.User.ID, TotalCount: 0, TotalAmount: 0})
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
	}
	// 获取数据
	book, err := services.GetBookByID(cartItem.BookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
	// 绑定CartID
	cartItem.CartID = cart.ID

	// 修改购物车总数量
	cart.TotalCount = cart.TotalCount + cartItem.Count
	// 修改购物车总价
	cart.TotalAmount = cart.TotalAmount + (book.Price * float32(cartItem.Count))

	// 查库是否存在 如果存在则更新数量与总价
	item, err := services.GetCartItemByBookIDAndCartID(cartItem.BookID, cart.ID)

	if err != nil {
		// 判断是否是空值
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		cartItem.Amount = book.Price
		// 新增项
		err = services.CreateCartItemAndUpdateCart(&cartItem, cart)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		c.JSON(http.StatusOK, utils.Success("add success", ""))
		return true, nil
	}
	// 修改金额
	item.Amount = item.Amount + book.Price
	// 修改数量
	item.Count = item.Count + cartItem.Count

	fmt.Printf("\n#### %#v ####\n", item)
	// 更新成功
	err = services.UpdateCartItemAndUpdateCart(item, cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Success("add success", ""))
	// 不存在则新增
	return true, nil
}

// Increase 添加数量
func Increase(c *gin.Context) (ok bool, err error) {
	var newCartItem models.CartItem
	c.BindJSON(&newCartItem)
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	// 获取书籍信息
	book, err := services.GetBookByID(newCartItem.BookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	// 判断书籍库存
	if UnderStock(book, newCartItem.Count) {
		c.JSON(http.StatusBadRequest, utils.BadRequest("under stock !", ""))
		return false, err
	}

	// 获取购物车
	cart, err := services.GetCartByUserID(cached.User.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	// 获取当前项
	item, err := services.GetCartItemByBookIDAndCartID(newCartItem.BookID, cart.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	// 添加
	err = increaseNumber(book.Price, newCartItem.Count, item, cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	c.JSON(http.StatusOK, utils.Success("increase success", ""))
	// 获取当前数量
	return true, nil
}

// Decrease 减少数量
func Decrease(c *gin.Context) (ok bool, err error) {
	var newCartItem models.CartItem
	c.BindJSON(&newCartItem)
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	// 获取书籍信息
	book, err := services.GetBookByID(newCartItem.BookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	// 获取购物车
	cart, err := services.GetCartByUserID(cached.User.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	// 获取当前项
	item, err := services.GetCartItemByBookIDAndCartID(newCartItem.BookID, cart.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	if MinNumber(item, newCartItem.Count) {
		c.JSON(http.StatusBadRequest, utils.BadRequest("at least one number!", ""))
		return false, err
	}
	err = decreaseNumber(book.Price, newCartItem.Count, item, cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	c.JSON(http.StatusOK, utils.Success("decrease success", ""))
	return true, nil
}

// increaseNumber 新增数量
func increaseNumber(price float32, number uint, cartItem *models.CartItem, cart *models.Cart) error {
	// 更新数据与金额
	cartItem.Amount = cartItem.Amount + price*float32(number)
	cartItem.Count = cartItem.Count + number
	// 更新购物车的数据与金额
	cart.TotalAmount = cart.TotalAmount + price*float32(number)
	cart.TotalCount = cart.TotalCount + number
	return services.UpdateCartItemAndUpdateCart(cartItem, cart)
}

// decreaseNumber 减少数据
func decreaseNumber(price float32, number uint, cartItem *models.CartItem, cart *models.Cart) error {
	// 更新数据与金额
	cartItem.Amount = cartItem.Amount - price*float32(number)
	cartItem.Count = cartItem.Count - number
	// 更新购物车的数据与金额
	cart.TotalAmount = cart.TotalAmount - price*float32(number)
	cart.TotalCount = cart.TotalCount - number
	return services.UpdateCartItemAndUpdateCart(cartItem, cart)
}

// UnderStock 库存不足
func UnderStock(book *models.Book, number uint) bool {
	return book.Stock < number
}

// MinNumber 校验数量
func MinNumber(cartItem *models.CartItem, number uint) bool {
	return cartItem.Count-number < 1
}

// DeleteCartItemByID 删除单项
func DeleteCartItemByID(c *gin.Context) (ok bool, err error) {
	// 获取ID
	var delete dao.ID
	c.BindUri(&delete)
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	//
	inItem, err := services.GetCartItemByID(uint(delete.ID))
	if err != nil {
		return false, err
	}
	//
	cart, err := services.GetCartByUserID(cached.User.ID)
	if err != nil {
		return false, err
	}
	// 更新数据
	cart.TotalAmount = cart.TotalAmount - inItem.Amount
	cart.TotalCount = cart.TotalCount - inItem.Count
	err = services.DeleteCartItem(inItem, cart)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if err != nil {
		return false, err
	}
	c.JSON(http.StatusOK, utils.Success("delete success", ""))
	return true, nil
}

// DeleteCartItemByIDs 删除多项
func DeleteCartItemByIDs(c *gin.Context) (ok bool, err error) {
	// 获取ID
	ids := c.Bind("ids")
	c.JSON(http.StatusOK, utils.Success("ok", ids))
	return true, nil
}

// ListCartItemByPage 用户列表分页
func ListCartItemByPage(c *gin.Context) {
	//
	var page dao.Page
	c.BindQuery(&page)
	//
	cart, err := services.GetCartByUserID(cached.User.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
	//
	list, total, err := services.GetListCartByPage(cart.ID, page.Page, page.PageSize)
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
