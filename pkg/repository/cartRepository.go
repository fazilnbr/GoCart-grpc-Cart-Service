package repository

import (
	"context"
	"errors"

	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/domain"
	interfaces "github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/repository/interface"
	"gorm.io/gorm"
)

type cartDatabase struct {
	DB *gorm.DB
}

// AddCartitemForUser implements interfaces.CartRepository
func (c *cartDatabase) AddCartitemForUser(ctx context.Context, cartItem domain.CartItem) (int64, error) {
	err := c.DB.Create(&cartItem).Error
	return cartItem.Id, err
}

// CreateCartForUser implements interfaces.CartRepository
func (c *cartDatabase) CreateCartForUser(ctx context.Context, userId int64) (int64, error) {
	cart := domain.Cart{User_id: userId}
	err := c.DB.Create(&cart).Error
	return cart.Id, err
}

// CheckInCartOfUser implements interfaces.CartRepository
func (c *cartDatabase) CheckInCartOfUser(ctx context.Context, userId int64) (int64, error) {
	var cart domain.Cart
	err := c.DB.Where("user_id = ?", userId).First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return -1, errors.New("there is no cart for this user")
	}
	return cart.Id, err

}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &cartDatabase{DB}
}
