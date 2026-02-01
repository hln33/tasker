package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"
	task "tasker/internal/Task"

	"github.com/gin-gonic/gin"
)

var (
	tasks  []task.Task
	nextID int = 1
)

func LoadTasks() {
	data, err := os.ReadFile("data.json")
	if err != nil || len(data) == 0 {
		return
	}

	json.Unmarshal(data, &tasks)

	// Find the highest ID to avoid collisions
	for _, t := range tasks {
		var id int
		fmt.Sscanf(t.ID, "TASK-%d", &id)
		if id >= nextID {
			nextID = id + 1
		}
	}
}

func saveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("data.json", data, 0644)
}

func generateNextID() string {
	id := fmt.Sprintf("TASK-%03d", nextID)
	nextID++
	return id
}

func validateTask(task task.Task) map[string]string {
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

func PostTaskHandler(c *gin.Context) {
	var task task.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if validationErrors := validateTask(task); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": validationErrors})
		return
	}

	task.ID = generateNextID()
	if task.Status == "" {
		task.Status = "TODO"
	}
	if task.Priority == "" {
		task.Priority = "Medium"
	}

	// Save to file
	tasks = append(tasks, task)
	if err := saveTasks(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}
