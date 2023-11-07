package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
)

type ProductRepository interface {
	CreateCategory(category helper.Category)(response.Category,error)
	UpdateCategory(category helper.Category, id int) (response.Category, error)
	DeleteCategory(id int) error
	ListAllCategories() ([]response.Category, error)
	ListCategory(id int) (response.Category, error)
}