package handlers

import (
	"net/http"
	"tasker/internal/repository"

	"github.com/gin-gonic/gin"
)

// GetTaskHandler returns all tasks
func GetTaskHandler(c *gin.Context) {
	tasks, err := repository.Tasks.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
