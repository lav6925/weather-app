package repository

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"weather-app/internal/database"
	entities "weather-app/internal/entities/weather-app"

	er "weather-app/internal/utils"

	"gorm.io/gorm"
)

// DeleteWeatherWithCity deletes weather data by city
func DeleteWeatherWithCity(city string) (int, error) {
	err := database.DB.Where("city = ?", city).Delete(&entities.Weather{}).Error
	if err != nil {
		log.Printf("DeleteWeatherWithCity: Failed to delete weather data in DB: %v", err)
		return 0, er.NewError(http.StatusServiceUnavailable, fmt.Sprintf("failed to delete weather data in DB: %s", err.Error()))
	}
	return 1, nil
}

// InsertWeather inserts weather data
func InsertWeather(weather *entities.Weather) error {
	err := database.DB.Create(&weather).Error
	if err != nil {
		log.Printf("InsertWeather: Failed to insert weather data in DB: %v", err)
		return er.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to insert weather data in DB: %s", err.Error()))
	}

	return nil
}

// QueryWeather queries weather data based on city and updates instance with the latest entry
func QueryWeather(weather *entities.Weather) error {
	err := database.DB.Where("city = ?", weather.City).Order("timestamp DESC").First(&weather).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("QueryWeather: entry not found for city: %s, failed with error: %s", weather.City, err.Error())
			return err
		}
		log.Printf("QueryWeather: failed to retrieve weather data from DB: %v", err)
		return er.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to retrieve weather data from DB: %s", weather.City))
	}
	return nil
}

// UpdateWeather updates weather data
func UpdateWeather(weather *entities.Weather) error {
	if err := database.DB.Model(&entities.Weather{}).Where("city = ?", weather.City).Updates(weather).Error; err != nil {
		log.Printf("UpdateWeather: failed to update weather data for city: %s in DB: %v", weather.City, err)
		return er.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to update weather data for city %s", weather.City))
	}
	return nil
}

// UpdateOnConflict updates weather data if already present else creates new row
func UpdateOnConflict(weather *entities.Weather) error {
	fmt.Printf("Updating for weather %v", weather)
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
		return nil, er.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to find weather data for city %s", city))
	}
	return &weather, nil
}
