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
		repository.NewSupAdminRepository,
		repository.NewProductRepository,
		repository.NewCartRepository,
		repository.NewOrderRepository,
		usecase.NewUserUseCase,
		usecase.NewAdminUseCase,
		usecase.NewSupAdminUseCase,
		usecase.NewProductUsecase,
		usecase.NewCartUsecase,
		usecase.NewOrderUseCase,
		handler.NewUserHandler,
		handler.NewAdminHandler,
		handler.NewSupAdminHandler,
		handler.NewProductHandler,
		handler.NewCartHandler,
		handler.NewOrderHandler,
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
