package usecase

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	interfaces "main.go/pkg/repository/interface"
	services "main.go/pkg/usecase/interface"
)

type ProductUsecase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUsecase(productRepo interfaces.ProductRepository) services.ProductUsecase {
	return &ProductUsecase{
		productRepo: productRepo,
	}
}

// -------------------------- Create-Category --------------------------//
func (c *ProductUsecase) CreateCategory(category helper.Category) (response.Category, error) {
	newCategory, err := c.productRepo.CreateCategory(category)
	return newCategory, err
}

// -------------------------- Update-Category --------------------------//

func (c *ProductUsecase) UpdateCategory(category helper.Category, id int) (response.Category, error) {
	updatedCategory, err := c.productRepo.UpdateCategory(category, id)
	return updatedCategory, err
}

// -------------------------- Delete-Category --------------------------//

func (c *ProductUsecase) DeleteCategory(id int) error {
	err := c.productRepo.DeleteCategory(id)
	return err
}

// -------------------------- List-All-Category --------------------------//

func (c *ProductUsecase) ListAllCategories() ([]response.Category, error) {
	categories, err := c.productRepo.ListAllCategories()
	return categories, err
}

// -------------------------- List-Single-Category --------------------------//

func (c *ProductUsecase) ListCategory(id int) (response.Category, error) {
	category, err := c.productRepo.ListCategory(id)
	return category, err
}