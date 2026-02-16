package handlers

import (
	"net/http"
	"slices"
	"strings"

	types "tasker/internal/types"
	"tasker/internal/repository"

	"github.com/gin-gonic/gin"
)

func validateTask(task types.Task) map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(task.Title) == "" {
		errors["title"] = "title is required"
	}

	validStatuses := []string{"TODO", "In Progress", "Done"}
	if task.Status != "" && !slices.Contains(validStatuses, task.Status) {
		errors["status"] = "status must be one of: TODO, In Progress, Done"
	}

	validPriorities := []string{"Low", "Medium", "High"}
	if task.Priority != "" && !slices.Contains(validPriorities, task.Priority) {
		errors["priority"] = "priority must be one of: Low, Medium, High"
	}

	return errors
}

// PostTaskHandler creates a new task
func PostTaskHandler(c *gin.Context) {
	var newTask types.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if validationErrors := validateTask(newTask); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": validationErrors})
		return
	}

	// Set defaults
	if newTask.Status == "" {
		newTask.Status = "TODO"
	}
	if newTask.Priority == "" {
		newTask.Priority = "Medium"
	}
	if newTask.BoardID == 0 {
		newTask.BoardID = 1 // Default to the default board
	}

	// Save via repository
	createdTask, err := repository.Tasks.CreateTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save task"})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}
