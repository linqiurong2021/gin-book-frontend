package main

import (
	"fmt"
	"linqiurong2021/gin-book-frontend/config"
	"linqiurong2021/gin-book-frontend/models"
	"linqiurong2021/gin-book-frontend/mysql"
	"linqiurong2021/gin-book-frontend/routers"
	"linqiurong2021/gin-book-frontend/validator"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 加载配置文件(这里可以使用默认的配置文件)
	if err := config.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	if err := mysql.InitMySQL(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 开启校验转换
	if err := validator.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	// 绑定数据表
	mysql.DB.AutoMigrate(&models.Book{}, &models.User{}, &models.Cart{}, &models.CartItem{}, &models.Order{}, &models.OrderItem{})
	// 注册路由
	routers.RegisterRouter(r)
	r.Run(":9001")
}
