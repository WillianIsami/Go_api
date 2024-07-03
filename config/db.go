package config

import (
	"fmt"
	"os"

	"github.com/WillianIsami/go_api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	err := godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Error loading .env: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL_MODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Postgres connection error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.Category{}, &models.Order{}, &models.OrderItem{}, &models.Product{}, &models.User{})
	if err != nil {
		logger.Errorf("Postgres automigrate error: %v", err)
		return nil, err
	}
	return db, nil
}
