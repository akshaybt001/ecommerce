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

// -------------------------- Create-Brand --------------------------//

func (c *ProductUsecase) AddBrand(Brand helper.Brands) (response.Brands, error) {
	newBrand, err := c.productRepo.AddBrand(Brand)
	return newBrand, err
}

// -------------------------- Update-Brand --------------------------//

func (c *ProductUsecase) UpdateBrand(id int, Brand helper.Brands) (response.Brands, error) {
	updatedBrand, err := c.productRepo.UpdateBrand(id, Brand)
	return updatedBrand, err
}
// -------------------------- Delete-Product --------------------------//

func (c *ProductUsecase) DeleteBrand(id int) error {
	err := c.productRepo.DeleteBrand(id)
	return err
}

// -------------------------- List-All-Product --------------------------//

func (c *ProductUsecase) ListAllBrand() ([]response.Brands, error) {
	brands, err := c.productRepo.ListAllBrand()
	return brands, err
}

// -------------------------- List-Single-Product --------------------------//

func (c *ProductUsecase) ListBrand(id int) (response.Brands, error) {
	product, err := c.productRepo.ListBrand(id)
	return product, err
}
// -------------------------- Add-Model --------------------------//

func (c *ProductUsecase) AddModel(model helper.Model) (response.Model, error) {
	newModel, err := c.productRepo.AddModel(model)
	return newModel, err
}

// -------------------------- Update-Model --------------------------//

func (c *ProductUsecase) UpdateModel(id int, model helper.Model) (response.Model, error) {
	updatedItem, err := c.productRepo.UpdateModel(id, model)
	return updatedItem, err
}
// -------------------------- Delete-Model --------------------------//

func (c *ProductUsecase) DeleteModel(id int) error {
	err := c.productRepo.DeleteModel(id)
	return err
}

// -------------------------- List-All-Model --------------------------//

func (c *ProductUsecase) ListAllModel() ([]response.Model, error) {
	model, err := c.productRepo.ListAllModel()
	return model, err
}

// -------------------------- List-Single-Model --------------------------//

func (c *ProductUsecase) ListModel(id int) (response.Model, error) {
	productItem, err := c.productRepo.ListModel(id)
	return productItem, err
}
