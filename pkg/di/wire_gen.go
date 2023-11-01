// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"main.go/pkg/api"
	"main.go/pkg/api/handler"
	"main.go/pkg/config"
	"main.go/pkg/db"
	"main.go/pkg/repository"
	"main.go/pkg/usercase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabse(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRespository(gormDB)
	userUseCase := usercase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}