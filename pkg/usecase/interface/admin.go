package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
)

type AdminUseCase interface {
	AdminLogin(admin helper.LoginReq) (string, error)
	ShowUser(userID int) (response.UserDetails, error)
	ShowAllUser() ([]response.UserDetails, error)
	BlockUser(body helper.BlockData, adminId int) error
	UnblockUser(id int) error
}
