package usercase

import (
	services "main.go/pkg/usercase/interface"

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
