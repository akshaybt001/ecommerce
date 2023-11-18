package domain

import "time"

type PaymentType struct {
	Id   uint   `gorm:"primaryKey;unique;not null"`
	Type string `gorm:"unique;not null"`
}

type Orders struct {
	Id              uint `gorm:"primaryKey;unique;not null"`
	UserId          uint
	Users           Users `gorm:"foreignKey:UserId" json:"-"`
	OrderDate       time.Time
	PaymentTypeId   uint
	PaymentType     PaymentType `gorm:"foreignKey:PaymentTypeId" json:"-"`
	ShippingAddress uint
	OrderTotal      int
	OrderStatusID   uint
	OrderStatus     OrderStatus `gorm:"foreignKey:OrderStatusID" json:"-"`
}
type OrderItem struct {
	Id       uint `gorm:"primaryKey;unique;not null"`
	OrdersId uint
	Orders   Orders `gorm:"foreignKey:OrdersId" json:"-"`
	ModelId  uint
	Model    Model `gorm:"foreignKey:ModelId" json:"-"`
	Quantity int
	Price    int
}

type OrderStatus struct {
	Id     uint   `gorm:"primaryKey;unique;not null"`
	Status string `grom:"unique"`
}