package domain

import "time"

type PaymentDetails struct {
	ID              uint          `gorm:"primaryKey" json:"id,omitempty"`
	OrdersID        uint          `json:"order_id,omitempty"`
	Orders          Orders        `gorm:"foreignKey:OrdersID" json:"-"`
	OrderTotal      float64       `json:"order_total"`
	PaymentTypeID   uint          `json:"payment_method_id"`
	PaymentType     PaymentType   `gorm:"foreignKey:PaymentTypeID"`
	PaymentStatusID uint          `json:"payment_status_id,omitempty"`
	PaymentStatus   PaymentStatus `gorm:"foreignKey:PaymentStatusID" json:"-"`
	PaymentRef      string        `json:"payment_ref"`
	UpdatedAt       time.Time
}

type PaymentStatus struct {
	ID            uint   `gorm:"primaryKey" json:"id,omitempty"`
	PaymentStatus string `json:"payment_status,omitempty"`
}