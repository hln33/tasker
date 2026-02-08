# Tasker

> **Jira-like organization, without the complexity**

A personal task management application for individuals who want more than a simple todo list. Tasker provides a Kanban-style workflow with drag-and-drop task management, priority levels, and project organization—built for personal productivity, not team collaboration.

![Tasker Board](docs/images/tasker-board.png)


## What Makes Tasker Different

| Todo Apps | Tasker | Jira |
|-----------|--------|------|
| Basic lists | Kanban workflow | Full project management |
| No organization | Projects/areas | Teams & collaboration |
| Static | Drag-and-drop flow | Overwhelming |

**Perfect for:**
- Individuals managing multiple projects or goals
- People who want visibility into progress (not just a todo list)
- Those who find Jira too heavy but todo apps too light

## Features

### Current (Phase 1 - Foundation)
- **Kanban Board**: Tasks flow through TODO → In Progress → Done
- **Drag & Drop**: Smoothly move tasks between columns
- **Priority Levels**: Low, Medium, High with visual indicators
- **Quick Edit**: Slide-over panel for editing tasks
- **Status Icons**: Visual column headers for at-a-glance status

### Coming Soon (Phase 2 - Organization)
- Multiple projects support
- Filter/search by project, priority, tags
- Due dates and task metadata
- Project switching UI

See [VISION.md](VISION.md) for the full product roadmap and [TODO.md](TODO.md) for what's being built next.

## Tech Stack

### Backend
- **[Go](https://go.dev/)** - High-performance language
- **[Gin](https://gin-gonic.com/)** - HTTP web framework
- **[PostgreSQL](https://www.postgresql.org/)** - Relational database
- **[Docker](https://www.docker.com/)** - Containerization

### Frontend
- **[Svelte 5](https://svelte.dev/)** - Reactive UI framework
- **[SvelteKit](https://kit.svelte.dev/)** - Full-stack web framework
- **[TypeScript](https://www.typescriptlang.org/)** - Type safety
- **[Tailwind CSS](https://tailwindcss.com/)** - Utility-first styling
- **[dnd-kit](https://dndkit.com/)** - Drag-and-drop library

## Development

### Prerequisites
- Go 1.24+
- Node.js 20+
- PostgreSQL 16+ (or use Docker)
- Docker Desktop (recommended)

## Quickstart
TODO: implement root level docker compose up to spin up front-end and back-end at same time.

## API Endpoints

### Tasks
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/task` | Get all tasks |
| GET | `/api/task/:id` | Get task by ID |
| POST | `/api/task` | Create new task |
| PUT | `/api/task/:id` | Update task |
| DELETE | `/api/task/:id` | Delete task |

### Health Check
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | API health check |

See [backend/README.md](backend/README.md) for full API documentation and examples.

## What Tasker is NOT
- ❌ Team collaboration tool (like Asana, Monday.com)
- ❌ Complex project management system (like MS Project)
- ❌ Enterprise software with permissions/roles
- ❌ Document/knowledge base (like Notion, Confluence)

---
