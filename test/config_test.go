package test

import (
	"fmt"
	"linqiurong2021/gin-book-frontend/config"
	"testing"

	"github.com/bmizerany/assert"
)

func TestConfigInit(t *testing.T) {
	err := config.Init("../config/config.ini")
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Printf("config:%v\n", config.Conf)
	assert.Equal(t, 9000, config.Conf.Port)
}
