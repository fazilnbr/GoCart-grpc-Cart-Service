package service

import (
	usecase "github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/usecase/interface"
)

type CartService struct {
	cartUseCase usecase.CartUseCase
}

func NewCartService(usecase usecase.CartUseCase) *CartService {
	return &CartService{
		cartUseCase: usecase,
	}
}
