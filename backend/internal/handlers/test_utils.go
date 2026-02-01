package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/task", GetTaskHandler)
	r.POST("/api/task", PostTaskHandler)

	return r
}

// marshalTaskBody creates a JSON byte slice from task field values
func marshalTaskBody(title string, description string, status string, priority string) []byte {
	taskBody := map[string]any{}
	if title != "" {
		taskBody["title"] = title
	}
	if description != "" {
		taskBody["description"] = description
	}
	if status != "" {
		taskBody["status"] = status
	}
	if priority != "" {
		taskBody["priority"] = priority
	}
	jsonBody, _ := json.Marshal(taskBody)
	return jsonBody
}
