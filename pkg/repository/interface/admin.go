package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
)

type AdminRespository interface {
	AdminLogin(email string) (domain.Admins, error)
	ShowUser(userID int) (response.UserDetails, error)
	ShowAllUser() ([]response.UserDetails, error)
	BlockUser(body helper.BlockData, adminId int) error
	UnblockUser(id int) error

}
