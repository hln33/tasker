package handlers

import (
	"net/http"
	"strings"
	"tasker/internal/repository"

	"github.com/gin-gonic/gin"
)

// DeleteTaskHandler handles DELETE /api/task/:id requests
func DeleteTaskHandler(c *gin.Context) {
	taskID := c.Param("id")

	if err := repository.Tasks.DeleteTask(taskID); err != nil {
		if strings.Contains(err.Error(), "task not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
		return
	}

	c.Status(http.StatusNoContent) // 204 No Content
}
