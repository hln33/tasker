# Tasker

A simple REST API built with Go and Gin, designed to be expanded into a Jira-like task management application.

## Requirements

- Go 1.18 or higher

## Installation

```bash
go mod download
```

## Running the Server

Start the server on port 8080:

```bash
go run main.go
```

The server will start at `http://localhost:8080`

## API Endpoints

### Health Check
```bash
curl http://localhost:8080/
```

Response:
```json
{
  "message": "hello world"
}
```

### Get Task
```bash
curl http://localhost:8080/api/task
```

Response:
```json
{
  "id": "TASK-001",
  "title": "Setup project repository",
  "description": "Initialize the repository with basic project structure",
  "status": "In Progress",
  "priority": "High"
}
```

## Running Tests

Run all tests:
```bash
go test -v
```

Run a specific test:
```bash
go test -v -run TestHealthCheckHandler
```

## Project Structure

```
.
├── main.go          # Application entry point and route handlers
├── main_test.go     # Test suite
├── go.mod           # Go module definition
└── README.md        # This file
```
