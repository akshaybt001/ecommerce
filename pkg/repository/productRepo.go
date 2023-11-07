package repository

import (
	"fmt"

	"gorm.io/gorm"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	interfaces "main.go/pkg/repository/interface"
)

type ProductDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &ProductDatabase{DB}
}

// -------------------------- Create-Category --------------------------//

func (c *ProductDatabase) CreateCategory(category helper.Category) (response.Category, error) {
	var newCategory response.Category
	query := `INSERT INTO categories (category,created_at)VAlues($1,NOW())RETURNING id,category`
	err := c.DB.Raw(query, category.Name).Scan(&newCategory).Error
	return newCategory, err
}

// -------------------------- Update-Category --------------------------//

func (c *ProductDatabase) UpdateCategory(category helper.Category, id int) (response.Category, error) {
	var updatedCategory response.Category
	query := `UPDATE categories
	          SET category = $1, updated_at = NOW()
	          WHERE id = $2
	          RETURNING id, category`
	err := c.DB.Raw(query, category.Name, id).Scan(&updatedCategory).Error
	if err != nil {
		return response.Category{}, err
	}
	if updatedCategory.Id == 0 {
		return response.Category{}, fmt.Errorf("no such category to update")
	}
	return updatedCategory, nil
}

// -------------------------- Delete-Category --------------------------//

func (c *ProductDatabase) DeleteCategory(id int) error {
	var exits bool

	query1 := `select exists(select 1 from categories where id=?)`
	c.DB.Raw(query1, id).Scan(&exits)
	if !exits {
		return fmt.Errorf("no category found")
	}
	query := `DELETE FROM categories WHERE id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}

// -------------------------- List-All-Category --------------------------//

func (c *ProductDatabase) ListAllCategories() ([]response.Category, error) {
	var categories []response.Category
	query := `SELECT * FROM categories`
	err := c.DB.Raw(query).Scan(&categories).Error
	return categories, err
}

// -------------------------- List-Single-Category --------------------------//

func (c *ProductDatabase) ListCategory(id int) (response.Category, error) {
	var category response.Category
	var exits bool

	query1 := `select exists(select 1 from categories where id=?)`
	c.DB.Raw(query1, id).Scan(&exits)
	if !exits {
		return response.Category{}, fmt.Errorf("no category found")
	}
	query := `SELECT * FROM categories WHERE id=$1`
	err := c.DB.Raw(query, id).Scan(&category).Error
	if err != nil {
		return response.Category{}, err
	}
	if category.Id == 0 {
		return response.Category{}, fmt.Errorf("no such category")
	}
	return category, nil
}
