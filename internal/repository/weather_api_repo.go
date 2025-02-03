package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"weather-app/internal/config"
	entities "weather-app/internal/entities/weather-app"
	"weather-app/internal/env"
	errors "weather-app/internal/utils"
	"weather-app/rpc/proto"
)

// FetchExternalWeather fetches weather data from an external API and validates the response.
func FetchExternalWeather(city string, appConfig config.Config) (*proto.GetWeatherResponse, error) {
	apiURL := fmt.Sprintf("%s?key=%s&q=%s", appConfig.Weather.APIURL, env.GetWeatherAPIKey(), url.QueryEscape(city))

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, errors.NewError(http.StatusServiceUnavailable, "failed to fetch external weather data")
	}
	defer resp.Body.Close()

	// Check if the response status is not 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewError(resp.StatusCode, fmt.Sprintf("unexpected API response: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode)))
	}

	// Decode JSON response
	result := &entities.WeatherAPIResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.NewError(http.StatusInternalServerError, "failed to decode response body")
	}

	// Validate the response data
	if result.Location.Name == "" || result.Current.Condition.Text == "" {
		return nil, errors.NewError(http.StatusBadGateway, "invalid API response: missing required weather data")
	}

	return &proto.GetWeatherResponse{
		City:        strings.ToUpper(result.Location.Name),
		Description: result.Current.Condition.Text,
		Temperature: float32(result.Current.TempC),
		Timestamp:   time.Now().Format(time.RFC3339),
	}, nil
}
