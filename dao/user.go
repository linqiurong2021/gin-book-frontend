package dao

// User 用户
type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
