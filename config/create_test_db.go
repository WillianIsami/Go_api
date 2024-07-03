package config

import (
	"fmt"
	"log"
	"os"

	"github.com/WillianIsami/go_api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func ConnectTestDatabase() {
	logger := GetLogger("postgres test db")
	err := godotenv.Load("../.env")
	if err != nil {
		logger.Errorf("Error loading .env: %v", err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_SSL_MODE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}
	db.AutoMigrate(&models.Order{}, &models.OrderItem{})
	TestDB = db
}
