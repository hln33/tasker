package main

import (
	"encoding/json"
	"log"
	"os"
	task "tasker/internal/Task"
)

func main() {
	tasks := []task.Task{
		{
			ID:          "TASK-001",
			Title:       "Setup project repository",
			Description: "Initialize the repository with basic project structure and configuration files",
			Status:      "Done",
			Priority:    "High",
		},
		{
			ID:          "TASK-002",
			Title:       "Design database schema",
			Description: "Create the database schema for tasks, users, and project assignments",
			Status:      "In Progress",
			Priority:    "High",
		},
		{
			ID:          "TASK-003",
			Title:       "Implement user authentication",
			Description: "Add login/logout functionality with JWT token-based authentication",
			Status:      "TODO",
			Priority:    "High",
		},
		{
			ID:          "TASK-004",
			Title:       "Create task CRUD API",
			Description: "Implement create, read, update, and delete endpoints for task management",
			Status:      "In Progress",
			Priority:    "Medium",
		},
		{
			ID:          "TASK-005",
			Title:       "Write unit tests",
			Description: "Add comprehensive unit tests for all API endpoints and business logic",
			Status:      "TODO",
			Priority:    "Low",
		},
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Write to backend/data.json (project root relative to script)
	outputPath := "data.json"
	if len(os.Args) > 1 {
		outputPath = os.Args[1]
	}

	err = os.WriteFile(outputPath, data, 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	log.Printf("Successfully generated %s with %d tasks", outputPath, len(tasks))
}
