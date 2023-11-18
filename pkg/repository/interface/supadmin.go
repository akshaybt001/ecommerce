package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/domain"
)

type SupAdminRepository interface {
	SupAdminLogin(email string) (domain.SupAdmins, error)
	BlockUser(body helper.BlockData) error
	UnblockUser(id int) error
}