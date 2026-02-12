# sample-test
hands on repo to test everything

## Hello Web Server

A simple Go web server application demonstrating REST API endpoints, JSON responses, and containerization.

### Features

- **Simple HTTP Web Server**: Built using Go's standard `net/http` package
- **REST API Endpoints**: 
  - `GET /` - Welcome message
  - `GET /api/hello` - Hello greeting (supports `?name=` query parameter)
  - `GET /health` - Health check with uptime information
- **JSON Responses**: All endpoints return JSON formatted responses
- **Request Logging**: Logs all incoming requests with timestamps
- **Docker Support**: Containerized application ready for deployment
- **Make Automation**: Makefile for easy building, testing, and deployment

### Project Structure

```
hello-web-server/
├── main.go          # Main application with HTTP handlers
└── go.mod           # Go module definition

Dockerfile           # Docker image configuration
Makefile            # Build automation
```

### Getting Started

#### Prerequisites

- Go 1.24 or later
- Docker (optional, for containerization)
- curl (for testing)

#### Building the Project

```bash
# Build the project
make build

# Prepare artifacts for deployment
make prepare-artifacts
```

#### Running the Server

```bash
# Run locally
make run

# Or run directly with Go
cd hello-web-server && go run main.go
```

The server will start on `http://localhost:8080`

#### Testing the Endpoints

```bash
# Test all endpoints (server must be running)
make test

# Or test manually:
curl http://localhost:8080/
curl http://localhost:8080/api/hello
curl http://localhost:8080/api/hello?name=Developer
curl http://localhost:8080/health
```

#### Docker Deployment

```bash
# Build the binary
make build

# Prepare artifacts
make prepare-artifacts

# Build Docker image
docker build -t hello-web-server:latest .

# Run the container
docker run -p 8080:8080 hello-web-server:latest
```

#### Clean Up

```bash
make clean
```

### Differences from hello_go

This project is similar in structure to the `hello_go` project but takes a different approach:

| Aspect | hello_go | hello-web-server |
|--------|----------|------------------|
| **Purpose** | Structured logging demonstration | HTTP web server with REST API |
| **Libraries** | dazl (third-party logging) | Standard library (net/http) |
| **Main Feature** | Various log levels and formats | HTTP endpoints with JSON responses |
| **Configuration** | YAML logging config | No external config needed |
| **Output** | Console logs | HTTP responses + request logs |
| **Use Case** | Learning structured logging | Building simple web services |

### Example Responses

**GET /**
```json
{
  "message": "Welcome to Hello Web Server! Visit /api/hello for a greeting.",
  "status": "success",
  "timestamp": "2026-02-12T09:08:00Z"
}
```

**GET /api/hello?name=Developer**
```json
{
  "message": "Hello, Developer!",
  "status": "success",
  "timestamp": "2026-02-12T09:08:00Z"
}
```

**GET /health**
```json
{
  "status": "healthy",
  "uptime": "5m30s",
  "timestamp": "2026-02-12T09:08:00Z"
}
```
