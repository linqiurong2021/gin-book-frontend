package order

// 'create' 'paid' 'cancle' 'refund' 'refunded'

// OrderStatus 订单状态
// type OrderStatus struct {
// 	Create   string `json:"create"`
// 	Paid     string `json:"paid"`
// 	Cancel   string `json:"cancel"`
// 	Refund   string `json:"refund"`
// 	Refunded string `json:"refunded"`
// }
//
const (
	// Create 创建状态
	Create = "create"
	// Paid 已支付
	Paid = "paid"
	// Cancel 已取消
	Cancel = "cancel"
	// Refund 已支付
	Refund = "refund"
	// Refunded 已退款
	Refunded = "refunded"
)
