package response

import "time"

type AdminOrder struct {
	OrderID       int    `gorm:"column:order_id"`
	PaymentTypeID int    `gorm:"column:payment_type_id"`
	OrderStatus   string `gorm:"column:order_status"`
	PaymentType   string `gorm:"column:payment_type"`
}
type OrderResponse struct {
	Id              uint
	OrderDate       time.Time
	UserId          uint
	UserName        string
	PaymentTypeId   uint
	PaymentType     string
	ShippingAddress uint
	Address         `gorm:"embedded" json:"ShippingAddress,omitempty"`
	ModelName       string
	OrderStatusID   uint
	OrderStatus     string
	PaymentStatusID uint `json:"payment_status_id,omitempty"`
	PaymentStatus   string
	OrderTotal      int
}