package usecase

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	interfaces "main.go/pkg/repository/interface"
	services "main.go/pkg/usecase/interface"
)

type supadminUseCase struct {
	supadminRepo interfaces.SupAdminRepository
}



func NewSupAdminUseCase(repo interfaces.SupAdminRepository) services.SupAdminUseCase {
	return &supadminUseCase{
		supadminRepo: repo,
	}
}

// -------------------------- Login --------------------------//

func (c *supadminUseCase) SupAdminLogin(supadmin helper.LoginReq) (string, error) {
	supadminData, err := c.supadminRepo.SupAdminLogin(supadmin.Email)
	if err != nil {
		return "", err
	}

	if supadmin.Email == "" {
		return "", fmt.Errorf("supadmin is not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(supadminData.Password), []byte(supadmin.Password)); err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":  supadminData.ID,
		"exp": time.Now().Add(time.Hour * 96).Unix(),
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := Token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

// CreateAdmin implements interfaces.SupAdminUseCase.
func (cr *supadminUseCase) CreateAdmin(admin helper.CreateAdmin) (response.AdminData, error) {
	var newAdmin response.AdminData
	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return newAdmin, fmt.Errorf("error hashing password")
	}
	admin.Password = string(hash)
	newAdmin, err = cr.supadminRepo.CreateAdmin(admin)
	return newAdmin, err
}

// ListAllAdmins implements interfaces.SuperAdminUseCase.
func (cr *supadminUseCase) ListAllAdmins() ([]response.AdminDetails, error) {
	admins, err := cr.supadminRepo.ListAllAdmins()
	return admins, err
}

// DisplayAdmin implements interfaces.SuperAdminUseCase.
func (cr *supadminUseCase) DisplayAdmin(id int) (response.AdminDetails, error) {
	admin, err := cr.supadminRepo.DisplayAdmin(id)
	return admin, err
}

// -------------------------- Block-User --------------------------//

func (c *supadminUseCase) BlockUser(body helper.BlockData) error {
	err := c.supadminRepo.BlockUser(body)
	return err
}

// -------------------------- UnBlock-User --------------------------//

func (c *supadminUseCase) UnblockUser(id int) error {
	err := c.supadminRepo.UnblockUser(id)
	return err
}


// -------------------------- Block-Admin --------------------------//

func (c *supadminUseCase) BlockAdmin(body helper.BlockAdminData) error {
	err := c.supadminRepo.BlockAdmin(body)
	return err
}

// -------------------------- UnBlock-Admin --------------------------//

func (c *supadminUseCase) UnblockAdmin(id int) error {
	err := c.supadminRepo.UnblockAdmin(id)
	return err
}