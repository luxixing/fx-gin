# Makefile
.PHONY: all build clean run test lint swagger package help

# Variable definitions
APP_NAME=fix-gin
BUILD_DIR=./build
MAIN_FILE=./cmd/server/main.go
SWAGGER_FILE=./cmd/swagger/main.go

# Default target
all: clean lint test build

# Build application
build:
	@echo "Building application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "Build completed: $(BUILD_DIR)/$(APP_NAME)"

# Clean build files
clean:
	@echo "Cleaning build files..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "Clean completed"

# Run application
run:
	@echo "Running application..."
	@go run $(MAIN_FILE)

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Code linting
lint:
	@echo "Running code linting..."
	@golangci-lint run

# Generate Swagger documentation
swagger:
	@echo "Checking if swag is installed..."
	@if ! command -v swag > /dev/null; then \
		echo "Installing swag..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	@echo "Generating Swagger documentation..."
	@PATH="$(shell go env GOPATH)/bin:$(PATH)" go run $(SWAGGER_FILE)

# Package project (excluding .env and *.db files)
package:
	@echo "Packaging project..."
	PACKAGE_NAME="fx-gin.tar.gz" && \
	echo "Creating $$PACKAGE_NAME..." && \
	tar --exclude='.git' --exclude='.env' --exclude='*.db' --exclude='build' \
		--exclude='*.tar.gz' --exclude='*.log' --exclude='tmp' --exclude='vendor' \
		-czf "$$PACKAGE_NAME" . && \
	echo "Project packaged: $$PACKAGE_NAME"

# Help information
help:
	@echo "Available commands:"
	@echo "  make build          - Build application"
	@echo "  make clean          - Clean build files"
	@echo "  make run            - Run application"
	@echo "  make test           - Run tests"
	@echo "  make lint           - Run code linting"
	@echo "  make swagger        - Generate Swagger documentation"
	@echo "  make package        - Package project"
	@echo "  make all            - Execute clean, lint, test, build"
	@echo "  make help           - Show help information"