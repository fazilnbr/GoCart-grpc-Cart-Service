package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/client"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/domain"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/pb"
	usecase "github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/usecase/interface"
)

type CartService struct {
	cartUseCase usecase.CartUseCase
	productSvc  client.ProductServiceClient
}

func (c *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	userId := req.UserId
	cartItems, err := c.cartUseCase.GetCart(ctx, userId)
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

	productId := req.ProductId

	product, err := c.productSvc.GetProduct(productId)
	if err != nil {
		return &pb.AddProductToCartResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}
	if product.Stock < req.Quantity {
		err = errors.New("Out Of Stock")
		return &pb.AddProductToCartResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

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
		Product_id: product.Id,
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

func NewCartService(usecase usecase.CartUseCase, productSvc client.ProductServiceClient) *CartService {
	return &CartService{
		cartUseCase: usecase,
		productSvc:  productSvc,
	}
}
