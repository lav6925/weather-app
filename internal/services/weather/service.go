package weather

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	entities "weather-app/internal/entities/weather-app"
	"weather-app/rpc/proto"

	"weather-app/internal/repository"

	er "weather-app/internal/utils"

	"gorm.io/gorm"
)

func (s *WeatherServiceServerImpl) DeleteWeather(ctx context.Context, req *proto.DeleteWeatherRequest) (*proto.DeleteWeatherResponse, error) {
	city := strings.ToUpper(req.GetCity())

	_, err := repository.DeleteWeatherWithCity(city)
	if err != nil {
		return nil, err
	}

	return entities.NewDeleteWeatherResponse(), nil
}

func (s *WeatherServiceServerImpl) CreateWeather(ctx context.Context, req *proto.CreateWeatherRequest) (*proto.CreateWeatherResponse, error) {
	weather, err := entities.NewWeather(req.GetCity(), req.GetDescription(), req.GetTemperature())
	if err != nil {
		return nil, err
	}

	_, queryDBErr := repository.FindWeatherByCity(weather.City)
	if queryDBErr == nil {
		return nil, queryDBErr
	}

	if dbErr := repository.InsertWeather(&weather); dbErr != nil {
		return nil, dbErr
	}

	return entities.NewCreateWeatherResponse(&weather), nil
}

func (s *WeatherServiceServerImpl) UpdateWeather(ctx context.Context, req *proto.UpdateWeatherRequest) (*proto.UpdateWeatherResponse, error) {
	weather, err := entities.NewWeather(req.GetCity(), req.GetDescription(), req.GetTemperature())
	if err != nil {
		return nil, err
	}

	// Query the record
	_, queryDBErr := repository.FindWeatherByCity(weather.City)
	if queryDBErr != nil {
		return nil, queryDBErr
	}

	// Save the updated record
	updateDBErr := repository.UpdateWeather(&weather)
	if updateDBErr != nil {
		return nil, err
	}

	return entities.NewUpdateWeatherResponse(&weather), nil
}

// GetWeather retrieves the weather data, either from the database or external source based on availability and freshness
func (s *WeatherServiceServerImpl) GetWeather(ctx context.Context, req *proto.GetWeatherRequest) (*proto.GetWeatherResponse, error) {
	city := strings.ToUpper(req.GetCity())
	maxAgeSeconds := s.Config.Weather.RefreshTime // threshold for data freshness

	weather, err := entities.NewWeatherFromCity(city)
	if err != nil {
		return nil, err
	}

	queryDBErr := repository.QueryWeather(&weather)

	if queryDBErr != nil {
		// if error is other than recordNotFound err then return
		if !errors.Is(queryDBErr, gorm.ErrRecordNotFound) {
			return nil, er.NewError(http.StatusNotFound, fmt.Sprintf("entry not found for city: %s", city))
		}
	} else {
		log.Printf("Weather data found: %+v", weather)

		// Calculate elapsed time since last recorded weather data
		elapsedSeconds := time.Since(weather.Timestamp).Seconds()
		if elapsedSeconds <= float64(maxAgeSeconds) {
			// Data is fresh, return it
			return entities.GetWeatherResponse(&weather), nil
		}
	}

	log.Printf("Weather data for %s is outdated or not found, fetching from external source...", city)
	externalWeather, err := repository.FetchExternalWeather(city, s.Config)
	if err != nil {
		return nil, err
	}

	weather = entities.Weather{
		City:        externalWeather.City,
		Description: externalWeather.Description,
		Temperature: externalWeather.Temperature,
		Timestamp:   time.Now(),
	}

	repository.UpdateOnConflict(&weather)

	return externalWeather, nil
}
