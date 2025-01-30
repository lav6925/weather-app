package main

import (
	"log"
	"weather-app/internal/services/weather"
)

func main() {
	if err := weather.StartServer(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
