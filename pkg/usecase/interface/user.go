package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
)

type UserUseCase interface {
	UserSignUp(user helper.UserReq) (response.UserData, error)
	UserLogin(user helper.LoginReq) (string, error)
	ViewProfile(userID int) (response.Userprofile, error)
	ForgotPassword(forgotPass helper.ForgotPassword) error
	EditProfile(userID int, updatingDetails helper.Userprofile) (response.Userprofile, error)
	UpdatePassword(userID int, Passwords helper.UpdatePassword) error
	AddAddress(id int, address helper.Address) error
	UpdateAddress(id, addressId int, address helper.Address) error
	ListAllAddresses(userId int)([]response.Address,error)
	CreateWallet(id int) error
}
