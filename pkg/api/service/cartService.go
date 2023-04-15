package service

import (
	"context"
	"net/http"

	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/domain"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/pb"
	usecase "github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/usecase/interface"
)

type CartService struct {
	cartUseCase usecase.CartUseCase
}

func (c *CartService) AddProductToCart(ctx context.Context, req *pb.AddProductToCartRequest) (*pb.AddProductToCartResponse, error) {
	userId := req.UserId

	id, err := c.cartUseCase.CheckorCreatecart(ctx, userId)
	if err != nil {
		return &pb.AddProductToCartResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	AddProduct := domain.CartItem{
		Cart_id:    id,
		Product_id: req.ProductId,
		Quantity:   int(req.Quantity),
	}
	id, err = c.cartUseCase.AddCartitemForUser(ctx, AddProduct)
	if err != nil {
		return &pb.AddProductToCartResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	return &pb.AddProductToCartResponse{
		Status: http.StatusOK,
		Id:     id,
	}, nil

}

func NewCartService(usecase usecase.CartUseCase) *CartService {
	return &CartService{
		cartUseCase: usecase,
	}
}
