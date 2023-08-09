.DEFAULT_GOAL := build

APP_NAME := test_app
CMD_DIR := cmd/test_app/main.go

build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) $(CMD_DIR)

clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)

.PHONY: build run clean
