package handlers

import (
	"net/http"
	task "tasker/internal/Task"

	"github.com/gin-gonic/gin"
)

// returns a single task
func GetTaskHandler(c *gin.Context) {
	task := task.Task{
		ID:          "TASK-001",
		Title:       "Setup project repository",
		Description: "Initialize the repository with basic project structure",
		Status:      "In Progress",
		Priority:    "High",
	}
	c.JSON(http.StatusOK, task)
}
