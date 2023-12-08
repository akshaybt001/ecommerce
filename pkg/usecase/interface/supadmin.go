package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
)

type SupAdminUseCase interface {
	SupAdminLogin(supadmin helper.LoginReq) (string, error)
	CreateAdmin(admin helper.CreateAdmin) (response.AdminData, error)
	ListAllAdmins() ([]response.AdminData, error)
	DisplayAdmin(id int) (response.AdminData, error)
	BlockUser(body helper.BlockData) error
	UnblockUser(id int) error
}
