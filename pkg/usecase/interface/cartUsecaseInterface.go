package interfaces

import (
	"context"

	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/domain"
)

type CartUseCase interface {
	CheckorCreatecart(ctx context.Context, userId int64) (int64, error)
	AddCartitemForUser(ctx context.Context, cartItem domain.CartItem) (int64, error)
	RemoveProductFromCart(ctx context.Context, productId int64) error
	GetCart(ctx context.Context, userId int64) ([]domain.CartItem, error)
}
