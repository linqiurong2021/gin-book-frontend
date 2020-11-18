package dao

// Login 登录表单
type Login struct {
	UserName string `json:"user_name" binding:"required,min=6,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Code     string `json:"code" binding:"required,len=6"`
}
