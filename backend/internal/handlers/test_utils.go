package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"

	task "tasker/internal/Task"
	"tasker/internal/repository"

	"github.com/gin-gonic/gin"
)

// MockTaskRepository is an in-memory implementation for testing
type MockTaskRepository struct {
	tasks map[string]task.Task
	mu    sync.RWMutex
}

func NewMockTaskRepository() *MockTaskRepository {
	return &MockTaskRepository{
		tasks: make(map[string]task.Task),
	}
}

func (m *MockTaskRepository) GetAllTasks() ([]task.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]task.Task, 0, len(m.tasks))
	for _, t := range m.tasks {
		result = append(result, t)
	}
	return result, nil
}

func (m *MockTaskRepository) GetTaskByID(id string) (*task.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if t, ok := m.tasks[id]; ok {
		return &t, nil
	}
	return nil, errors.New("task not found: " + id)
}

func (m *MockTaskRepository) CreateTask(t task.Task) (*task.Task, error) {
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

func (m *MockTaskRepository) UpdateTask(id string, t task.Task) (*task.Task, error) {
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

	// Update timestamp
	existing.UpdatedAt = time.Now()

	m.tasks[id] = existing
	return &existing, nil
}

func (m *MockTaskRepository) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tasks = make(map[string]task.Task)
}

var mockRepo *MockTaskRepository

func setupTest() {
	gin.SetMode(gin.TestMode)
	nextID = 1
	mockRepo = NewMockTaskRepository()
	repository.Tasks = mockRepo
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
