package main

import (
	"fmt"
	"linqiurong2021/gin-book-frontend/config"
	"linqiurong2021/gin-book-frontend/controller"
	"linqiurong2021/gin-book-frontend/models"
	"linqiurong2021/gin-book-frontend/mysql"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	// 加载配置文件(这里可以使用默认的配置文件)
	if err := config.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	if err := mysql.InitMySQL(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	mysql.DB.AutoMigrate(&models.Book{}, &models.User{}, &models.Cart{}, &models.CartItem{}, &models.Order{}, &models.OrderItem{})
	r.Run(":9001")
}
