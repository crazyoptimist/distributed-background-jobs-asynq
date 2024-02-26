# Build the application
all: build

build:
	@echo "Building..."

	@go build -o build/server cmd/api/main.go
	@go build -o build/worker cmd/worker/main.go

# Run the API
run_api:
	@go run cmd/api/main.go

# Run the worker
run_worker:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f build/server
	@rm -f build/worker

docker:
	@docker compose -f deployments/compose.yaml build

up:
	@docker compose -f deployments/compose.yaml up

upd:
	@docker compose -f deployments/compose.yaml up -d

down:
	@docker compose -f deployments/compose.yaml down

log:
	@docker compose -f deployments/compose.yaml logs -f

.PHONY: all build run test clean
