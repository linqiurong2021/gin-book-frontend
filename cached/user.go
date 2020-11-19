package cached

import (
	"linqiurong2021/gin-book-frontend/myjwt"
)

// User 存储当前登录用户信息
var User CurrentUser

// CurrentUser 存储当前登录用户信息
type CurrentUser struct {
	ID   uint
	Name string
}

// Save 存储
func Save(token *myjwt.MyClaims) {
	User.ID = token.ID
	User.Name = token.Name
}
