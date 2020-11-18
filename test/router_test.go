package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"linqiurong2021/gin-book-frontend/config"
	"linqiurong2021/gin-book-frontend/dao"
	"linqiurong2021/gin-book-frontend/models"
	"linqiurong2021/gin-book-frontend/mysql"
	"linqiurong2021/gin-book-frontend/routers"
	"linqiurong2021/gin-book-frontend/validator"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	// "github.com/stretchr/testify/assert"
)

var r *gin.Engine

//
func init() {
	r = gin.Default()
	// 加载配置文件(这里可以使用默认的配置文件)
	config.Init("../config/config.ini")
	fmt.Printf("%v########config.Conf", config.Conf)
	if err := mysql.InitMySQL(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 开启校验转换
	if err := validator.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	// 启用Mysql
	mysql.DB.AutoMigrate(&models.Book{}, &models.User{}, &models.Cart{}, &models.CartItem{}, &models.Order{}, &models.OrderItem{})
	// 注册路由
	routers.RegisterRouter(r)
	r.Run(":9001")
}

func TestPingRoute(t *testing.T) {

	response := httptest.NewRecorder()
	pingRequest, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(response, pingRequest)

	assert.Equal(t, 200, response.Code)
	wantResp := "{\"code\":200,\"msg\":\"pong\",\"data\":\"\"}"
	fmt.Println("response body:", response.Body.String())
	assert.Equal(t, wantResp, response.Body.String())

}

func TestV1UserLoginRoute(t *testing.T) {

	response := httptest.NewRecorder()
	// 添加参数
	var data = new(dao.Login)
	data.UserName = "17605048999"
	data.Code = "123456"
	data.Password = "123456"

	bytesData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("strings.NewReader(data.Encode())", string(bytesData))
	//
	v1UserRequest, _ := http.NewRequest("POST", "/v1/user/login", bytes.NewReader(bytesData))
	v1UserRequest.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// v1UserRequest.Header.Add("Content-Length", strconv.Itoa(len(string(bytesData))))
	r.ServeHTTP(response, v1UserRequest)
	fmt.Println("response body:", response.Body.String())
	// 返回数据
	assert.Equal(t, 200, response.Code)

}
