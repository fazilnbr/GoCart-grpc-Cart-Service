package interfaces

import (
	"context"

	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/domain"
)

type CartRepository interface {
	CheckInCartOfUser(ctx context.Context, userId int64) (int64, error)
	CreateCartForUser(ctx context.Context, userId int64) (int64, error)
	AddCartitemForUser(ctx context.Context, cartItem domain.CartItem) (int64, error)
	RemoveProductFromCart(ctx context.Context, productId int64) error
	GetCart(ctx context.Context, cartId int64) ([]domain.CartItem, error)
}
