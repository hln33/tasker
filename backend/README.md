# Tasker

A simple REST API built with Go and Gin, designed to be expanded into a Jira-like task management application.

## Requirements

- Docker Desktop 4.0 or higher

## Quick Start

1. **Start the application:**
```bash
docker compose up
```

This automatically:
- Pulls and starts PostgreSQL 16
- Builds the Go backend
- Runs database migrations
- Starts the API server on `http://localhost:8080`

2. **Verify it's working:**
```bash
curl http://localhost:8080/
# Returns: {"message":"hello world"}
```

3. **Stop the application:**
```bash
docker compose down
```

## Development Commands

### Running the Application
```bash
# Start services (with rebuild if needed)
docker compose up --build

# Start in background (detached mode)
docker compose up -d

# View logs
docker compose logs -f

# View logs for specific service
docker compose logs -f backend
docker compose logs -f postgres

# Stop services
docker compose down

# Reset everything (including database)
docker compose down -v
```

### Running Tests
```bash
# Run all tests
go test -v

# Run specific test
go test -v -run TestHealthCheckHandler

# Run all tests recursively
go test -v ./...
```

### Database Management
```bash
# Access PostgreSQL directly
docker compose exec postgres psql -U tasker -d tasker_db

# Backup database
docker compose exec postgres pg_dump -U tasker tasker_db > backup.sql

# Restore database
docker compose exec -T postgres psql -U tasker tasker_db < backup.sql
```

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

### Get All Tasks
```bash
curl http://localhost:8080/api/task
```

Response:
```json
[
  {
    "id": "TASK-001",
    "title": "Setup project repository",
    "description": "Initialize the repository with basic project structure",
    "status": "In Progress",
    "priority": "High",
    "created_at": "2026-02-01T18:11:08Z",
    "updated_at": "2026-02-01T18:11:08Z"
  }
]
```

### Create Task
```bash
curl -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New task",
    "description": "Task description",
    "status": "TODO",
    "priority": "Medium"
  }'
```

### Delete Task
```bash
curl -X DELETE http://localhost:8080/api/task/TASK-001
```

## Architecture

This application follows the **Repository Pattern** to separate business logic from data access:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           REQUEST FLOW                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  HTTP Request                                                               │
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

### Docker Image

The application uses a multi-stage Docker build:

- **Build stage**: `golang:1.24-alpine` (~75MB) - Compiles the Go binary
- **Runtime stage**: `alpine:latest` (~7MB) - Runs the compiled binary
- **Final image size**: ~20MB

Benefits:
- Minimal production image size
- No build tools in production (better security)
- Fast container startup times
- Easy deployment to container registries

## Production Deployment

The Docker setup is production-ready and can be deployed to:

- **AWS ECS** / **Google Cloud Run** - Fully managed container services
- **DigitalOcean App Platform** - Simple container hosting
- **Railway** / **Render** - Push-to-deploy platforms
- **Kubernetes** - For larger scale deployments

### Security Notes

- Change `POSTGRES_PASSWORD` in production
- Use environment variables or Docker secrets for sensitive data
- Run backend in release mode: `GIN_MODE=release`
- Enable SSL for database connections in production

## Troubleshooting

### Port Already in Use
```bash
# Check what's using port 8080
lsof -ti:8080

# Kill the process
lsof -ti:8080 | xargs kill -9

# Or change port in docker-compose.yml
ports:
  - "8081:8080"
```

### Database Connection Issues
```bash
# Check container status
docker compose ps

# Check PostgreSQL is healthy
docker compose exec postgres pg_isready -U tasker -d tasker_db

# View backend logs
docker compose logs backend
```

### Reset Everything
```bash
# Stop and remove all containers, networks, and volumes
docker compose down -v

# Rebuild and start fresh
docker compose up --build
```
