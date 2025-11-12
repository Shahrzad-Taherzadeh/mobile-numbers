BINARY_NAME=mobile-numbers
CMD_PATH=./cmd/main.go
SWAG_DOCS=./docs

all: run

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(CMD_PATH)
	@echo "Build successful: ./${BINARY_NAME}"

run: build
	@echo "Starting $(BINARY_NAME)..."
	@./$(BINARY_NAME)

deps:
	@echo "Installing dependencies and tidying modules..."
	@go mod tidy
	@echo "Dependencies installed."

swagger: deps
	@echo "Generating Swagger documentation..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@swag init -g $(CMD_PATH)
	@echo "Swagger documentation generated in $(SWAG_DOCS)."

clean:
	@echo "Cleaning up build artifacts..."
	@rm -f $(BINARY_NAME)
	@rm -rf $(SWAG_DOCS)
	@echo "Cleanup complete."

.PHONY: all build run clean swagger deps