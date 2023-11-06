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

type adminUseCase struct {
	adminRepo interfaces.AdminRespository
}

func NewAdminUseCase(repo interfaces.AdminRespository) services.AdminUseCase {
	return &adminUseCase{
		adminRepo: repo,
	}
}

// -------------------------- Login --------------------------//

func (c *adminUseCase) AdminLogin(admin helper.LoginReq) (string, error) {
	adminData, err := c.adminRepo.AdminLogin(admin.Email)
	if err != nil {
		return "", err
	}

	if admin.Email == "" {
		return "", fmt.Errorf("admin is not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(adminData.Password), []byte(admin.Password)); err != nil {
		return "", err
	}

	if adminData.IsBlocked {
		return "", fmt.Errorf("admin is blocked")
	}

	claims := jwt.MapClaims{
		"id":  adminData.ID,
		"exp": time.Now().Add(time.Hour * 96).Unix(),
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := Token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
