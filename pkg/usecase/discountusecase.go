package usecase

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	interfaces "main.go/pkg/repository/interface"
	services "main.go/pkg/usecase/interface"
)

type DiscountUsecase struct {
	discountrepo interfaces.DiscountRepository
}

func NewDiscountUsecase(discountrepo interfaces.DiscountRepository) services.DiscountUsecase {
	return &DiscountUsecase{
		discountrepo: discountrepo,
	}
}

// AddDiscount implements interfaces.DiscountUsecase.
func (c *DiscountUsecase) AddDiscount(discount helper.Discount) error {
	err := c.discountrepo.AddDiscount(discount)
	return err
}

// EditDiscount implements interfaces.DiscountUsecase.
func (c *DiscountUsecase) EditDiscount(id int, discount helper.Discount) (response.Discount, error) {
	updatedDiscount, err := c.discountrepo.EditDiscount(id, discount)
	return updatedDiscount, err
}

// DeleteDiscount implements interfaces.DiscountUsecase.
func (c *DiscountUsecase) DeleteDiscount(id int) error {
	err := c.discountrepo.DeleteDiscount(id)
	return err
}

// ListAllDiscount implements interfaces.DiscountUsecase.
func (c *DiscountUsecase) ListAllDiscount() ([]response.Discount, error) {
	discount, err := c.discountrepo.ListAllDiscount()
	return discount, err
}

// ListDiscount implements interfaces.DiscountUsecase.
func (c *DiscountUsecase) ListDiscount(id int) (response.Discount, error) {
	discount, err := c.discountrepo.ListDiscount(id)
	return discount, err
}
