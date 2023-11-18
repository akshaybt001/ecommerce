package usecase

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	services "main.go/pkg/usecase/interface"

	interfaces "main.go/pkg/repository/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) UserSignUp(user helper.UserReq) (response.UserData, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return response.UserData{}, err
	}
	user.Password = string(hash)
	userData, err := c.userRepo.UserSignUp(user)
	return userData, err
}

//-------------------------- Login --------------------------//

func (c *userUseCase) UserLogin(user helper.LoginReq) (string, error) {
	userData, err := c.userRepo.UserLogin(user.Email)
	if err != nil {
		return "", err
	}

	if user.Email == "" {
		return "", fmt.Errorf("no user found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	if userData.IsBlocked {
		return "", fmt.Errorf("user is blocked")
	}

	claims := jwt.MapClaims{
		"id":  userData.ID,
		"exp": time.Now().Add(time.Hour * 96).Unix(),
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := Token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

//-------------------------- Forgot-Password --------------------------//

func (c *userUseCase) ForgotPassword(forgotPass helper.ForgotPassword) error {
	UserData, err := c.userRepo.ForgotPassword(forgotPass.Email)
	if err != nil {
		return fmt.Errorf("there is no such user")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(forgotPass.NewPassword), 10)
	if err != nil {
		return fmt.Errorf("error in hashing the password")
	}
	forgotPass.NewPassword = string(hash)
	if err = c.userRepo.UpdatePassword(UserData.Id, forgotPass.NewPassword); err != nil {
		return err
	}
	return nil
}
//-------------------------- View-Profile --------------------------//

func (c *userUseCase) ViewProfile(userID int) (response.Userprofile, error) {
	profile, err := c.userRepo.ViewProfile(userID)
	return profile, err
}
//-------------------------- Edit-Profile --------------------------//

func (c *userUseCase) EditProfile(userID int, updatingDetails helper.Userprofile) (response.Userprofile, error) {
	updatedProfile, err := c.userRepo.EditProfile(userID, updatingDetails)
	return updatedProfile, err
}

//-------------------------- Update-Password --------------------------//

func (c *userUseCase) UpdatePassword(userID int, Passwords helper.UpdatePassword) error {

	orginalPassword, err := c.userRepo.FindPassword(userID)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(orginalPassword), []byte(Passwords.OldPassword))
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(Passwords.NewPasswoed), 10)
	if err != nil {
		return err
	}
	newPassword := string(hash)

	err = c.userRepo.UpdatePassword(userID, newPassword)
	return err
}


//-------------------------- Add-Address --------------------------//

func (c *userUseCase) AddAddress(userID int, address helper.Address) error {
	err := c.userRepo.AddAddress(userID, address)
	return err
}

//-------------------------- Update-Address --------------------------//

func (c *userUseCase) UpdateAddress(id, addressId int, address helper.Address) error {
	err := c.userRepo.UpdateAddress(id, addressId, address)
	return err
}