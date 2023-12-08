package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
)

type SupAdminRepository interface {
	SupAdminLogin(email string) (domain.SupAdmins, error)
	CreateAdmin(admin helper.CreateAdmin) (response.AdminData, error)
	ListAllAdmins() ([]response.AdminData, error)
	DisplayAdmin(id int) (response.AdminData, error)
	BlockUser(body helper.BlockData) error
	UnblockUser(id int) error
}