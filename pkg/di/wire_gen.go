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
	"main.go/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabse(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRespository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	adminRespository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRespository)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	productRepository := repository.NewProductRepository(gormDB)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)
	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, productHandler)
	return serverHTTP, nil
}
