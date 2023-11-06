package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
)

type UserUseCase interface {
	UserSignUp(user helper.UserReq) (response.UserData, error)
	UserLogin(user helper.LoginReq) (string, error)
}
