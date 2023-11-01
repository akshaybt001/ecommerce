package repository

import (
	"gorm.io/gorm"
	interfaces "main.go/pkg/repository/interface"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRespository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}
