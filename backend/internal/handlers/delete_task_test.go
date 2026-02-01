package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	task "tasker/internal/Task"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func makeDeleteRequest(r *gin.Engine, id string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("DELETE", "/api/task/"+id, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

func TestDeleteTaskHandler_Success(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// First, create a task
	jsonBody := marshalTaskBody("Task to delete", "This will be deleted", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	assert.Equal(t, http.StatusCreated, postW.Code)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Now delete it
	deleteW := makeDeleteRequest(r, createdTask.ID)
	assert.Equal(t, http.StatusNoContent, deleteW.Code)

	// Verify it's gone
	getW := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/task", nil)
	r.ServeHTTP(getW, req)

	var remainingTasks []task.Task
	json.Unmarshal(getW.Body.Bytes(), &remainingTasks)

	assert.Equal(t, 0, len(remainingTasks))
}

func TestDeleteTaskHandler_TaskNotFound(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	deleteW := makeDeleteRequest(r, "NON-EXISTENT")
	assert.Equal(t, http.StatusNotFound, deleteW.Code)

	var response map[string]any
	json.Unmarshal(deleteW.Body.Bytes(), &response)

	assert.Equal(t, "task not found", response["error"])
}

func TestDeleteTaskHandler_EmptyID(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	deleteW := makeDeleteRequest(r, "")
	assert.Equal(t, http.StatusNotFound, deleteW.Code)
}

func TestDeleteTaskHandler_MultipleTasks(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create multiple tasks
	jsonBody1 := marshalTaskBody("Task 1", "", "TODO", "High")
	postW1 := makePostRequest(r, jsonBody1)

	jsonBody2 := marshalTaskBody("Task 2", "", "In Progress", "Medium")
	postW2 := makePostRequest(r, jsonBody2)

	jsonBody3 := marshalTaskBody("Task 3", "", "Done", "Low")
	postW3 := makePostRequest(r, jsonBody3)

	var task1, task2, task3 task.Task
	json.Unmarshal(postW1.Body.Bytes(), &task1)
	json.Unmarshal(postW2.Body.Bytes(), &task2)
	json.Unmarshal(postW3.Body.Bytes(), &task3)

	// Delete middle task
	deleteW := makeDeleteRequest(r, task2.ID)
	assert.Equal(t, http.StatusNoContent, deleteW.Code)

	// Verify correct tasks remain
	getW := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/task", nil)
	r.ServeHTTP(getW, req)

	var remainingTasks []task.Task
	json.Unmarshal(getW.Body.Bytes(), &remainingTasks)

	assert.Equal(t, 2, len(remainingTasks))

	// Verify the remaining tasks are task1 and task3
	taskIDs := make(map[string]bool)
	for _, task := range remainingTasks {
		taskIDs[task.ID] = true
	}

	assert.True(t, taskIDs[task1.ID])
	assert.True(t, taskIDs[task3.ID])
	assert.False(t, taskIDs[task2.ID])
}

func TestDeleteTaskHandler_Persistence(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a task
	jsonBody := marshalTaskBody("Persistent task", "", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Delete it
	deleteW := makeDeleteRequest(r, createdTask.ID)
	assert.Equal(t, http.StatusNoContent, deleteW.Code)

	// Verify file was updated
	if _, err := os.Stat("data.json"); os.IsNotExist(err) {
		t.Error("data.json file should exist")
	}

	data, err := os.ReadFile("data.json")
	assert.NoError(t, err)

	var savedTasks []task.Task
	json.Unmarshal(data, &savedTasks)

	assert.Equal(t, 0, len(savedTasks))
}

func TestDeleteTaskHandler_DeleteTwice(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a task
	jsonBody := marshalTaskBody("Task to delete twice", "", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Delete it first time
	deleteW1 := makeDeleteRequest(r, createdTask.ID)
	assert.Equal(t, http.StatusNoContent, deleteW1.Code)

	// Try to delete it again
	deleteW2 := makeDeleteRequest(r, createdTask.ID)
	assert.Equal(t, http.StatusNotFound, deleteW2.Code)

	var response map[string]any
	json.Unmarshal(deleteW2.Body.Bytes(), &response)

	assert.Equal(t, "task not found", response["error"])
}
