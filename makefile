# Define variables
BINARY_NAME=kogab_simple_kv
BINARY_DIR=bin
SOURCE_DIR=cmd

# Default target
all: run

# Build the go project
build:
	@mkdir -p $(BINARY_DIR)
	@echo "Building the project..."
	@go build -o $(BINARY_DIR)/$(BINARY_NAME) $(SOURCE_DIR)/main.go

# Run the go project
run: build
	@echo "Running the project..."
	@./$(BINARY_DIR)/$(BINARY_NAME)

# Clean the project
clean:
	@echo "Cleaning the project..."
	@rm -rf $(BINARY_DIR)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Run the tests
test:
	@echo "Running the tests..."
	@go test -v ./...

# Format the code
fmt:
	@echo "Formatting the code..."
	@go fmt ./...

# Lint the code
lint:
	@echo "Linting the code..."
	@golangci-lint run
