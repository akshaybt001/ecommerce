package repository

import (
	"fmt"

	"gorm.io/gorm"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	interfaces "main.go/pkg/repository/interface"
)

type DiscountDatabase struct {
	DB *gorm.DB
}

func NewDiscountRepository(DB *gorm.DB) interfaces.DiscountRepository {
	return &DiscountDatabase{DB}
}

// AddDiscount implements interfaces.DiscountRepository.
func (c *DiscountDatabase) AddDiscount(discount helper.Discount) error {
	add := `INSERT INTO discounts(brand_id,discount_Percent,discount_maximum_Amount,minimum_Purchase_Amount,expiration_date) VALUES (?, ?, ?, ?, ?)`
	err := c.DB.Exec(add, discount.Brand_id, discount.DiscountPercent, discount.DiscountMaximumAmount, discount.MinimumPurchaseAmount, discount.ExpirationDate).Error

	return err

}

// EditDiscount implements interfaces.DiscountRepository.
func (c *DiscountDatabase) EditDiscount(id int, discount helper.Discount) (response.Discount, error) {
	var updatedDiscount response.Discount
	query := `UPDATE discounts SET brand_id=$1,discount_percent=$2,discount_maximum_Amount=$3,minimum_Purchase_Amount=$4,expiration_date=$5 WHERE id=$6
		RETURNING id,brand_id,discount_percent,discount_maximum_Amount,minimum_Purchase_Amount,expiration_date`
	err := c.DB.Raw(query, discount.Brand_id, discount.DiscountPercent, discount.DiscountMaximumAmount, discount.MinimumPurchaseAmount, discount.ExpirationDate, id).
		Scan(&updatedDiscount).Error
	if err != nil {
		return response.Discount{}, err
	}
	if updatedDiscount.Id == 0 {
		return response.Discount{}, fmt.Errorf("there is no such discount id ")
	}
	return updatedDiscount, nil
}

// DeleteDiscount implements interfaces.DiscountRepository.
func (c *DiscountDatabase) DeleteDiscount(id int) error {
	var exists bool
	isExists := `SELECT EXISTS (SELECT 1 FROM discounts WHERE id=$1)`
	c.DB.Raw(isExists, id).Scan(&exists)
	if !exists {
		return fmt.Errorf("there is no such discount to delete")
	}
	query := `DELETE FROM discounts WHERE id=?`
	err := c.DB.Exec(query, id).Error
	return err
}

// ListAllDiscount implements interfaces.DiscountRepository.
func (c *DiscountDatabase) ListAllDiscount() ([]response.Discount, error) {

	var discount []response.Discount
	query := `SELECT * FROM discounts`
	err := c.DB.Raw(query).Scan(&discount).Error
	return discount, err

}

// ListDiscount implements interfaces.DiscountRepository.

func (c *DiscountDatabase) ListDiscount(id int) (response.Discount, error) {
	var discount response.Discount
	var exits bool

	query1 := `select exists(select 1 from discounts where id=?)`
	c.DB.Raw(query1, id).Scan(&exits)
	if !exits {
		return response.Discount{}, fmt.Errorf("no discount found")
	}
	query := `SELECT * FROM discounts WHERE id=$1`
	err := c.DB.Raw(query, id).Scan(&discount).Error
	if err != nil {
		return response.Discount{}, err
	}
	if discount.Id == 0 {
		return response.Discount{}, fmt.Errorf("no such discount")
	}
	return discount, nil
}
