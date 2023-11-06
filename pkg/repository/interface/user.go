package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
)

type UserRepository interface {
	UserSignUp(user helper.UserReq) (response.UserData, error)
	UserLogin(email string) (domain.Users, error)
}
