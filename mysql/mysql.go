package mysql

import (
	"fmt"
	"linqiurong2021/gin-book-frontend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB ORM
var DB *gorm.DB

// InitMySQL 初始化数据库连接
func InitMySQL(cfg *config.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB, cfg.Charset)
	fmt.Printf("init DB: %s \n", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	// 如果是非生产环境则使用Debug
	if !config.Conf.Release {
		DB = DB.Debug()
	}
	return
}
