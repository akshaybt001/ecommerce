package interfaces

import (
	"main.go/pkg/common/helper"
)

type SupAdminUseCase interface {
	SupAdminLogin(supadmin helper.LoginReq) (string, error)
	BlockUser(body helper.BlockData) error
	UnblockUser(id int) error
}
