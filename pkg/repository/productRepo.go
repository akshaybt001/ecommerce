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

// -------------------------- Create-Brand --------------------------//

func (c *ProductDatabase) AddBrand(Brand helper.Brands) (response.Brands, error) {
	var newBrand response.Brands
	var category response.Category

	query1 := `SELECT * FROM categories WHERE id=$1`
	err := c.DB.Raw(query1, Brand.CategoryId).Scan(&category).Error
	if err != nil {
		return response.Brands{}, err
	}

	query := `INSERT INTO brands (brand,description,category_id,created_at)
		VALUES ($1,$2,$3,NOW())
		RETURNING id,brand AS name,description,category_id`
	err = c.DB.Raw(query, Brand.Name, Brand.Description, Brand.CategoryId).
		Scan(&newBrand).Error
	if err != nil {
		return newBrand, err
	}
	newBrand.CategoryName = category.Category
	return newBrand, err
}

// -------------------------- Update-Product --------------------------//

func (c *ProductDatabase) UpdateBrand(id int, Brand helper.Brands) (response.Brands, error) {
	var updatedBrand response.Brands
	query2 := `UPDATE brands SET brand=$1,description=$2,category_id=$3,updated_at=NOW() WHERE id=$4
		RETURNING id,brand,description,category_id`
	err := c.DB.Raw(query2, Brand.Name, Brand.Description, Brand.CategoryId, id).
		Scan(&updatedBrand).Error
	if err != nil {
		return response.Brands{}, err
	}
	if updatedBrand.Id == 0 {
		return response.Brands{}, fmt.Errorf("there is no such product")
	}
	return updatedBrand, nil
}

// -------------------------- Delete-Product --------------------------//

func (c *ProductDatabase) DeleteBrand(id int) error {
	var exists bool
	isExists := `SELECT EXISTS (SELECT 1 FROM brands WHERE id=$1)`
	c.DB.Raw(isExists, id).Scan(&exists)
	if !exists {
		return fmt.Errorf("there is no such product to delete")
	}
	query := `DELETE FROM brands WHERE id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}

// -------------------------- List-All-Product --------------------------//

func (c *ProductDatabase) ListAllBrand() ([]response.Brands, error) {
	var brands []response.Brands
	getbrandsDetails := `SELECT p.id,p.brand AS name,
		p.description,
		c.category AS category_name, c.id AS category_id
		 FROM brands p JOIN categories c ON p.category_id=c.id `
	err := c.DB.Raw(getbrandsDetails).Scan(&brands).Error
	if err != nil {
		return []response.Brands{}, err
	}
	return brands, nil
}

// -------------------------- List-Single-Product --------------------------//

func (c *ProductDatabase) ListBrand(id int) (response.Brands, error) {
	var brand response.Brands
	query := `SELECT p.id,p.brand AS name,p.description,c.category AS category_name, c.id AS category_id FROM brands p 
		JOIN categories c ON p.category_id=c.id WHERE p.id=$1`
	err := c.DB.Raw(query, id).Scan(&brand).Error
	if err != nil {
		return response.Brands{}, err
	}
	if brand.Id == 0 {
		return response.Brands{}, fmt.Errorf("there is no such brand")
	}
	return brand, err
}

// -------------------------- Add-Model --------------------------//

func (c *ProductDatabase) AddModel(model helper.Model) (response.Model, error) {
	var newModel response.Model
	query := `INSERT INTO models (brand_id,
		model_name,
		sku,
		qty_in_stock,
		color,
		ram,
		battery,
		screen_size,
		storage,
		camera,
		price,
		image,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,NOW())
		RETURNING 
		id,
		model_name,
		brand_id,
		sku,
		qty_in_stock,
		color,
		ram,
		battery,
		screen_size,
		storage,
		camera,
		price,
		image`
	err := c.DB.Raw(query, model.Brand_id,
		model.Model_name,
		model.Sku,
		model.Qty,
		model.Color,
		model.Ram,
		model.Battery,
		model.Screen_size,
		model.Storage,
		model.Camera,
		model.Price,
		model.Image).Scan(&newModel).Error
	return newModel, err
}
