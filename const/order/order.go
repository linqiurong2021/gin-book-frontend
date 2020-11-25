package order

const (
	// Init 初始状态
	Init = 0
	// WaitPay 待支付
	WaitPay = 1
	// Paid 已支付
	Paid = 2
	// Refund 申请退款
	Refund = 3
	// Refunded 已退款
	Refunded = 4
)

//
