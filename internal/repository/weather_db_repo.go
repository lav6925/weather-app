package repository

import (
	"errors"
	"fmt"
	"log"

	"weather-app/internal/database"
	entities "weather-app/internal/entities/weather-app"

	"gorm.io/gorm"
)

// DeleteWeatherWithCity deletes weather data by city
func DeleteWeatherWithCity(city string) (int, error) {
	err := database.DB.Where("city = ?", city).Delete(&entities.Weather{}).Error
	if err != nil {
		log.Printf("\nFailed to delete weather data in DB: %v", err)
		return 0, fmt.Errorf("failed to delete weather data in DB: %w", err)
	}
	return 1, nil
}

// InsertWeather inserts weather data
func InsertWeather(weather *entities.Weather) error {
	err := database.DB.Create(&weather).Error
	if err != nil {
		log.Printf("\nFailed to insert weather data in DB: %v", err)
		return fmt.Errorf("failed to insert weather data in DB: %w", err)
	}

	return nil
}

// QueryWeather queries weather data based on city and updates instance with the latest entry
func QueryWeather(weather *entities.Weather) error {
	err := database.DB.Where("city = ?", weather.City).Order("timestamp DESC").First(&weather).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("entry not found to update for city: %s", weather.City)
			return err
		}
		log.Printf("\nFailed to retrieve weather data from DB: %v", err)
		return err
	}
	return nil
}

// UpdateWeather updates weather data
func UpdateWeather(weather *entities.Weather) error {
	if err := database.DB.Model(&entities.Weather{}).Where("city = ?", weather.City).Updates(weather).Error; err != nil {
		log.Printf("\nFailed to update weather data for city: %s in DB: %v", weather.City, err)
		return fmt.Errorf("failed to update weather data for city %s in DB: %w", weather.City, err)
	}
	return nil
}

// UpdateOnConflict updates weather data on conflict
func UpdateOnConflict(weather *entities.Weather) error {
	err := database.DB.Where("city = ?", weather.City).Assign(weather).FirstOrCreate(&weather).Error
	if err != nil {
		log.Printf("\nFailed to update weather data in DB: %v", err)
		return fmt.Errorf("failed to update weather data in DB: %w", err)
	}
	return nil
}

// FindWeatherByCity finds weather data by city and returns updated instance
func FindWeatherByCity(city string) (*entities.Weather, error) {
	var weather entities.Weather
	err := database.DB.Where("city = ?", city).First(&weather).Error
	if err != nil {
		return nil, err
	}
	return &weather, nil
}
