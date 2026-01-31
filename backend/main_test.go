package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	// Set Gin to test mode to disable unnecessary logging
	gin.SetMode(gin.TestMode)

	// Create a test router
	r := setupRouter()

	// Create a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "hello world")
}

func TestGetTaskHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/task", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Assert JSON contains expected fields
	body := w.Body.String()
	assert.Contains(t, body, "TASK-001")
	assert.Contains(t, body, "Setup project repository")
	assert.Contains(t, body, "Initialize the repository with basic project structure")
	assert.Contains(t, body, "In Progress")
	assert.Contains(t, body, "High")
}

func TestNotFoundRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/nonexistent", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
