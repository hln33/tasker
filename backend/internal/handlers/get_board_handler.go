package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	types "tasker/internal/types"
	"tasker/internal/repository"

	"github.com/gin-gonic/gin"
)

// GetBoardHandler handles GET requests for boards
// GET /api/boards - returns all boards
// GET /api/boards/:id - returns a specific board
func GetBoardHandler(c *gin.Context) {
	boardIDParam := c.Param("id")

	// If no ID provided, return all boards
	if boardIDParam == "" {
		boards, err := repository.Boards.GetAllBoards()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get boards"})
			return
		}

		c.JSON(http.StatusOK, boards)
		return
	}

	b, err := getBoardById(c, boardIDParam)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, b)
}

// GetTasksByBoardHandler returns all tasks for a specific board
func GetTasksByBoardHandler(c *gin.Context) {
	boardIDParam := c.Param("id")

	// First verify the board exists
	boardID, err := strconv.Atoi(boardIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid board ID"})
		return
	}

	_, err = getBoardById(c, boardIDParam)
	if err != nil {
		return
	}

	// Get tasks for this board
	tasks, err := repository.Boards.GetTasksByBoardID(boardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get tasks for board"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func getBoardById(c *gin.Context, boardIDParam string) (*types.Board, error) {
	boardID, err := strconv.Atoi(boardIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid board ID"})
		return nil, err
	}

	b, err := repository.Boards.GetBoardByID(boardID)
	if err != nil {
		if err.Error() == fmt.Sprintf("board not found: %d", boardID) {
			c.JSON(http.StatusNotFound, gin.H{"error": "board not found"})
			return nil, err
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get board"})
		return nil, err
	}
	return b, nil
}
