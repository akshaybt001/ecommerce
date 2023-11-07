package di

import (
	"github.com/google/wire"
	http "main.go/pkg/api"
	"main.go/pkg/api/handler"
	"main.go/pkg/config"
	"main.go/pkg/db"
	"main.go/pkg/repository"
	"main.go/pkg/usecase"
)

func InitializeAPI1(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabse,
		repository.NewUserRespository,
		repository.NewAdminRepository,
		repository.NewProductRepository,
		usecase.NewUserUseCase,
		usecase.NewAdminUseCase,
		usecase.NewProductUsecase,
		handler.NewUserHandler,
		handler.NewAdminHandler,
		handler.NewProductHandler,
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
