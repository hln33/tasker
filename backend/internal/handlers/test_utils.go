package handlers

import (
	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/task", GetTaskHandler)
	r.POST("/api/task", PostTaskHandler)

	return r
}
