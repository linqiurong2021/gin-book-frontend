package services

import (
	"linqiurong2021/gin-book-frontend/models"
)

// CreateBook 创建书籍
func CreateBook(inBook *models.Book) (outBook *models.Book, err error) {
	return models.CreateBook(inBook)
}

// GetBookByID 通过书籍ID获取
func GetBookByID(bookID uint) (book *models.Book, err error) {
	return models.GetBookByID(bookID)
}

// GetBookListByIDs 通过书籍ID获取
func GetBookListByIDs(bookIDs []uint) (books []*models.Book, err error) {
	return models.GetBookListByIDs(bookIDs)
}

// UpdateBook 更新数据
func UpdateBook(info *models.Book) (book *models.Book, err error) {

	return models.UpdateBook(info)
}

// DeleteBookByID 通过ID删除书籍
func DeleteBookByID(userID uint, bookID int) (err error) {
	err = models.DeleteBookByID(userID, bookID)
	return
}

// GetListBookByPage 通过ID删除书籍
func GetListBookByPage(page int, pageSize int) (bookList []*models.Book, count int64, err error) {
	bookList, count, err = models.GetListBookByPage(page, pageSize)
	return
}
