package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"train-http/config"
	"train-http/internal/models"
)

func DB() *gorm.DB {
	cfg := config.Config{}
	db, err := gorm.Open(postgres.Open(cfg.DB_URL()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	return db
}

func InitDB() {
	db := DB()
	err := db.AutoMigrate(&models.User{}, &models.Card{})
	if err != nil {
		log.Fatal("Failed to auto migrate users table")
	}
}
