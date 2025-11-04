# Makefile for Order Service Go Microservice

.PHONY: build run test clean docker-run docker-build prepare

# Prepare the project (generate go.sum)
prepare:
	go mod tidy

# Build the application
build: prepare
	go build -o analytics cmd/main.go

# Run the application
run: prepare
	go run cmd/main.go

# Run tests
test: prepare
	go test ./...

# Clean build files
clean:
	rm -rf bin/

# Build Docker image
docker-build: prepare
	docker build -t analytics .

# Run with Docker
docker-run: docker-build
	docker run -p 8080:8080 analytics

# Run with docker-compose
docker-compose-up: prepare
	docker-compose up --build

# Stop docker-compose
docker-compose-down:
	docker-compose down