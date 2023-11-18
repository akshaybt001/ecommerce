package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
)

type ProductRepository interface {
	CreateCategory(category helper.Category) (response.Category, error)
	UpdateCategory(category helper.Category, id int) (response.Category, error)
	DeleteCategory(id int) error
	ListAllCategories() ([]response.Category, error)
	ListCategory(id int) (response.Category, error)

	AddBrand(product helper.Brands) (response.Brands, error)
	UpdateBrand(id int, Brand helper.Brands) (response.Brands, error)
	DeleteBrand(id int) error
	ListAllBrand() ([]response.Brands, error)
	ListBrand(id int) (response.Brands, error)

	AddModel(model helper.Model) (response.Model, error)
	UpdateModel(id int, productItem helper.Model) (response.Model, error)
	DeleteModel(id int) error
	ListAllModel(queryParams helper.QueryParams) ([]response.Model, error)
	ListModel(id int) (response.Model, error)
	UploadImage(filepath string, productId int) error


}
