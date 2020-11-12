package controller

import (
	"linqiurong2021/gin-book-frontend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 健康检查
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Success("pong", ""))
}
