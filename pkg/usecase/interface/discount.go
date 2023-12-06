package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
)

type DiscountUsecase interface {
	AddDiscount(discount helper.Discount)error
	EditDiscount(id int , discount helper.Discount)(response.Discount,error)
	DeleteDiscount(id int)error
	ListAllDiscount()([]response.Discount,error)
	ListDiscount(id int)(response.Discount,error)

}