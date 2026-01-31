package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
}

var (
	tasks  []Task
	nextID int = 1
)

func loadTasks() {
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

func validateTask(task Task) map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(task.Title) == "" {
		errors["title"] = "title is required"
	}

	validStatuses := map[string]bool{"TODO": true, "In Progress": true, "Done": true}
	if task.Status != "" && !validStatuses[task.Status] {
		errors["status"] = "status must be one of: TODO, In Progress, Done"
	}

	validPriorities := map[string]bool{"Low": true, "Medium": true, "High": true}
	if task.Priority != "" && !validPriorities[task.Priority] {
		errors["priority"] = "priority must be one of: Low, Medium, High"
	}

	return errors
}

// postTaskHandler creates a new task
func postTaskHandler(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// Validate
	if validationErrors := validateTask(task); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": validationErrors})
		return
	}

	// Generate ID and set defaults
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
