package weather

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"weather-app/internal/config"
	database "weather-app/internal/database"
	model "weather-app/internal/entities/weather-app"
	"weather-app/internal/env"
	"weather-app/rpc/proto"

	"gorm.io/gorm"
)

// WeatherServiceServerImpl implements the WeatherServiceServer interface
type WeatherServiceServerImpl struct {
	proto.UnimplementedWeatherServiceServer
	Config config.Config
}

// fetchExternalWeather fetches weather data from an external API
func fetchExternalWeather(city string, appConfig config.Config) (*proto.GetWeatherResponse, error) {
	url := fmt.Sprintf("%s?key=%s&q=%s", appConfig.Weather.APIURL, env.GetWeatherAPIKey(), url.QueryEscape(city))
	fmt.Println("url formed is: ", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch external weather data: %w", err)
	}
	defer resp.Body.Close()

	// Use io.TeeReader to log while preserving the original body
	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)

	// Read the body normally (without consuming it completely)
	body, err := io.ReadAll(tee)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Log the response body
	log.Printf("\nResponse Body: \n%s", body)

	// Restore the response body for further use
	resp.Body = io.NopCloser(&buf)

	var result struct {
		Location struct {
			Name      string `json:"name"`
			Localtime string `json:"localtime"`
		} `json:"location"`
		Current struct {
			TempC     float64 `json:"temp_c"`
			Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
		} `json:"current"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &proto.GetWeatherResponse{
		City:        strings.ToUpper(result.Location.Name),
		Description: result.Current.Condition.Text,
		Temperature: float32(result.Current.TempC),
	}, nil
}

// GetWeather retrieves the weather data, either from the database or external source
func (s *WeatherServiceServerImpl) GetWeather(ctx context.Context, req *proto.GetWeatherRequest) (*proto.GetWeatherResponse, error) {
	city := strings.ToUpper(req.GetCity())
	maxAgeSeconds := s.Config.Weather.RefreshTime //threshold for data freshness

	var weather model.Weather
	err := database.DB.Where("city = ?", city).
		Order("timestamp DESC").
		First(&weather).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("No weather data found for", city)
		} else {
			log.Printf("Database error: %v", err)
			return nil, fmt.Errorf("database error: %w", err)
		}
	} else {
		log.Printf("Weather data found: %+v", weather)

		// Calculate elapsed time since last recorded weather data
		elapsedSeconds := time.Since(weather.Timestamp).Seconds()
		if elapsedSeconds <= float64(maxAgeSeconds) {
			// Data is fresh, return it
			return &proto.GetWeatherResponse{
				City:        weather.City,
				Description: weather.Description,
				Temperature: weather.Temperature,
				Timestamp:   time.Now().Format(time.RFC3339),
			}, nil
		}
	}

	// Data is either not in the DB or outdated, fetch it from the external API
	log.Printf("Weather data for %s is outdated or not found, fetching from external source...", city)
	externalWeather, err := fetchExternalWeather(city, s.Config)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch external weather: %w", err)
	}

	// Create a Weather instance for insertion or update
	weather = model.Weather{
		City:        externalWeather.City,
		Description: externalWeather.Description,
		Temperature: externalWeather.Temperature,
		Timestamp:   time.Now(),
	}

	// Use GORM's `Save` method to insert or update based on conflict
	err = database.DB.Where("city = ?", weather.City).Assign(weather).FirstOrCreate(&weather).Error
	if err != nil {
		log.Printf("\nFailed to update weather data in DB: %v", err)
		return nil, fmt.Errorf("failed to update weather data in DB: %w", err)
	}

	// Return the newly fetched data
	return externalWeather, nil
}

// Other CRUD methods: CreateWeather, UpdateWeather, DeleteWeather...
