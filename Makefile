.PHONY: all build prepare-artifacts display-artifacts clean run test

# Default target
all: build prepare-artifacts display-artifacts

# Build the Go project
build:
	@echo "Building Go web server..."
	cd hello-web-server && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-web-server -v .

# Prepare artifacts for publishing
prepare-artifacts:
	@echo "Preparing artifacts..."
	mkdir -p publish_output
	cp hello-web-server/hello-web-server publish_output/

# Display artifact contents
display-artifacts:
	@echo "Artifact contents:"
	ls -lh publish_output/

# Run the web server locally
run:
	@echo "Running web server..."
	cd hello-web-server && go run main.go

# Test the web server endpoints
test:
	@echo "Testing endpoints (server must be running)..."
	@echo "Testing root endpoint:"
	curl -s http://localhost:8080/ | jq . || echo "Server not running or jq not installed"
	@echo "\nTesting hello endpoint:"
	curl -s http://localhost:8080/api/hello | jq . || echo "Server not running or jq not installed"
	@echo "\nTesting hello with name:"
	curl -s http://localhost:8080/api/hello?name=Developer | jq . || echo "Server not running or jq not installed"
	@echo "\nTesting health endpoint:"
	curl -s http://localhost:8080/health | jq . || echo "Server not running or jq not installed"

# Clean up build artifacts and publish_output
clean:
	@echo "Cleaning up..."
	rm -rf publish_output
	rm -f hello-web-server/hello-web-server
