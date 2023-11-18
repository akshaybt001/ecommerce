package response

type AdminOrder struct {
	OrderId       uint
	PaymentTypeId uint
	OrderStatus   string
	PaymentStatus string
}
