package repository

import (
	"database/sql"
	"fmt"
	"time"

	types "tasker/internal/types"

	"github.com/jmoiron/sqlx"
)

// BoardRepositoryInterface defines the contract for board data operations
type BoardRepositoryInterface interface {
	GetAllBoards() ([]types.Board, error)
	GetBoardByID(id int) (*types.Board, error)
	CreateBoard(b types.Board) (*types.Board, error)
	UpdateBoard(id int, b types.Board) (*types.Board, error)
	DeleteBoard(id int) error
	GetTasksByBoardID(boardID int) ([]types.Task, error)
}

type BoardRepository struct {
	db *sqlx.DB
}

var Boards BoardRepositoryInterface

func NewBoardRepository(db *sqlx.DB) *BoardRepository {
	return &BoardRepository{db: db}
}

func (r *BoardRepository) GetAllBoards() ([]types.Board, error) {
	var boards []types.Board
	query := `SELECT id, name, description, color, created_at, updated_at FROM boards ORDER BY created_at DESC`

	err := r.db.Select(&boards, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get boards: %w", err)
	}

	return boards, nil
}

func (r *BoardRepository) GetBoardByID(id int) (*types.Board, error) {
	var b types.Board
	query := `SELECT id, name, description, color, created_at, updated_at FROM boards WHERE id = $1`
	err := r.db.Get(&b, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("board not found: %d", id)
		}
		return nil, fmt.Errorf("failed to get board: %w", err)
	}

	return &b, nil
}

func (r *BoardRepository) CreateBoard(b types.Board) (*types.Board, error) {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now

	query := `
		INSERT INTO boards (name, description, color, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, description, color, created_at, updated_at
	`

	var createdBoard types.Board
	err := r.db.QueryRow(
		query,
		b.Name, b.Description, b.Color, b.CreatedAt, b.UpdatedAt,
	).Scan(
		&createdBoard.ID,
		&createdBoard.Name,
		&createdBoard.Description,
		&createdBoard.Color,
		&createdBoard.CreatedAt,
		&createdBoard.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create board: %w", err)
	}

	return &createdBoard, nil
}

func (r *BoardRepository) UpdateBoard(id int, b types.Board) (*types.Board, error) {
	b.UpdatedAt = time.Now()

	query := `
		UPDATE boards
		SET name = COALESCE(NULLIF($1, ''), name),
		    description = COALESCE(NULLIF($2, ''), description),
		    color = COALESCE(NULLIF($3, ''), color),
		    updated_at = $4
		WHERE id = $5
		RETURNING id, name, description, color, created_at, updated_at
	`

	var updatedBoard types.Board
	err := r.db.QueryRow(
		query,
		b.Name, b.Description, b.Color, b.UpdatedAt, id,
	).Scan(
		&updatedBoard.ID,
		&updatedBoard.Name,
		&updatedBoard.Description,
		&updatedBoard.Color,
		&updatedBoard.CreatedAt,
		&updatedBoard.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("board not found: %d", id)
		}
		return nil, fmt.Errorf("failed to update board: %w", err)
	}

	return &updatedBoard, nil
}

func (r *BoardRepository) DeleteBoard(id int) error {
	// First check if board has any tasks
	var taskCount int
	countQuery := `SELECT COUNT(*) FROM tasks WHERE board_id = $1`
	err := r.db.Get(&taskCount, countQuery, id)
	if err != nil {
		return fmt.Errorf("failed to check board tasks: %w", err)
	}

	if taskCount > 0 {
		return fmt.Errorf("cannot delete board with %d tasks. Please reassign or delete tasks first", taskCount)
	}

	query := `DELETE FROM boards WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete board: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("board not found: %d", id)
	}

	return nil
}

func (r *BoardRepository) GetTasksByBoardID(boardID int) ([]types.Task, error) {
	tasks := []types.Task{}
	query := `SELECT id, title, description, status, priority, board_id, created_at, updated_at FROM tasks WHERE board_id = $1 ORDER BY created_at DESC`

	err := r.db.Select(&tasks, query, boardID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks for board: %w", err)
	}

	return tasks, nil
}
