# The shell to run the makefile with must be defined to work properly in Linux systems
SHELL := /bin/bash

# all the recipes are phony (no files to check).
.PHONY: build dev test fmt clean compose all-docker all-local help

.DEFAULT_GOAL := help

# Output descriptions of all commands
help:
	@echo "Please use 'make <target>', where <target> is one of"
	@echo ""
	@echo "  help                             outputs this helper"
	@echo "  build						      builds the project"
	@echo "  dev 							  runs the project in development mode"
	@echo "  test							  runs the tests"
	@echo "  fmt						  	  runs the golang fmt"
	@echo "  clean						  	  cleans the project"
	@echo "  compose						  runs the docker-compose"
	@echo "  all-docker						  runs all the targets for docker"
	@echo "  all-local						  runs all the targets for local"
	@echo ""
	@echo ""
	@echo "Check the Makefile to know exactly what each target is doing."

# Build the project
build:
	@echo "Building the project..."
	@go build -o bin/main cmd/main.go

# Run the project in development mode
dev:
	@echo "Running the project in development mode..."
	@go run cmd/main.go

# Run the tests
test:
	@echo "Running the tests..."
	@go test -v ./...

# Run the golang fmt
fmt:
	@echo "Running the golang fmt..."
	@go mod tidy
	@go fmt ./...

# Clean the project
clean:
	@echo "Cleaning the project..."
	@rm -rf bin

# Run the docker-compose
compose:
	@echo "Running the docker-compose..."
	@docker-compose -f compose.yml up -d --build

# Run all-docker targets
all-docker: clean fmt build compose

# Run all-local targets
all-local: clean fmt build dev

