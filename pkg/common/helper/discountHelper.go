package helper

import "time"

type Discount struct {
	Id                    uint `gorm:"primarykey;unique;not null"`
	Brand_id              uint
	Brand                 Brands    `gorm:"foreignKey:Brand_id"`
	DiscountPercent       int       `json:"discountpercent"`
	DiscountMaximumAmount int       `json:"discountmaximum"`
	MinimumPurchaseAmount int       `json:"minimumpurchase"`
	ExpirationDate        time.Time `json:"expirationdate"`
}
