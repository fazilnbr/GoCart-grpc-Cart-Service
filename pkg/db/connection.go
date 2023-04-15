package db

import (
	"fmt"

	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/config"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBCart, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(
		domain.Cart{},
		domain.CartItem{},
	)

	return db, dbErr
}
