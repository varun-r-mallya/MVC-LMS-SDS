.DEFAULT_GOAL := help

GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin
BUILD_OUTPUT := ./build/MVC-LMS-SDS
GO_PACKAGES := $(shell go list ./... | grep -v vendor)
BUILD_INPUT := cmd/main.go
UNAME := $(shell uname)
AIR := $(GOPATH_BIN)/air

help:
		@echo "MVC-LMS-SDS Makefile"
		@echo "build   - Build MVC-LMS"
		@echo "dev     - Run development environment"
		@echo "help    - Prints help message"
		@echo "vendor  - Vendor dependencies"
		@echo "init	- Initialize apache"

build:
		@echo "Building..."
		@test -d build || mkdir build
		@$(GO) build -o $(BUILD_OUTPUT) $(BUILD_INPUT)
		@echo "Built as $(BUILD_OUTPUT)"

vendor:
		@echo "Tidy up go.mod..."
		@$(GO) mod tidy
		@echo "Vendoring..."
		@$(GO) mod vendor
		@echo "Done!"

dev:
		@echo "Starting development server..."
		@$(AIR)

init:
		@echo "Initializing apache..."
		@sudo cp ./apache/mvc-lms-sds.conf /etc/apache2/sites-available/
		@sudo a2ensite mvc-lms-sds.conf
		@sudo systemctl restart apache2
		@echo "Done!"