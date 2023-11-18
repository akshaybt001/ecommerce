package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/pkg/config"
	"main.go/pkg/domain"
)

func ConnectDatabse(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s", cfg.DBHost, cfg.DBName, cfg.DBUser, cfg.DBPassword, cfg.DBSslmode)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate(
		&domain.Users{},
		&domain.Admins{},
		&domain.SupAdmins{},
		&domain.UserBlockInfo{},
		&domain.Category{},
		&domain.Brand{},
		&domain.Model{},
		&domain.Carts{},
		&domain.CartItem{},
		&domain.OrderItem{},
		&domain.OrderStatus{},
		&domain.PaymentType{},
		&domain.PaymentDetails{},
		&domain.Address{},
		&domain.Images{},
	)
	return db, dbErr
}
