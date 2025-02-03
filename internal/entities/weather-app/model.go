package entities

import (
	"net/http"
	"strings"
	"time"

	er "weather-app/internal/utils"
	"weather-app/rpc/proto"
)

type WeatherAPIResponse struct {
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

// Weather model (no DB dependencies)
type Weather struct {
	ID          uint   `gorm:"primaryKey"`
	City        string `gorm:"index"`
	Description string
	Temperature float32
	Timestamp   time.Time
}

func NewWeatherFromCity(city string) (Weather, error) {
	city = strings.ToUpper(city)
	if city == "" {
		return Weather{}, er.NewError(http.StatusBadRequest, "city cannot be empty")
	}
	return Weather{
		City: city,
	}, nil
}

func NewWeather(city, description string, temperature float32) (Weather, error) {
	city = strings.ToUpper(city)
	if city == "" {
		return Weather{}, er.NewError(http.StatusBadRequest, "city cannot be empty")
	}
	return Weather{
		City:        city,
		Description: description,
		Temperature: temperature,
		Timestamp:   time.Now(),
	}, nil
}

func NewCreateWeatherResponse(weather *Weather) *proto.CreateWeatherResponse {
	return &proto.CreateWeatherResponse{
		City:        weather.City,
		Description: weather.Description,
		Temperature: weather.Temperature,
		Timestamp:   weather.Timestamp.Format(time.RFC3339),
	}
}

func NewUpdateWeatherResponse(weather *Weather) *proto.UpdateWeatherResponse {
	return &proto.UpdateWeatherResponse{
		City:        weather.City,
		Description: weather.Description,
		Temperature: weather.Temperature,
		Timestamp:   weather.Timestamp.Format(time.RFC3339),
	}
}

func NewDeleteWeatherResponse() *proto.DeleteWeatherResponse {
	return &proto.DeleteWeatherResponse{
		Message: "success",
	}
}

func GetWeatherResponse(weather *Weather) *proto.GetWeatherResponse {
	return &proto.GetWeatherResponse{
		City:        weather.City,
		Description: weather.Description,
		Temperature: weather.Temperature,
		Timestamp:   time.Now().Format(time.RFC3339),
	}
}
