package usecase

import (
	repository "github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/repository/interface"
	interfaces "github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/usecase/interface"
)

type cartUseCase struct {
	cartRepo repository.CartRepository
}

func NewCartUseCase(repo repository.CartRepository) interfaces.CartUseCase {
	return &cartUseCase{
		cartRepo: repo,
	}
}
