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

func (c *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	cartId := req.CartId
	cartItems, err := c.cartUseCase.GetCart(ctx, cartId)
	if err != nil {
		return &pb.GetCartResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}
	var pbCartItems []*pb.CartItem
	for _, c := range cartItems {
		pbCartItems = append(pbCartItems, &pb.CartItem{
			Id:        c.Id,
			CartId:    c.Cart_id,
			ProductId: c.Product_id,
			Quantity:  c.Quantity,
		})
	}

	return &pb.GetCartResponse{
		Status:   http.StatusOK,
		Cartlist: pbCartItems,
	}, nil

}

func (c *CartService) RemoveProductFromCart(ctx context.Context, req *pb.RemoveProductFromCartRequest) (*pb.RemoveProductFromCartResponse, error) {
	productId := req.ProductId
	err := c.cartUseCase.RemoveProductFromCart(ctx, productId)
	if err != nil {
		return &pb.RemoveProductFromCartResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	return &pb.RemoveProductFromCartResponse{
		Status: http.StatusOK,
	}, nil

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
		Quantity:   req.Quantity,
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
