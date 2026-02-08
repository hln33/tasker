# Tasker - Product Vision

## Vision Statement
**Tasker is a personal task management app for individuals who want Jira-like organization without the complexity.**

## Core Philosophy
- **Individual-first**: Built for personal productivity, not team collaboration
- **Simple over powerful**: Do fewer things really well
- **Workflow-driven**: Smooth transitions between task states (like Jira, but personal)

## Target User
Individuals who:
- Have multiple projects/goals in flight
- Want visibility into progress (not just a todo list)
- Care about organization and workflow
- Find Jira too heavy but todo apps too light

## Key Differentiators
| Todo Apps | Tasker | Jira |
|-----------|--------|------|
| Basic lists | Kanban workflow | Full project management |
| No organization | Projects/areas | Teams & collaboration |
| Static | Drag-and-drop flow | Overwhelming |

## Tech Stack
- **Frontend**: Svelte 5 + TypeScript + Tailwind CSS
- **Backend**: Go + Gin framework
- **Database**: PostgreSQL
- **Infrastructure**: Docker + Docker Compose

## Feature Roadmap

### Phase 1: Foundation ✓ (Current)
- Kanban board (TODO → In Progress → Done)
- Priority levels (Low, Medium, High)
- CRUD operations for tasks
- Drag-and-drop between columns
- Slide-over edit panel
- Status icons for columns

### Phase 2: Organization (Next)
- Multiple projects support
- Filter/search by project
- Enhanced task metadata (due dates, tags)
- Project switching UI

### Phase 3: Workflow
- Subtasks
- Task dependencies (blocking relationships)
- Calendar view (display tasks by due date)

### Phase 4: Polish (Future)
- Goals/OKRs tracking
- Analytics dashboard (what did I accomplish?)
- Keyboard shortcuts
- Power user features

## What Tasker is NOT
- Team collaboration tool
- Complex project management system
- Enterprise software with permissions/roles
- Document/knowledge base (like Notion)

## Success Metrics
What makes Tasker successful:
- Users can quickly organize tasks across multiple projects
- Workflow (dragging tasks between states) feels smooth and satisfying
- Complex enough for real work, simple enough for daily use
- Calendar and dependency features help with planning ahead
