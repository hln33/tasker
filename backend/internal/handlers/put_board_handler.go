package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	types "tasker/internal/types"
	"tasker/internal/repository"

	"github.com/gin-gonic/gin"
)

// PutBoardHandler updates an existing board
func PutBoardHandler(c *gin.Context) {
	boardIDParam := c.Param("id")
	if boardIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "board ID is required"})
		return
	}

	boardID, err := strconv.Atoi(boardIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid board ID"})
		return
	}

	var updateBoard types.Board
	if err := c.ShouldBindJSON(&updateBoard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// Validate only if fields are provided (partial update)
	if updateBoard.Name != "" && strings.TrimSpace(updateBoard.Name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": map[string]string{
			"name": "name cannot be empty",
		}})
		return
	}

	// Validate color if provided
	if updateBoard.Color != "" && strings.HasPrefix(updateBoard.Color, "#") && len(updateBoard.Color) != 7 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": map[string]string{
			"color": "color must be a valid hex color code (e.g., #3B82F6)",
		}})
		return
	}

	// Update via repository
	updatedBoard, err := repository.Boards.UpdateBoard(boardID, updateBoard)
	if err != nil {
		if err.Error() == fmt.Sprintf("board not found: %d", boardID) {
			c.JSON(http.StatusNotFound, gin.H{"error": "board not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update board"})
		return
	}

	c.JSON(http.StatusOK, updatedBoard)
}
