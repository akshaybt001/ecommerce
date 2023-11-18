package domain

type Carts struct{
	Id       uint `gorm:"primaryKey;unique;not null"`
	User_id  uint
	Users    Users `gorm:"foreignKey:User_id"`
	SubTotal int
	Total    int
}

type CartItem struct {
	Id       uint `gorm:"primaryKey;unique;not null"`
	Carts_id uint
	Carts    Carts `gorm:"foreignKey:Carts_id"`
	Model_id uint
	Model    Model `gorm:"foreignKey:Model_id"`
	Quantity int
}
