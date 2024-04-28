#!/bin/bash

# GAIKeep Makefile - Defines automation tasks for building, testing, and running the GAIKeep application.

# *** Building ***

all: clean build run  # This target depends on completing clean and build steps before running the application.

# Build
# Compiles the Go code and places the app in the bin directory, then makes it executable.
# For more information on the 'go build' command, see the official docs: [https://go.dev/doc/cmd/go/#hdr-Compile_packages_and_dependencies]
# For how the Go compiler works, including building for multiple targets: [https://go.dev/doc/compiler]
build: 
	@echo "Building..." && \
	go build -o ./bin/gaikeep . && \
    chmod +x ./bin/gaikeep && \
    echo "Done building!"

# Test
# Runs unit tests for the GAIKeep application
# For more information on the 'go test' command, see the official docs: [https://go.dev/doc/cmd/go/#hdr-Test_packages]
test:
	@echo "Running tests..." && \
	go test ./... && \
	echo "Done running tests!"

# Clean
# Removes any previously built artifacts
# For more information on the 'rm' command, see the man page: [https://linux.die.net/man/1/rm]
clean:
	@echo "Removing bin folder..." && \
	rm -rf bin/ && \
	echo "bin folder removed!"

# Setup
# Prepares the project environment (assuming Go is installed)
# For more information on the 'go mod init' command, see the official docs: [https://go.dev/doc/cmd/go/#hdr-Create_a_new_module]
# For more information on the 'go get' command, see the official docs: [https://go.dev/doc/cmd/go/#hdr-Download_and_install_packages_and_dependencies]
setup:
	@echo "Setting up the GAIKeep project environment..."
	@which go || (echo "Go is not installed. Please install Go first."; exit 1)
	@if [ ! -f go.mod ]; then \
       	echo "Initializing new Go module..."; \
        go mod init github.com/christimiller/gaikeep; \
    fi
	@echo "Installing dependencies..."
	go get -u ./...
	@echo "Installing Cobra CLI..."
	go get -u github.com/spf13/cobra@latest
	go install github.com/spf13/cobra-cli@latest
	@echo "Setup completed!"

# Run
# Executes the GAIKeep application
run:  
	@./bin/gaikeep search 
	@./bin/gaikeep generate 
	@./bin/gaikeep store

# *** Troubleshooting Resources ***

# Common troubleshooting steps:
# 1. Verify Go installation and version (run 'go version').
# 2. Ensure you're in the correct project directory when running Makefile commands.
# 3. Carefully examine error messages for clues.
# 4. Search online for solutions to common Go build errors. 
# 5. Consider using LLMs like ChatGPT or Gemini for assistance with specific errors or commands. 

.PHONY: all build test clean setup run 

