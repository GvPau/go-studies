# Variables
BINARY_NAME=fs
BUILD_DIR=bin
MAIN_DIR=cmd

# Build the binary
build:
	@echo "Building the binary..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_DIR)/main.go

# Run the application
run: build
	@echo "Running the application..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	@go test ./... -v

# Clean up binary and other build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(BUILD_DIR)/$(BINARY_NAME)
