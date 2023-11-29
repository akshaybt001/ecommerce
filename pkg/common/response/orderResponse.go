package response


type AdminOrder struct {
	OrderID       int    `gorm:"column:order_id"`
	PaymentTypeID int    `gorm:"column:payment_type_id"`
	OrderStatus   string `gorm:"column:order_status"`
	PaymentType   string `gorm:"column:payment_type"`
}
