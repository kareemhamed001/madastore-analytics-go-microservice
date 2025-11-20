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

PROTOC = protoc
PROTO_DIR = proto
OUT_DIR = common/genproto/analytics
PROTO_FILE = analytics.proto

protoc-gen:
	@echo "🚀 Generating Go and gRPC code from $(PROTO_FILE)..."
	@$(PROTOC) \
		--proto_path=$(PROTO_DIR) $(PROTO_DIR)/$(PROTO_FILE) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative
	@echo "✅ Generation complete! Files saved in $(OUT_DIR)"