package db

import (
	"fmt"

	"github.com/jxeldotdev/url-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect(cfg *config.Config) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", cfg.Database.Host.Value, cfg.Database.User.Value, cfg.Database.Password.Value, cfg.Database.Name.Value, cfg.Database.Port.Value)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
