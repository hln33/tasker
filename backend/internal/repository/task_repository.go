package repository

import (
	"database/sql"
	"fmt"
	"time"

	types "tasker/internal/types"

	"github.com/jmoiron/sqlx"
)

// TaskRepositoryInterface defines the contract for task data operations
type TaskRepositoryInterface interface {
	GetAllTasks() ([]types.Task, error)
	GetTaskByID(id int) (*types.Task, error)
	CreateTask(t types.Task) (*types.Task, error)
	UpdateTask(id int, t types.Task) (*types.Task, error)
	DeleteTask(id int) error
}

type TaskRepository struct {
	db *sqlx.DB
}

var Tasks TaskRepositoryInterface

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAllTasks() ([]types.Task, error) {
	tasks := []types.Task{}
	query := `SELECT id, title, description, status, priority, board_id, created_at, updated_at FROM tasks ORDER BY created_at DESC`

	err := r.db.Select(&tasks, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}

	return tasks, nil
}

func (r *TaskRepository) GetTaskByID(id int) (*types.Task, error) {
	var t types.Task
	query := `SELECT id, title, description, status, priority, board_id, created_at, updated_at FROM tasks WHERE id = $1`
	err := r.db.Get(&t, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task not found: %d", id)
		}
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return &t, nil
}

func (r *TaskRepository) CreateTask(t types.Task) (*types.Task, error) {
	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now

	query := `
		INSERT INTO tasks (title, description, status, priority, board_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, title, description, status, priority, board_id, created_at, updated_at
	`

	var createdTask types.Task
	err := r.db.QueryRow(
		query,
		t.Title, t.Description, t.Status, t.Priority, t.BoardID, t.CreatedAt, t.UpdatedAt,
	).Scan(
		&createdTask.ID,
		&createdTask.Title,
		&createdTask.Description,
		&createdTask.Status,
		&createdTask.Priority,
		&createdTask.BoardID,
		&createdTask.CreatedAt,
		&createdTask.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return &createdTask, nil
}

func (r *TaskRepository) DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task not found: %d", id)
	}

	return nil
}

func (r *TaskRepository) UpdateTask(id int, t types.Task) (*types.Task, error) {
	t.UpdatedAt = time.Now()

	query := `
		UPDATE tasks
		SET title = COALESCE(NULLIF($1, ''), title),
		    description = COALESCE(NULLIF($2, ''), description),
		    status = COALESCE(NULLIF($3, ''), status),
		    priority = COALESCE(NULLIF($4, ''), priority),
		    board_id = COALESCE(NULLIF($5, 0), board_id),
		    updated_at = $6
		WHERE id = $7
		RETURNING id, title, description, status, priority, board_id, created_at, updated_at
	`

	var updatedTask types.Task
	err := r.db.QueryRow(
		query,
		t.Title, t.Description, t.Status, t.Priority, t.BoardID, t.UpdatedAt, id,
	).Scan(
		&updatedTask.ID,
		&updatedTask.Title,
		&updatedTask.Description,
		&updatedTask.Status,
		&updatedTask.Priority,
		&updatedTask.BoardID,
		&updatedTask.CreatedAt,
		&updatedTask.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task not found: %d", id)
		}
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return &updatedTask, nil
}
