package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"tasker/internal/repository"

	"github.com/gin-gonic/gin"
)

// DeleteBoardHandler deletes a board
func DeleteBoardHandler(c *gin.Context) {
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

	isDefaultBoard := boardID == 1
	if isDefaultBoard {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete the default board"})
		return
	}

	err = repository.Boards.DeleteBoard(boardID)
	if err != nil {
		if err.Error() == fmt.Sprintf("board not found: %d", boardID) {
			c.JSON(http.StatusNotFound, gin.H{"error": "board not found"})
			return
		}
		// Check if it's a "has tasks" error
		if strings.Contains(err.Error(), "cannot delete board with") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete board"})
		return
	}

	c.Status(http.StatusNoContent)
}
