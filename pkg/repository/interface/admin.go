package interfaces

import "main.go/pkg/domain"

type AdminRespository interface {
	AdminLogin(email string) (domain.Admins, error)
}