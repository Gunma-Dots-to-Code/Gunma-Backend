package db

import (
	"os"

	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Question{})
	db.AutoMigrate(&model.Answer{})

	return db
}
