package env

import (
	"os"
)

//TODO: use toml instead of env vars

func GetWeatherAPIURL() string {
	environment := os.Getenv("WEATHER_API_URL")
	if environment == "" {
		environment = "http://api.weatherapi.com/v1/current.json"
	}

	return environment
}

func GetWeatherAPIKey() string {
	key := os.Getenv("WEATHER_API_KEY")
	if key == "" {
		panic("WEATHER_API_KEY is not set")
	}
	return key
}
