package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTaskHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := setupTestRouter()
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
