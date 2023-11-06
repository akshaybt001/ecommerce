package domain

import "time"

type Category struct {
	Id         uint   `gorm:"primaryKey;unique;not null"`
	category   string `gorm:"unique;not null"`
	Created_at time.Time
	Updated_at time.Time
}

type Model struct {
	Id          uint   `gorm:"primaryKey;unique;not null"`
	ModelName   string `gorm:"unique;not null"`
	Description string
	Category_id uint
	Category    Category `gorm:"foreignKey:Category_id"`
	Created_at  time.Time
	Updated_at  time.Time
}

type Product struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	ProductName    string `gorm:"not null"`
	Model_id     uint
	Model        Model  `gorm:"foreignKey:Model_id"`
	Sku          string `gorm:"not null"`
	Qty_in_stock int
	Color        string
	Ram          int
	Battery      int
	Screen_size  float64
	Storage      int
	Camera       int
	Price        int
	Image        string
	Created_at   time.Time
	Updated_at   time.Time
}

type Images struct {
	Id       uint `gorm:"primaryKey;unique;not null"`
	ProductId  uint
	Product    Product `gorm:"foreignKey:ProductId"`
	FileName string
}
