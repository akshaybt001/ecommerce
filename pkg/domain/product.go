package domain

import "time"

type Category struct {
	Id         uint   `gorm:"primaryKey;unique;not null"`
	Category   string `gorm:"unique;not null"`
	Created_at time.Time
	Updated_at time.Time
}

type Brand struct {
	Id          uint   `gorm:"primaryKey;unique;not null"`
	Brand       string `gorm:"unique;not null"`
	Description string
	Category_id uint
	Category    Category `gorm:"foreignKey:Category_id"`
	Created_at  time.Time
	Updated_at  time.Time
}

type Model struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	ModelName    string `gorm:"not null"`
	Brand_id     uint
	Brand        Brand  `gorm:"foreignKey:Brand_id"`
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
	Id        uint `gorm:"primaryKey;unique;not null"`
	ProductId uint
	Brand     Brand `gorm:"foreignKey:ProductId"`
	FileName  string
}
