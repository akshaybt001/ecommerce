package domain

import "time"

type Discount struct {
	Id                    uint `gorm:"primarykey;unique;not null"`
	Brand_id              uint
	Brand                 Brand `gorm:"foreignKey:Brand_id"`
	DiscountPercent       float64
	DiscountMaximumAmount int
	MinimumPurchaseAmount int
	ExpirationDate        time.Time
}
