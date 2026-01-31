package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	loadTasks()
}

// setupRouter configures and returns the Gin router with all routes
func setupRouter() *gin.Engine {
	r := gin.Default()

	// Configure CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/", healthCheckHandler)
	r.GET("/api/task", getTaskHandler)
	r.POST("/api/task", postTaskHandler)

	return r
}

// returns a hello world message
func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

// returns a single task
func getTaskHandler(c *gin.Context) {
	task := Task{
		ID:          "TASK-001",
		Title:       "Setup project repository",
		Description: "Initialize the repository with basic project structure",
		Status:      "In Progress",
		Priority:    "High",
	}
	c.JSON(http.StatusOK, task)
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
