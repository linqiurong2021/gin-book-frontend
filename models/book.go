package models

import (
	"linqiurong2021/gin-book-frontend/mysql"
)

// Book 书集
type Book struct {
	Common  `gorm:"embedded"`
	Title   string  `json:"title" gorm:"title;type:varchar(30);" binding:"required,min=1,max=30" label:"书籍名称"`
	Author  string  `json:"author" gorm:"author;type:varchar(30);"  binding:"required,min=1,max=30" label:"书籍作者"`
	Price   float32 `json:"price" gorm:"price;type:float(7,2);"  binding:"required,numeric,gte=0" label:"价格"`
	Sales   uint    `json:"sales" gorm:"sales;type:int;" binding:"numeric,gte=0" label:"销售量"`
	Stock   uint    `json:"stock" gorm:"stock;type:int;" binding:"numeric,gte=0" label:"库存量"`
	ImgPath string  `json:"img_path" gorm:"img_path;type:varchar(255)"  label:"图片"`
}

// CreateBook 创建书
func CreateBook(inBook *Book) (book *Book, err error) {
	if err := mysql.DB.Create(&inBook).Error; err != nil {
		return nil, err
	}
	book = inBook
	return
}

// GetBookByID 通过ID获取书籍信息
func GetBookByID(bookID uint) (outBook *Book, err error) {
	var book = new(Book)
	if err := mysql.DB.Where("id = ?", bookID).First(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// UpdateBook 更新数据
func UpdateBook(info *Book) (outBook *Book, err error) {
	if err := mysql.DB.Debug().Where("id = ?", info.ID).Save(info).Error; err != nil {
		return nil, err
	}
	outBook = info
	return
}

// DeleteBookByID 通过ID删除书籍
func DeleteBookByID(bookID int) (err error) {
	if err := mysql.DB.Debug().Where("id = ?", bookID).Delete(&Book{}).Error; err != nil {
		return err
	}
	return nil
}

// GetListBookByPage 获取列表 分页
func GetListBookByPage(page int, pageSize int) (bookList []*Book, count int64, err error) {
	if err := mysql.DB.Debug().Offset((page - 1) * pageSize).Limit(pageSize).Find(&bookList).Error; err != nil {
		return nil, 0, err
	}
	if err := mysql.DB.Debug().Find(&Book{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return bookList, count, nil
}

// GetListBook 获取列表 不分页
func GetListBook() (bookList []*Book, err error) {
	if err := mysql.DB.Debug().Find(&bookList).Error; err != nil {
		return nil, err
	}
	return bookList, nil
}
