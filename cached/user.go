package cached

// User 存储当前登录用户信息
var User CurrentUser

// CurrentUser 存储当前登录用户信息
type CurrentUser struct {
	ID   uint
	Name string
}

// Save 存储
func Save(token *Claims) {
	User.ID = token.ID
	User.Name = token.Name
}
