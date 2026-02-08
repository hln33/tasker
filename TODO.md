# Tasker - TODO

## Phase 0: Foundation (Infrastructure & Auth)

> **Note**: These foundational items should be implemented alongside or before Phase 2 features for a production-ready application.

### Authentication & User Accounts
- [ ] Add `users` table to database
  - Fields: id, email, password_hash, name, created_at, updated_at
- [ ] Implement password hashing (bcrypt/argon2)
- [ ] Create authentication endpoints
  - `POST /api/auth/register` - User registration
  - `POST /api/auth/login` - User login (returns JWT/session)
  - `POST /api/auth/logout` - User logout
  - `GET /api/auth/me` - Get current user
- [ ] Add JWT middleware for protected routes
- [ ] Create login/register UI pages
- [ ] Add session management (refresh tokens, expiration)
- [ ] Update all existing tables to include `user_id` foreign key
- [ ] Add data migration script for existing data
- [ ] Implement password reset flow (email)
- [ ] Add email verification (optional)

### Database Migrations
- [ ] Choose migration tool (golang-migrate, goose, or custom)
- [ ] Create initial migration schema from current database
- [ ] Set up migration runner in Go application startup
- [ ] Create migration for users table
- [ ] Create migration for adding user_id to existing tables
- [ ] Document migration workflow for developers

### Testing
- [ ] Write integration tests for API endpoints
- [ ] Set up E2E testing (Playwright or Cypress)
- [ ] Configure CI to run tests automatically

### Security Hardening
- [ ] Add CSRF protection for state-changing requests
- [ ] Implement rate limiting middleware
  - Per-IP rate limits
  - Per-user rate limits
- [ ] Add input validation/sanitization
  - Validate all request bodies
  - Sanitize user-generated content
- [ ] Set secure HTTP headers
  - Helmet.js equivalent for Go (secure headers middleware)
  - CSP headers
  - XSS protection
- [ ] Add SQL injection prevention (parameterized queries)
- [ ] Implement CORS policy
- [ ] Add request logging/audit trail
- [ ] Secure cookie configuration (httpOnly, secure, sameSite)

### DevOps & Deployment
- [ ] Environment configuration
  - Support .env files
  - Separate dev/staging/prod configs
  - Document required env vars
- [ ] Structured logging (JSON format)
  - Request logging
  - Error logging
  - Performance metrics
- [ ] Set up CI/CD pipeline
  - Run tests on PR
  - Automated builds
  - Deployment automation
- [ ] Add database backup strategy
- [ ] Monitoring/observability setup
  - Error tracking (Sentry or similar)
  - Performance monitoring (optional)

### Data Management
- [ ] Data export functionality
  - Export all user data as JSON/CSV
- [ ] Data import functionality
  - Bulk import tasks/projects
- [ ] Account deletion with data cleanup (GDPR)
- [ ] Soft delete option (deleted_at column)
- [ ] Archive old completed tasks

### Documentation
- [ ] API documentation (OpenAPI/Swagger)
- [ ] Developer setup guide (README.md)
  - Prerequisites
  - Local development setup
  - Environment variables
- [ ] Deployment guide
  - Docker deployment
  - Production configuration
- [ ] Architecture documentation
  - System overview
  - Database schema diagram
  - Request flow diagrams

---

## Phase 2: Organization (Current Focus)

### 1. Multiple Projects Support
- [ ] Add `projects` table to database
  - Fields: id, name, description, color_hex, created_at, updated_at
- [ ] Create Go backend endpoints
  - `GET /api/projects` - List all projects
  - `POST /api/projects` - Create new project
  - `PUT /api/projects/:id` - Update project
  - `DELETE /api/projects/:id` - Delete project
- [ ] Add `project_id` foreign key to `tasks` table
- [ ] Create project management UI (Svelte components)
  - Project list view
  - Create/edit project form
  - Project color picker

### 2. Project Switching UI
- [ ] Add project selector/dropdown to header
- [ ] Implement "All Projects" view (current behavior)
- [ ] Filter tasks by selected project
- [ ] Update board header to show current project name
- [ ] Add breadcrumb or indicator for active project

### 3. Enhanced Task Metadata
- [ ] Add `due_date` field to `tasks` table
  - Backend: Update task CRUD to handle due dates
  - Frontend: Add date picker to task create/edit form
  - Display due dates on task cards
- [ ] Add `tags` support
  - Create `tags` table (id, name, color_hex)
  - Create `task_tags` join table (task_id, tag_id)
  - Backend endpoints for tag management
  - Tag selector in task form
  - Display tags on task cards

### 4. Filter & Search
- [ ] Add search bar to board view
  - Search by task title and description
- [ ] Add filter controls
  - Filter by project
  - Filter by priority
  - Filter by due date (overdue, today, upcoming, none)
  - Filter by tags
- [ ] Persist filter state in URL query params

### 5. UI/UX Improvements
- [ ] Add empty states for "No tasks in project"
- [ ] Add project color theming (subtle accent colors)
- [ ] Show task count per column per project
- [ ] Improve mobile responsiveness for project selector

---

## Phase 3: Workflow (Future)

### Subtasks
- [ ] Design subtask data model
- [ ] Add subtask CRUD endpoints
- [ ] Create subtask UI within task cards
- [ ] Implement progress indicator based on subtasks

### Task Dependencies
- [ ] Add `blocked_by` relationships to tasks
- [ ] Visual indicators for blocked tasks
- [ ] Prevent moving blocked tasks to "Done"
- [ ] Show dependency chains

### Calendar View
- [ ] Create calendar view route
- [ ] Display tasks grouped by due date
- [ ] Month/week navigation
- [ ] Link calendar events to task edit panel

---

## Phase 4: Polish (Future)

### Goals/OKRs
- [ ] Design goals tracking system
- [ ] Link tasks to goals
- [ ] Progress tracking toward goals

### Analytics Dashboard
- [ ] Task completion rate over time
- [ ] Tasks completed per project
- [ ] Average time in each status
- [ ] Productivity insights

### Power User Features
- [ ] Keyboard shortcuts (n for new, / for search, etc.)
- [ ] Bulk task operations
- [ ] Task templates
- [ ] Quick add via global hotkey
