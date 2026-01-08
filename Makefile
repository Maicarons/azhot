# Makefile for azhot project
#
# This Makefile provides convenient targets for common build operations.
# It works on both local environments and CI systems.

.PHONY: help build clean test run dev package release build-multiarch

# Default target
help:
	@echo "azhot Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build            - Build for current platform"
	@echo "  make run              - Build and run the application"
	@echo "  make dev              - Run in development mode"
	@echo "  make test             - Run tests"
	@echo "  make test-all         - Run all tests with coverage"
	@echo "  make package          - Build and package for all platforms"
	@echo "  make release          - Build and package release versions"
	@echo "  make build-multiarch  - Build for multiple architectures"
	@echo "  make clean            - Clean build artifacts"
	@echo "  make fmt              - Format code"
	@echo "  make tidy             - Tidy go modules"
	@echo "  make docker-build     - Build Docker image"
	@echo "  make docker-run       - Run application in Docker"
	@echo ""

# Build for current platform
build:
	@echo "Building for current platform..."
	@mkdir -p build
	@cd build && cmake .. && cmake --build . --target build

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@cd build && cmake .. && cmake --build . --target azhot_clean || rm -rf build dist package

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run all tests with coverage
test-all:
	@echo "Running all tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...

# Run the application
run: build
	@echo "Running application..."
	@./build/azhot

# Development mode
dev:
	@echo "Starting development mode..."
	@go run main.go

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Tidy go modules
tidy:
	@echo "Tidying go modules..."
	@go mod tidy

# Build and package for all platforms
package:
	@echo "Building and packaging for all platforms..."
	@mkdir -p build
	@cd build && cmake .. && cmake --build . --target package

# Build and package release versions
release: package
	@echo "Release packages created in package/ directory"

# Build for multiple architectures
build-multiarch:
	@echo "Building for multiple architectures..."
	@mkdir -p build
	@cd build && cmake .. && cmake --build . --target build-multiarch

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	@docker build -t azhot .

# Run application in Docker
docker-run: docker-build
	@echo "Running application in Docker..."
	@docker run -d -p 8080:8080 azhot

# Build cross-platform binaries using scripts
build-release-linux:
	@echo "Building release for Linux..."
	@chmod +x scripts/build-release.sh
	@./scripts/build-release.sh

# Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	@swag init

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go install github.com/swaggo/swag/cmd/swag@latest