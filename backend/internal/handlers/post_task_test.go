package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	task "tasker/internal/Task"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTest() {
	gin.SetMode(gin.TestMode)
	tasks = nil
	nextID = 1
	os.Remove("data.json")
}

func tearDownTest() {
	os.Remove("data.json")
}

func makePostRequest(r *gin.Engine, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/api/task", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

func TestPostTaskHandler_Success(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	taskBody := map[string]any{
		"title":       "Test task.Task",
		"description": "Test Description",
		"status":      "In Progress",
		"priority":    "High",
	}
	jsonBody, _ := json.Marshal(taskBody)
	w := makePostRequest(r, jsonBody)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response task.Task
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "TASK-001", response.ID)
	assert.Equal(t, "Test task.Task", response.Title)
	assert.Equal(t, "Test Description", response.Description)
	assert.Equal(t, "In Progress", response.Status)
	assert.Equal(t, "High", response.Priority)
}

func TestPostTaskHandler_WithDefaults(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	taskBody := map[string]any{
		"title": "Minimal task.Task",
	}
	jsonBody, _ := json.Marshal(taskBody)
	w := makePostRequest(r, jsonBody)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response task.Task
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "TODO", response.Status)
	assert.Equal(t, "Medium", response.Priority)
	assert.Equal(t, "", response.Description)
}

func TestPostTaskHandler_MissingTitle(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	taskBody := map[string]any{
		"description": "task.Task without title",
	}
	jsonBody, _ := json.Marshal(taskBody)
	w := makePostRequest(r, jsonBody)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "validation failed", response["error"])

	details, ok := response["details"].(map[string]any)
	assert.True(t, ok)
	assert.NotNil(t, details["title"])
}

func TestPostTaskHandler_InvalidStatus(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	taskBody := map[string]any{
		"title":  "Test task.Task",
		"status": "InvalidStatus",
	}
	jsonBody, _ := json.Marshal(taskBody)
	w := makePostRequest(r, jsonBody)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	details, ok := response["details"].(map[string]any)
	assert.True(t, ok)
	assert.NotNil(t, details["status"])
}

func TestPostTaskHandler_InvalidPriority(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	taskBody := map[string]any{
		"title":    "Test task.Task",
		"priority": "InvalidPriority",
	}
	jsonBody, _ := json.Marshal(taskBody)
	w := makePostRequest(r, jsonBody)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	details, ok := response["details"].(map[string]any)
	assert.True(t, ok)
	assert.NotNil(t, details["priority"])
}

func TestPostTaskHandler_InvalidJSON(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()
	w := makePostRequest(r, []byte("invalid json"))

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "invalid JSON", response["error"])
}

func TestPostTaskHandler_IDIncrementing(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create first task
	taskBody1 := map[string]any{"title": "task.Task 1"}
	jsonBody1, _ := json.Marshal(taskBody1)
	w1 := makePostRequest(r, jsonBody1)

	var response1 task.Task
	json.Unmarshal(w1.Body.Bytes(), &response1)
	assert.Equal(t, "TASK-001", response1.ID)

	// Create second task
	taskBody2 := map[string]any{"title": "task.Task 2"}
	jsonBody2, _ := json.Marshal(taskBody2)
	w2 := makePostRequest(r, jsonBody2)

	var response2 task.Task
	json.Unmarshal(w2.Body.Bytes(), &response2)
	assert.Equal(t, "TASK-002", response2.ID)
}

func TestValidateTask(t *testing.T) {
	tests := []struct {
		name        string
		task        task.Task
		expectError bool
		errorFields []string
	}{
		{
			name: "Valid task",
			task: task.Task{
				Title:    "Valid Title",
				Status:   "TODO",
				Priority: "High",
			},
			expectError: false,
		},
		{
			name: "Empty title",
			task: task.Task{
				Title: "",
			},
			expectError: true,
			errorFields: []string{"title"},
		},
		{
			name: "Whitespace only title",
			task: task.Task{
				Title: "   ",
			},
			expectError: true,
			errorFields: []string{"title"},
		},
		{
			name: "Invalid status",
			task: task.Task{
				Title:  "Valid Title",
				Status: "Invalid",
			},
			expectError: true,
			errorFields: []string{"status"},
		},
		{
			name: "Invalid priority",
			task: task.Task{
				Title:    "Valid Title",
				Priority: "Invalid",
			},
			expectError: true,
			errorFields: []string{"priority"},
		},
		{
			name: "Multiple errors",
			task: task.Task{
				Title:    "",
				Status:   "Invalid",
				Priority: "Invalid",
			},
			expectError: true,
			errorFields: []string{"title", "status", "priority"},
		},
		{
			name: "Valid status values",
			task: task.Task{
				Title:  "Test",
				Status: "In Progress",
			},
			expectError: false,
		},
		{
			name: "Valid priority values",
			task: task.Task{
				Title:    "Test",
				Priority: "Low",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := validateTask(tt.task)

			if tt.expectError && len(errors) == 0 {
				t.Error("Expected validation errors, got none")
			}

			if !tt.expectError && len(errors) > 0 {
				t.Errorf("Expected no errors, got %v", errors)
			}

			for _, field := range tt.errorFields {
				if errors[field] == "" {
					t.Errorf("Expected error for field %s, got none", field)
				}
			}
		})
	}
}

func TestGenerateNextID(t *testing.T) {
	setupTest()
	defer tearDownTest()

	tests := []struct {
		name     string
		setupID  int
		expected string
	}{
		{"First ID", 1, "TASK-001"},
		{"Tenth ID", 10, "TASK-010"},
		{"Hundredth ID", 100, "TASK-100"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextID = tt.setupID
			result := generateNextID()

			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.setupID+1, nextID)
		})
	}
}

func TestSaveAndLoadTasks(t *testing.T) {
	setupTest()
	defer tearDownTest()

	tasks = []task.Task{
		{ID: "TASK-001", Title: "task.Task 1", Status: "TODO", Priority: "High"},
		{ID: "TASK-005", Title: "task.Task 2", Status: "Done", Priority: "Low"},
	}

	err := saveTasks()
	assert.NoError(t, err)

	if _, err := os.Stat("data.json"); os.IsNotExist(err) {
		t.Error("data.json file was not created")
	}

	tasks = nil
	nextID = 1
	LoadTasks()

	assert.Equal(t, 2, len(tasks))
	assert.Equal(t, 6, nextID) // Should be 6 since TASK-005 exists
}

func TestLoadTasks_EmptyFile(t *testing.T) {
	setupTest()
	defer tearDownTest()

	os.WriteFile("data.json", []byte(""), 0644)

	tasks = []task.Task{{ID: "test"}}
	nextID = 999

	LoadTasks()

	// Should not modify tasks if file is empty
	assert.Equal(t, 1, len(tasks))
}

func TestLoadTasks_NonExistentFile(t *testing.T) {
	setupTest()
	defer tearDownTest()

	os.Remove("data.json")

	tasks = []task.Task{{ID: "test"}}
	nextID = 999

	LoadTasks()

	// Should not crash or modify tasks
	assert.Equal(t, 1, len(tasks))
}
