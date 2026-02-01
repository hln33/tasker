# Tasker

A simple REST API built with Go and Gin, designed to be expanded into a Jira-like task management application.

## Requirements

- Go 1.18 or higher
- PostgreSQL 12 or higher

## Installation

1. Install dependencies:
```bash
go mod download
```

2. Install PostgreSQL:
```bash
# macOS
brew install postgresql
brew services start postgresql

# Ubuntu
sudo apt-get install postgresql
sudo systemctl start postgresql
```

3. Create database:
```bash
createdb tasker_db
```

4. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your database credentials
```

5. Run the server:
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

## Architecture

### Request Flow: Handler → Repository → Database

This application follows the **Repository Pattern** to separate business logic from data access:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           REQUEST FLOW                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  1. HTTP Request                                                            │
│     │                                                                       │
│     ▼                                                                       │
│  ┌──────────────────┐                                                       │
│  │   HTTP Handler   │  ◄── Parses JSON, validates input, returns HTTP      │
│  └────────┬─────────┘       responses, handles status codes                  │
│           │                                                                 │
│           │ calls repository methods                                         │
│           ▼                                                                 │
│  ┌──────────────────┐                                                       │
│  │    Repository    │  ◄── Executes SQL queries, manages database           │
│  └────────┬─────────┘       connection, handles data logic                   │
│           │                                                                 │
│           │ executes SQL                                                     │
│           ▼                                                                 │
│  ┌──────────────────┐                                                       │
│  │   PostgreSQL DB  │  ◄── Stores and retrieves task data                   │
│  └──────────────────┘                                                       │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Layer Responsibilities

| Layer          | Responsibility                                          | Location                |
|----------------|---------------------------------------------------------|-------------------------|
| **Handler**    | - Parse HTTP requests<br>- Validate JSON input<br>- Return HTTP status codes<br>- No SQL/database logic | `internal/handlers/`    |
| **Repository** | - Execute SQL queries<br>- Map database rows to structs<br>- Handle connection errors<br>- Transaction management | `internal/repository/`  |
| **Database**   | - Persistent data storage<br>- ACID guarantees<br>- Indexing and constraints | PostgreSQL              |
andlers focus on HTTP, repository focuses on data
