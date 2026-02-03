package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"
	task "tasker/internal/Task"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func makePutRequest(r *gin.Engine, id string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("PUT", "/api/task/"+id, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

func TestPutTaskHandler_Success_FullUpdate(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// First, create a task
	jsonBody := marshalTaskBody("Original Title", "Original Description", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Wait a bit to ensure UpdatedAt will be different
	time.Sleep(10 * time.Millisecond)

	// Update all fields
	updateBody := marshalTaskBody("Updated Title", "Updated Description", "Done", "High")
	putW := makePutRequest(r, createdTask.ID, updateBody)

	assert.Equal(t, http.StatusOK, putW.Code)

	var response task.Task
	json.Unmarshal(putW.Body.Bytes(), &response)

	assert.Equal(t, createdTask.ID, response.ID)
	assert.Equal(t, "Updated Title", response.Title)
	assert.Equal(t, "Updated Description", response.Description)
	assert.Equal(t, "Done", response.Status)
	assert.Equal(t, "High", response.Priority)
	assert.True(t, response.UpdatedAt.After(createdTask.UpdatedAt))
}

func TestPutTaskHandler_Success_PartialUpdate(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a task
	jsonBody := marshalTaskBody("Original Title", "Original Description", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Update only status
	updateBody := marshalTaskBody("", "", "In Progress", "")
	putW := makePutRequest(r, createdTask.ID, updateBody)

	assert.Equal(t, http.StatusOK, putW.Code)

	var response task.Task
	json.Unmarshal(putW.Body.Bytes(), &response)

	assert.Equal(t, createdTask.ID, response.ID)
	assert.Equal(t, "Original Title", response.Title)
	assert.Equal(t, "Original Description", response.Description)
	assert.Equal(t, "In Progress", response.Status)
	assert.Equal(t, "Medium", response.Priority)
}

func TestPutTaskHandler_TaskNotFound(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	updateBody := marshalTaskBody("Updated Title", "Updated Description", "Done", "High")
	putW := makePutRequest(r, "NON-EXISTENT", updateBody)

	assert.Equal(t, http.StatusNotFound, putW.Code)

	var response map[string]any
	json.Unmarshal(putW.Body.Bytes(), &response)

	assert.Equal(t, "task not found", response["error"])
}

func TestPutTaskHandler_InvalidStatus(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a task
	jsonBody := marshalTaskBody("Test Task", "Test Description", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Try to update with invalid status
	updateBody := marshalTaskBody("Updated Title", "Updated Description", "InvalidStatus", "High")
	putW := makePutRequest(r, createdTask.ID, updateBody)

	assert.Equal(t, http.StatusBadRequest, putW.Code)

	var response map[string]any
	json.Unmarshal(putW.Body.Bytes(), &response)

	assert.Equal(t, "validation failed", response["error"])

	details, ok := response["details"].(map[string]any)
	assert.True(t, ok)
	assert.NotNil(t, details["status"])
}

func TestPutTaskHandler_InvalidPriority(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a task
	jsonBody := marshalTaskBody("Test Task", "Test Description", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Try to update with invalid priority
	updateBody := marshalTaskBody("Updated Title", "Updated Description", "Done", "InvalidPriority")
	putW := makePutRequest(r, createdTask.ID, updateBody)

	assert.Equal(t, http.StatusBadRequest, putW.Code)

	var response map[string]any
	json.Unmarshal(putW.Body.Bytes(), &response)

	details, ok := response["details"].(map[string]any)
	assert.True(t, ok)
	assert.NotNil(t, details["priority"])
}

func TestPutTaskHandler_EmptyTitle(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a task
	jsonBody := marshalTaskBody("Original Title", "Original Description", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Empty title should be treated as "no change", not a validation error
	// This allows partial updates
	updateBody := marshalTaskBody("", "Updated Description", "Done", "High")
	putW := makePutRequest(r, createdTask.ID, updateBody)

	// Should succeed because empty title means "skip updating title"
	assert.Equal(t, http.StatusOK, putW.Code)

	var response task.Task
	json.Unmarshal(putW.Body.Bytes(), &response)

	// Title should remain unchanged (not set to empty)
	assert.Equal(t, "Original Title", response.Title)
	assert.Equal(t, "Updated Description", response.Description)
	assert.Equal(t, "Done", response.Status)
	assert.Equal(t, "High", response.Priority)
}

func TestPutTaskHandler_MultipleValidationErrors(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a task
	jsonBody := marshalTaskBody("Original Title", "Original Description", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Try to update with invalid status and priority (empty title is treated as "no change")
	updateBody := marshalTaskBody("", "Updated Description", "InvalidStatus", "InvalidPriority")
	putW := makePutRequest(r, createdTask.ID, updateBody)

	assert.Equal(t, http.StatusBadRequest, putW.Code)

	var response map[string]any
	json.Unmarshal(putW.Body.Bytes(), &response)

	details, ok := response["details"].(map[string]any)
	assert.True(t, ok)
	// Empty title is not a validation error (it means "no change")
	assert.Nil(t, details["title"])
	assert.NotNil(t, details["status"])
	assert.NotNil(t, details["priority"])
}

func TestPutTaskHandler_InvalidJSON(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a task
	jsonBody := marshalTaskBody("Original Title", "Original Description", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Try to update with invalid JSON
	putW := makePutRequest(r, createdTask.ID, []byte("invalid json"))

	assert.Equal(t, http.StatusBadRequest, putW.Code)

	var response map[string]any
	json.Unmarshal(putW.Body.Bytes(), &response)

	assert.Equal(t, "invalid JSON", response["error"])
}

func TestPutTaskHandler_UpdatedAtTimestampUpdates(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a task
	jsonBody := marshalTaskBody("Test Task", "Test Description", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)
	originalUpdatedAt := createdTask.UpdatedAt

	// Wait to ensure timestamp difference
	time.Sleep(10 * time.Millisecond)

	// Update the task
	updateBody := marshalTaskBody("Updated Title", "", "", "")
	putW := makePutRequest(r, createdTask.ID, updateBody)

	assert.Equal(t, http.StatusOK, putW.Code)

	var response task.Task
	json.Unmarshal(putW.Body.Bytes(), &response)

	assert.True(t, response.UpdatedAt.After(originalUpdatedAt))
}

func TestPutTaskHandler_StatusOnlyUpdate(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a task
	jsonBody := marshalTaskBody("Test Task", "Test Description", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Update only status
	updateBody := marshalTaskBody("", "", "Done", "")
	putW := makePutRequest(r, createdTask.ID, updateBody)

	assert.Equal(t, http.StatusOK, putW.Code)

	var response task.Task
	json.Unmarshal(putW.Body.Bytes(), &response)

	assert.Equal(t, "Test Task", response.Title)
	assert.Equal(t, "Test Description", response.Description)
	assert.Equal(t, "Done", response.Status)
	assert.Equal(t, "Medium", response.Priority)
}

func TestPutTaskHandler_PriorityOnlyUpdate(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a task
	jsonBody := marshalTaskBody("Test Task", "Test Description", "TODO", "Medium")
	postW := makePostRequest(r, jsonBody)

	var createdTask task.Task
	json.Unmarshal(postW.Body.Bytes(), &createdTask)

	// Update only priority
	updateBody := marshalTaskBody("", "", "", "High")
	putW := makePutRequest(r, createdTask.ID, updateBody)

	assert.Equal(t, http.StatusOK, putW.Code)

	var response task.Task
	json.Unmarshal(putW.Body.Bytes(), &response)

	assert.Equal(t, "Test Task", response.Title)
	assert.Equal(t, "Test Description", response.Description)
	assert.Equal(t, "TODO", response.Status)
	assert.Equal(t, "High", response.Priority)
}

func TestPutTaskHandler_EmptyID(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	updateBody := marshalTaskBody("Updated Title", "Updated Description", "Done", "High")
	putW := makePutRequest(r, "", updateBody)

	assert.Equal(t, http.StatusNotFound, putW.Code)
}
