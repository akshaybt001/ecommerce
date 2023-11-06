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
	)
	return db, dbErr
}
