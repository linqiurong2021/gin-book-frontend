package dao

// Login 登录表单
type Login struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Code     string `json:"code"`
}
