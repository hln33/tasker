package handlers

import (
	"net/http"
	"slices"
	"strings"
	task "tasker/internal/Task"
	"tasker/internal/repository"

	"github.com/gin-gonic/gin"
)

// validateTaskUpdate validates only the fields that are being updated (non-empty)
// For PUT /api/task/:id, we allow partial updates, so we only validate provided fields
func validateTaskUpdate(t task.Task) map[string]string {
	errors := make(map[string]string)

	// Only validate title if it's being updated (non-empty)
	// Check if title is provided but consists only of whitespace
	if t.Title != "" && strings.TrimSpace(t.Title) == "" {
		errors["title"] = "title cannot be empty or whitespace only"
	}

	// Only validate status if it's being updated (non-empty)
	validStatuses := []string{"TODO", "In Progress", "Done"}
	if t.Status != "" && !slices.Contains(validStatuses, t.Status) {
		errors["status"] = "status must be one of: TODO, In Progress, Done"
	}

	// Only validate priority if it's being updated (non-empty)
	validPriorities := []string{"Low", "Medium", "High"}
	if t.Priority != "" && !slices.Contains(validPriorities, t.Priority) {
		errors["priority"] = "priority must be one of: Low, Medium, High"
	}

	return errors
}

// PutTaskHandler handles PUT /api/task/:id requests
func PutTaskHandler(c *gin.Context) {
	taskID := c.Param("id")

	var task task.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if validationErrors := validateTaskUpdate(task); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": validationErrors})
		return
	}

	updatedTask, err := repository.Tasks.UpdateTask(taskID, task)
	if err != nil {
		if strings.Contains(err.Error(), "task not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}
