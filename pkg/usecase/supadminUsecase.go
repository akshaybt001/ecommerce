package usecase

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"main.go/pkg/common/helper"
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