package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"

	"tasker/internal/repository"
	types "tasker/internal/types"

	"github.com/gin-gonic/gin"
)

// MockTaskRepository is an in-memory implementation for testing
type MockTaskRepository struct {
	tasks map[string]types.Task
	mu    sync.RWMutex
}

func NewMockTaskRepository() *MockTaskRepository {
	return &MockTaskRepository{
		tasks: make(map[string]types.Task),
	}
}

func (m *MockTaskRepository) GetAllTasks() ([]types.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]types.Task, 0, len(m.tasks))
	for _, t := range m.tasks {
		result = append(result, t)
	}
	return result, nil
}

func (m *MockTaskRepository) GetTaskByID(id string) (*types.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if t, ok := m.tasks[id]; ok {
		return &t, nil
	}
	return nil, errors.New("task not found: " + id)
}

func (m *MockTaskRepository) CreateTask(t types.Task) (*types.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now

	m.tasks[t.ID] = t
	return &t, nil
}

func (m *MockTaskRepository) DeleteTask(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.tasks[id]; !ok {
		return errors.New("task not found: " + id)
	}
	delete(m.tasks, id)
	return nil
}

func (m *MockTaskRepository) UpdateTask(id string, t types.Task) (*types.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	existing, exists := m.tasks[id]
	if !exists {
		return nil, errors.New("task not found: " + id)
	}

	// Update only non-empty fields (except Status/Priority which can be empty)
	if t.Title != "" {
		existing.Title = t.Title
	}
	if t.Description != "" {
		existing.Description = t.Description
	}
	if t.Status != "" {
		existing.Status = t.Status
	}
	if t.Priority != "" {
		existing.Priority = t.Priority
	}
	if t.BoardID != 0 {
		existing.BoardID = t.BoardID
	}

	// Update timestamp
	existing.UpdatedAt = time.Now()

	m.tasks[id] = existing
	return &existing, nil
}

func (m *MockTaskRepository) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tasks = make(map[string]types.Task)
}

// MockBoardRepository is an in-memory implementation for testing
type MockBoardRepository struct {
	boards map[int]types.Board
	mu     sync.RWMutex
	nextID int
}

func NewMockBoardRepository() *MockBoardRepository {
	return &MockBoardRepository{
		boards: make(map[int]types.Board),
		nextID: 1,
	}
}

func (m *MockBoardRepository) GetAllBoards() ([]types.Board, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]types.Board, 0, len(m.boards))
	for _, b := range m.boards {
		result = append(result, b)
	}
	return result, nil
}

func (m *MockBoardRepository) GetBoardByID(id int) (*types.Board, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if b, ok := m.boards[id]; ok {
		return &b, nil
	}
	return nil, fmt.Errorf("board not found: %d", id)
}

func (m *MockBoardRepository) CreateBoard(b types.Board) (*types.Board, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	b.ID = m.nextID
	m.nextID++

	m.boards[b.ID] = b
	return &b, nil
}

func (m *MockBoardRepository) UpdateBoard(id int, b types.Board) (*types.Board, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	existing, exists := m.boards[id]
	if !exists {
		return nil, fmt.Errorf("board not found: %d", id)
	}

	// Update only non-empty fields
	if b.Name != "" {
		existing.Name = b.Name
	}
	if b.Description != "" {
		existing.Description = b.Description
	}
	if b.Color != "" {
		existing.Color = b.Color
	}

	// Update timestamp
	existing.UpdatedAt = time.Now()

	m.boards[id] = existing
	return &existing, nil
}

func (m *MockBoardRepository) DeleteBoard(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.boards[id]; !ok {
		return fmt.Errorf("board not found: %d", id)
	}
	delete(m.boards, id)
	return nil
}

func (m *MockBoardRepository) GetTasksByBoardID(boardID int) ([]types.Task, error) {
	// This is a simplified implementation for testing
	// In real implementation, this would query the task repository
	return []types.Task{}, nil
}

var mockRepo *MockTaskRepository
var mockBoardRepo *MockBoardRepository

func setupTest() {
	gin.SetMode(gin.TestMode)
	nextID = 1
	mockRepo = NewMockTaskRepository()
	mockBoardRepo = NewMockBoardRepository()
	repository.Tasks = mockRepo
	repository.Boards = mockBoardRepo
}

func tearDownTest() {
	// Mock repository doesn't need cleanup
}

func setupTestRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/task", GetTaskHandler)
	r.POST("/api/task", PostTaskHandler)
	r.PUT("/api/task/:id", PutTaskHandler)
	r.DELETE("/api/task/:id", DeleteTaskHandler)

	// Board routes
	r.GET("/api/boards", GetBoardHandler)
	r.GET("/api/boards/:id", GetBoardHandler)
	r.GET("/api/boards/:id/tasks", GetTasksByBoardHandler)
	r.POST("/api/boards", PostBoardHandler)
	r.PUT("/api/boards/:id", PutBoardHandler)
	r.DELETE("/api/boards/:id", DeleteBoardHandler)

	return r
}

// marshalTaskBody creates a JSON byte slice from task field values
func marshalTaskBody(title string, description string, status string, priority string) []byte {
	taskBody := map[string]any{}
	if title != "" {
		taskBody["title"] = title
	}
	if description != "" {
		taskBody["description"] = description
	}
	if status != "" {
		taskBody["status"] = status
	}
	if priority != "" {
		taskBody["priority"] = priority
	}
	jsonBody, _ := json.Marshal(taskBody)
	return jsonBody
}

// makePostRequest creates a POST request for testing
func makePostRequest(r *gin.Engine, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/api/task", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

// marshalBoardBody creates a JSON byte slice from board field values
func marshalBoardBody(name string, description string, color string) []byte {
	boardBody := map[string]any{}
	if name != "" {
		boardBody["name"] = name
	}
	if description != "" {
		boardBody["description"] = description
	}
	if color != "" {
		boardBody["color"] = color
	}
	jsonBody, _ := json.Marshal(boardBody)
	return jsonBody
}

// makePostBoardRequest creates a POST request to /api/boards for testing
func makePostBoardRequest(r *gin.Engine, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/api/boards", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

// makePutBoardRequest creates a PUT request to /api/boards/:id for testing
func makePutBoardRequest(r *gin.Engine, id string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("PUT", "/api/boards/"+id, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

// makeDeleteBoardRequest creates a DELETE request to /api/boards/:id for testing
func makeDeleteBoardRequest(r *gin.Engine, id string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("DELETE", "/api/boards/"+id, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

// makeGetBoardRequest creates a GET request to /api/boards for testing
func makeGetBoardRequest(r *gin.Engine) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", "/api/boards", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

// makeGetBoardByIDRequest creates a GET request to /api/boards/:id for testing
func makeGetBoardByIDRequest(r *gin.Engine, id string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", "/api/boards/"+id, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}
