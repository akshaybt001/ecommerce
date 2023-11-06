package interfaces

import "main.go/pkg/common/helper"

type AdminUseCase interface {
	AdminLogin(admin helper.LoginReq)(string, error)
}