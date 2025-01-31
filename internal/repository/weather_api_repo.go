package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"weather-app/internal/config"
	"weather-app/internal/env"
	"weather-app/rpc/proto"
)

// FetchExternalWeather fetches weather data from an external API and validates the response.
func FetchExternalWeather(city string, appConfig config.Config) (*proto.GetWeatherResponse, error) {
	apiURL := fmt.Sprintf("%s?key=%s&q=%s", appConfig.Weather.APIURL, env.GetWeatherAPIKey(), url.QueryEscape(city))
	log.Println("Fetching weather data from:", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch external weather data: %w", err)
	}
	defer resp.Body.Close() // The underlying connection (socket) remains open if you don’t close it, which can lead to resource exhaustion (too many open connections).

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected API response status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

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

	if result.Location.Name == "" || result.Current.Condition.Text == "" {
		return nil, fmt.Errorf("invalid API response: missing required weather data")
	}

	return &proto.GetWeatherResponse{
		City:        strings.ToUpper(result.Location.Name),
		Description: result.Current.Condition.Text,
		Temperature: float32(result.Current.TempC),
		Timestamp:   time.Now().Format(time.RFC3339),
	}, nil
}
