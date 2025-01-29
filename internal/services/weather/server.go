package weather

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"weather-app/internal/config"
	"weather-app/internal/database"
	"weather-app/rpc/proto"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func StartServer() {
	// Step 1: Load configuration
	loader := config.NewDefaultLoader()
	var appConfig config.Config
	err := loader.Load("default", "dev", &appConfig)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Step 2: Initialize the database using the loaded config
	database.InitDB(appConfig)

	// Step 3: Initialize WeatherServiceServer
	service := &WeatherServiceServerImpl{
		Config: appConfig, // Passing the config to the service
	}

	// Initialize the gRPC server
	grpcServer := grpc.NewServer()

	// Register the service with the gRPC server
	proto.RegisterWeatherServiceServer(grpcServer, service)

	// Set up a listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", appConfig.Server.Port))
	if err != nil {
		log.Fatalf("Failed to start TCP listener: %v", err)
	}

	// Run the gRPC server in a goroutine
	go func() {
		log.Printf("Starting gRPC server on port %d...\n", appConfig.Server.Port)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// Set up HTTP routing using Gorilla Mux
	router := mux.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// HTTP handler to call GetWeather from the gRPC service
	router.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		// Extract city parameter from query string
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "City parameter is required", http.StatusBadRequest)
			return
		}

		// Call the GetWeather gRPC method
		req := &proto.GetWeatherRequest{City: city}
		resp, err := service.GetWeather(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get weather: %v", err), http.StatusInternalServerError)
			return
		}

		// Return the response as JSON (you can format this according to your needs)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"city": "%s", "description": "%s", "temperature": %.2f}`,
			resp.GetCity(), resp.GetDescription(), resp.GetTemperature())
	}).Methods("GET")

	// Start the HTTP server
	log.Println("Starting HTTP server on port 9603...")
	log.Fatal(http.ListenAndServe(":9603", router))
}
