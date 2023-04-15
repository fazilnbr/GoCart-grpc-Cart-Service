package usecase

import (
	"context"

	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/domain"
	repository "github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/repository/interface"
	interfaces "github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/usecase/interface"
)

type cartUseCase struct {
	cartRepo repository.CartRepository
}

// GetCart implements interfaces.CartUseCase
func (c *cartUseCase) GetCart(ctx context.Context, userId int64) ([]domain.CartItem, error) {
	cartId,err:=c.cartRepo.CheckInCartOfUser(ctx,userId)
	if err!=nil{
		return []domain.CartItem{},err
	}
	cartItems, err := c.cartRepo.GetCart(ctx, cartId)
	return cartItems, err
}

// RemoveProductFromCart implements interfaces.CartUseCase
func (c *cartUseCase) RemoveProductFromCart(ctx context.Context, productId int64) error {
	return c.cartRepo.RemoveProductFromCart(ctx, productId)
}

// AddCartitemForUser implements interfaces.CartUseCase
func (c *cartUseCase) AddCartitemForUser(ctx context.Context, cartItem domain.CartItem) (int64, error) {
	id, err := c.cartRepo.AddCartitemForUser(ctx, cartItem)
	return id, err
}

// CheckorCreatecart implements interfaces.CartUseCase
func (c *cartUseCase) CheckorCreatecart(ctx context.Context, userId int64) (int64, error) {
	id, err := c.cartRepo.CheckInCartOfUser(ctx, userId)
	if err == nil {
		return id, nil
	}
	if err.Error() != "there is no cart for this user" {
		return -1, err
	}
	id, err = c.cartRepo.CreateCartForUser(ctx, userId)

	return id, err

}

func NewCartUseCase(repo repository.CartRepository) interfaces.CartUseCase {
	return &cartUseCase{
		cartRepo: repo,
	}
}
