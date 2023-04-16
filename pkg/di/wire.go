//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/api"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/api/service"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/config"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/db"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/client"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/repository"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewCartRepository,
		usecase.NewCartUseCase,
		client.InitProductServiceClient,
		service.NewCartService,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
