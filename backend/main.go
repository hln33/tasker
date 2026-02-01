package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"tasker/internal/config"
	"tasker/internal/database"
	"tasker/internal/handlers"
	"tasker/internal/repository"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: error loading .env file: %v", err)
	}

	// Load configuration
	cfg := config.Load()

	// Connect to database with retry logic for containerized environments
	var db *sqlx.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = database.Connect(cfg.DatabaseURL())
		if err == nil {
			break
		}
		log.Printf("Database connection attempt %d failed: %v", i+1, err)
		if i < 9 {
			time.Sleep(2 * time.Second)
		}
	}
	if err != nil {
		log.Fatalf("Failed to connect to database after 10 attempts: %v", err)
	}
	defer database.Close()

	// Run migrations
	if err := runMigrations(cfg); err != nil {
		log.Printf("Warning: failed to run migrations: %v", err)
	}

	// Initialize repository
	repository.Tasks = repository.NewTaskRepository(db)

	// Initialize ID generator from existing tasks
	existingTasks, _ := repository.Tasks.GetAllTasks()
	if len(existingTasks) > 0 {
		handlers.InitTaskIDGenerator(existingTasks)
	}

	// Setup router
	r := setupRouter()

	// Start server with graceful shutdown
	startServer(r)
}

func runMigrations(cfg *config.Config) error {
	migrationFiles := []string{
		"migrations/000001_create_tasks_table.up.sql",
	}

	for _, file := range migrationFiles {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		if _, err := database.DB.Exec(string(content)); err != nil {
			return err
		}
		log.Printf("Migration completed: %s", file)
	}

	return nil
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/", healthCheckHandler)
	r.GET("/api/task", handlers.GetTaskHandler)
	r.POST("/api/task", handlers.PostTaskHandler)
	r.DELETE("/api/task/:id", handlers.DeleteTaskHandler)

	return r
}

func startServer(r *gin.Engine) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		log.Println("Server starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Give the server 5 seconds to complete existing requests
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
