package dao

// UserCreate 用户创建
type UserCreate struct {
	Name     string `json:"name" binding:"required,min=6,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Phone    string `json:"phone" binding:"required,len=11"`
}
