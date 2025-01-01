# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	@go build -o caliborg cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f caliborg

.PHONY: all build run test clean
