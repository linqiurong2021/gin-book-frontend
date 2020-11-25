package dao

// Login 登录表单
type Login struct {
	UserName string `json:"user_name" binding:"required,min=6,max=20" label:"账号"`
	Password string `json:"password" binding:"required,min=6,max=20"  label:"密码"`
	Code     string `json:"code" binding:"required,len=6"  label:"验证码"`
}

// ToUserInfo 返回的结构体
type ToUserInfo struct {
	UserName string `json:"user_name"`
}

// Page 分页
type Page struct {
	Page     int `form:"page" binding:"required,numeric,gte=1" label:"页码"`       // 大于1 GET 需要用formtag
	PageSize int `form:"page_size" binding:"required,numeric,gte=10" label:"条数"` // 大于 10 GET 需要用formtag
}

// ID 删除操作与编辑操作
type ID struct {
	ID int `uri:"id" json:"id" binding:"required,numeric,gte=1" label:"ID"`
}

// IDs 多删除操作
type IDs struct {
	IDs []int `uri:"ids" json:"ids" binding:"required,array" label:"IDs"`
}

// ListPage 分页返回数据格式
type ListPage struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

// OrderItem 订单项
type OrderItem struct {
	BookID uint `json:"book_id" binding:"required" label:"ID"`
	Number uint `json:"number" binding:"required,gt=0" label:"数量"`
}

// OrderItems 多个订单项
type OrderItems struct {
	OrderItem []*OrderItem `json:"order_items" label:"订单项"`
}
