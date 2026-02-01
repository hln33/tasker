package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteTaskHandler handles DELETE /api/task/:id requests
func DeleteTaskHandler(c *gin.Context) {
	taskID := c.Param("id")

	// Find the task by ID
	found := false
	for i, task := range tasks {
		if task.ID == taskID {
			// Remove task from slice
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	// Save to file
	if err := saveTasks(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save tasks"})
		return
	}

	c.Status(http.StatusNoContent) // 204 No Content
}
