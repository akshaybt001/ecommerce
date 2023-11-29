package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
)

type UserRepository interface {
	UserSignUp(user helper.UserReq) (response.UserData, error)
	UserLogin(email string) (domain.Users, error)
	ForgotPassword(email string) (response.UserData, error)
	ViewProfile(userID int) (response.Userprofile, error)
	EditProfile(userID int, updatingDetails helper.Userprofile) (response.Userprofile, error)
	FindPassword(id int) (string, error)
	UpdatePassword(id int, newPassword string) error
	AddAddress(id int, address helper.Address) error
	ListAllAddresses(userId int)([]response.Address,error)
	UpdateAddress(userID, addressID int, address helper.Address) error
	CreateWallet(id int) error
}
