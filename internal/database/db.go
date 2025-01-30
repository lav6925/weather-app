package database

import (
	"fmt"
	"log"

	model "weather-app/internal/entities/weather-app"

	"weather-app/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config config.Config) error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", config.Database.Host, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.Port, config.Database.SslMode)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	// Auto-migrate the Weather struct
	err = DB.AutoMigrate(&model.Weather{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return err
	}

	log.Println("Database initialized successfully.")
	return nil
}
