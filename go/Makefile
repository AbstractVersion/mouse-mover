# Go Mouse Movement Makefile

# Variables
BINARY_NAME=mouse-mover
DOCKER_IMAGE=go-mouse-mover
GO_FILES=$(shell find . -name '*.go' -type f)
MAIN_FILE=main.go

# Default target
.DEFAULT_GOAL := help

# Help target
.PHONY: help
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Go development targets
.PHONY: deps
deps: ## Download and install dependencies
	@echo "Installing dependencies..."
	go mod download
	go mod tidy
	@echo "Dependencies installed successfully!"

.PHONY: build
build: deps ## Build the application
	@echo "Building $(BINARY_NAME)..."
	CGO_ENABLED=1 go build -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "Build completed: $(BINARY_NAME)"

.PHONY: run
run: build ## Build and run the application continuously
	@echo "Running $(BINARY_NAME) in continuous mode..."
	./$(BINARY_NAME)

.PHONY: run-once
run-once: build ## Build and run the application once
	@echo "Running $(BINARY_NAME) once..."
	./$(BINARY_NAME) once

.PHONY: clean
clean: ## Remove built binaries and temporary files
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME).exe
	rm -f $(BINARY_NAME)-*
	go clean -cache
	@echo "Cleanup completed!"

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

.PHONY: fmt
fmt: ## Format Go code
	@echo "Formatting code..."
	go fmt ./...
	@echo "Code formatted!"

.PHONY: vet
vet: ## Run go vet
	@echo "Running go vet..."
	go vet ./...
	@echo "Vet check completed!"

.PHONY: lint
lint: fmt vet ## Run formatting and vetting

# Cross-compilation targets
.PHONY: build-linux
build-linux: deps ## Build for Linux (amd64)
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(BINARY_NAME)-linux-amd64 $(MAIN_FILE)
	@echo "Linux build completed: $(BINARY_NAME)-linux-amd64"

.PHONY: build-windows
build-windows: deps ## Build for Windows (amd64)
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o $(BINARY_NAME)-windows-amd64.exe $(MAIN_FILE)
	@echo "Windows build completed: $(BINARY_NAME)-windows-amd64.exe"

.PHONY: build-mac
build-mac: deps ## Build for macOS (amd64)
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o $(BINARY_NAME)-darwin-amd64 $(MAIN_FILE)
	@echo "macOS build completed: $(BINARY_NAME)-darwin-amd64"

.PHONY: build-mac-arm
build-mac-arm: deps ## Build for macOS (arm64/M1)
	@echo "Building for macOS ARM64..."
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o $(BINARY_NAME)-darwin-arm64 $(MAIN_FILE)
	@echo "macOS ARM64 build completed: $(BINARY_NAME)-darwin-arm64"

.PHONY: build-all
build-all: build-linux build-windows build-mac build-mac-arm ## Build for all platforms
	@echo "All platform builds completed!"

# Docker targets
.PHONY: docker-build
docker-build: ## Build Docker image
	@echo "Building Docker image: $(DOCKER_IMAGE)..."
	docker build -t $(DOCKER_IMAGE) .
	@echo "Docker image built successfully!"

.PHONY: docker-run
docker-run: docker-build ## Build and run Docker container
	@echo "Running Docker container..."
	docker run -it --rm \
		-e DISPLAY=$${DISPLAY} \
		-v /tmp/.X11-unix:/tmp/.X11-unix:rw \
		-v ~/.Xauthority:/root/.Xauthority:ro \
		$(DOCKER_IMAGE)

.PHONY: docker-run-once
docker-run-once: docker-build ## Build and run Docker container once
	@echo "Running Docker container once..."
	docker run -it --rm \
		-e DISPLAY=$${DISPLAY} \
		-v /tmp/.X11-unix:/tmp/.X11-unix:rw \
		-v ~/.Xauthority:/root/.Xauthority:ro \
		$(DOCKER_IMAGE) ./mouse-mover once

.PHONY: docker-clean
docker-clean: ## Remove Docker image
	@echo "Removing Docker image..."
	docker rmi $(DOCKER_IMAGE) || true
	@echo "Docker cleanup completed!"

# Docker Compose targets
.PHONY: compose-up
compose-up: ## Start with docker-compose
	@echo "Starting with docker-compose..."
	docker-compose up --build

.PHONY: compose-up-detached
compose-up-detached: ## Start with docker-compose in detached mode
	@echo "Starting with docker-compose (detached)..."
	docker-compose up --build -d

.PHONY: compose-down
compose-down: ## Stop docker-compose services
	@echo "Stopping docker-compose services..."
	docker-compose down

.PHONY: compose-logs
compose-logs: ## Show docker-compose logs
	docker-compose logs -f

# Development targets
.PHONY: dev
dev: fmt vet build ## Format, vet, and build for development

.PHONY: watch
watch: ## Watch for file changes and rebuild (requires entr)
	@echo "Watching for changes... (Ctrl+C to stop)"
	@echo "Note: This requires 'entr' to be installed (brew install entr / apt install entr)"
	find . -name '*.go' | entr -r make run-once

# Install targets
.PHONY: install
install: build ## Install the binary to $GOPATH/bin
	@echo "Installing $(BINARY_NAME) to $$GOPATH/bin..."
	cp $(BINARY_NAME) $$GOPATH/bin/
	@echo "Installation completed!"

.PHONY: uninstall
uninstall: ## Remove the binary from $GOPATH/bin
	@echo "Removing $(BINARY_NAME) from $$GOPATH/bin..."
	rm -f $$GOPATH/bin/$(BINARY_NAME)
	@echo "Uninstallation completed!"

# Release targets
.PHONY: release
release: clean build-all ## Create a release with all platform binaries
	@echo "Creating release directory..."
	mkdir -p release
	mv $(BINARY_NAME)-* release/ 2>/dev/null || true
	@echo "Release files created in ./release/"
	@ls -la release/

.PHONY: clean-release
clean-release: ## Clean release directory
	@echo "Cleaning release directory..."
	rm -rf release/
	@echo "Release directory cleaned!"

# System check targets
.PHONY: check-deps
check-deps: ## Check if required system dependencies are installed
	@echo "Checking system dependencies..."
	@command -v go >/dev/null 2>&1 || { echo "Go is not installed!"; exit 1; }
	@command -v docker >/dev/null 2>&1 || { echo "Docker is not installed!"; exit 1; }
	@command -v docker-compose >/dev/null 2>&1 || { echo "Docker Compose is not installed!"; exit 1; }
	@echo "All required dependencies are installed!"

.PHONY: check-x11
check-x11: ## Check X11 setup for GUI automation
	@echo "Checking X11 setup..."
	@echo "DISPLAY: $$DISPLAY"
	@command -v xset >/dev/null 2>&1 && xset q >/dev/null 2>&1 || { echo "X11 not properly configured!"; exit 1; }
	@echo "X11 setup looks good!"

# Quick targets
.PHONY: quick
quick: fmt build run-once ## Quick format, build, and test run

.PHONY: all
all: clean deps fmt vet build test ## Run all checks and build