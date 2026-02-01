package handlers

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	task "tasker/internal/Task"
	"tasker/internal/repository"

	"github.com/gin-gonic/gin"
)

var (
	nextID int = 1
)

// InitTaskIDGenerator initializes the ID generator from existing tasks
func InitTaskIDGenerator(existingTasks []task.Task) {
	for _, t := range existingTasks {
		var id int
		fmt.Sscanf(t.ID, "TASK-%d", &id)
		if id >= nextID {
			nextID = id + 1
		}
	}
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

// PostTaskHandler creates a new task
func PostTaskHandler(c *gin.Context) {
	var newTask task.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if validationErrors := validateTask(newTask); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": validationErrors})
		return
	}

	// Generate ID and set defaults
	newTask.ID = generateNextID()
	if newTask.Status == "" {
		newTask.Status = "TODO"
	}
	if newTask.Priority == "" {
		newTask.Priority = "Medium"
	}

	// Save via repository
	createdTask, err := repository.Tasks.CreateTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save task"})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}
