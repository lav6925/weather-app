# Project Variables
BINARY_NAME=weather-app
BUILD_DIR=bin/
SRC_DIR=cmd/
PORT=9602

# Database Variables
DB_USER=weather_user
DB_NAME=weather_app
DB_HOST=localhost
DB_PORT=5432

# Default target
all: build

# Build the Go application
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)$(BINARY_NAME) $(SRC_DIR)/main.go

# Run the application
run: build
	$(BUILD_DIR)$(BINARY_NAME)

# Stop any process using the application port
stop:
	@echo "Stopping any process on port $(PORT)..."
	@fuser -k $(PORT)/tcp || true

# Restart the application
restart: stop run

# Clean the build files
clean:
	rm -rf $(BUILD_DIR)

# Refresh protobuf-generated code
proto-refresh:
	protoc --go_out=. --go-grpc_out=. proto/weather.proto

# Format the Go code
fmt:
	go fmt ./...

# Help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build          Build the application"
	@echo "  run            Run the application"
	@echo "  stop           Stop any running process on port $(PORT)"
	@echo "  restart        Restart the application"
	@echo "  clean          Remove built files"
	@echo "  proto-refresh  Regenerate gRPC and protobuf files"
	@echo "  fmt            Format the Go code"

.PHONY: all build run stop restart clean proto-refresh fmt lint dev db-init migrate-up migrate-down help
