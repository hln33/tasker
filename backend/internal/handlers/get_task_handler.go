package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	task "tasker/internal/Task"

	"github.com/gin-gonic/gin"
)

// returns a single task
func GetTaskHandler(c *gin.Context) {
	data, err := os.ReadFile("data.json")
	if err != nil || len(data) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No tasks found"})
		return
	}

	fetchedTasks := []task.Task{}
	if err := json.Unmarshal(data, &fetchedTasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse tasks"})
		return
	}

	c.JSON(http.StatusOK, fetchedTasks)
}
