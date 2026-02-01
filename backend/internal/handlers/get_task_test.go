package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTaskHandler(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a test task first
	jsonBody := marshalTaskBody("Setup project repository", "Initialize the repository with basic project structure", "In Progress", "High")
	postReq, _ := http.NewRequest("POST", "/api/task", bytes.NewBuffer(jsonBody))
	postReq.Header.Set("Content-Type", "application/json")
	postW := httptest.NewRecorder()
	r.ServeHTTP(postW, postReq)

	assert.Equal(t, http.StatusCreated, postW.Code)

	// Now test GET
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/task", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, "TASK-001")
	assert.Contains(t, body, "Setup project repository")
	assert.Contains(t, body, "Initialize the repository with basic project structure")
	assert.Contains(t, body, "In Progress")
	assert.Contains(t, body, "High")
}
