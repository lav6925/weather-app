package env

import (
	"os"
)

func GetWeatherAPIKey() string {
	key := os.Getenv("WEATHER_API_KEY")
	if key == "" {
		panic("WEATHER_API_KEY is not set")
	}
	return key
}
