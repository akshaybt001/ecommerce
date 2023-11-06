package repository

import (
	"gorm.io/gorm"
	"main.go/pkg/domain"
	interfaces "main.go/pkg/repository/interface"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRespository {
	return &adminDatabase{DB}
}

//-------------------------- Login --------------------------//

func (c *adminDatabase) AdminLogin(email string) (domain.Admins, error) {
	var adminData domain.Admins
	err := c.DB.Raw("SELECT * FROM admins WHERE email=?", email).Scan(&adminData).Error
	return adminData, err
}
