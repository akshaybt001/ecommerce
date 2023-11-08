package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
)

type ProductUsecase interface {
	CreateCategory(category helper.Category) (response.Category, error)
	UpdateCategory(category helper.Category, id int) (response.Category, error)
	DeleteCategory(id int) error
	ListAllCategories() ([]response.Category, error)
	ListCategory(id int) (response.Category, error)

	AddBrand(Brand helper.Brands) (response.Brands, error)
	UpdateBrand(id int, Brand helper.Brands) (response.Brands, error)
	DeleteBrand(id int) error
	ListAllBrand() ([]response.Brands, error)
	ListBrand(id int) (response.Brands, error)

	AddModel(model helper.Model) (response.Model, error)

}
