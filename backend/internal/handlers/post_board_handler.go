package handlers

import (
	"net/http"
	"strings"

	types "tasker/internal/types"
	"tasker/internal/repository"

	"github.com/gin-gonic/gin"
)

func validateBoard(b types.Board) map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(b.Name) == "" {
		errors["name"] = "name is required"
	}

	// Optional color validation - validate hex color format
	if b.Color != "" {
		if strings.HasPrefix(b.Color, "#") && len(b.Color) != 7 {
			errors["color"] = "color must be a valid hex color code (e.g., #3B82F6)"
		}
	}

	return errors
}

// PostBoardHandler creates a new board
func PostBoardHandler(c *gin.Context) {
	var newBoard types.Board
	if err := c.ShouldBindJSON(&newBoard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if validationErrors := validateBoard(newBoard); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": validationErrors})
		return
	}

	if newBoard.Color == "" {
		newBoard.Color = "#3B82F6" // Default blue color
	}

	// Save via repository
	createdBoard, err := repository.Boards.CreateBoard(newBoard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save board"})
		return
	}

	c.JSON(http.StatusCreated, createdBoard)
}
