package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	types "tasker/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBoardHandler_GetAllBoards(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create some boards first
	board1Body := marshalBoardBody("Board 1", "Description 1", "#FF5733")
	makePostBoardRequest(r, board1Body)

	board2Body := marshalBoardBody("Board 2", "Description 2", "#33FF57")
	makePostBoardRequest(r, board2Body)

	// Get all boards
	w := makeGetBoardRequest(r)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []types.Board
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(response))
}

func TestGetBoardHandler_GetAllBoards_Empty(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Get all boards (none created)
	w := makeGetBoardRequest(r)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []types.Board
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, 0, len(response))
}

func TestGetBoardHandler_GetBoardByID(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a board
	boardBody := marshalBoardBody("Test Board", "Test Description", "#FF5733")
	w1 := makePostBoardRequest(r, boardBody)

	var createdBoard types.Board
	json.Unmarshal(w1.Body.Bytes(), &createdBoard)

	// Get the board by ID
	w2 := makeGetBoardByIDRequest(r, fmt.Sprintf("%d", createdBoard.ID))

	assert.Equal(t, http.StatusOK, w2.Code)

	var response types.Board
	err := json.Unmarshal(w2.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, createdBoard.ID, response.ID)
	assert.Equal(t, "Test Board", response.Name)
	assert.Equal(t, "Test Description", response.Description)
	assert.Equal(t, "#FF5733", response.Color)
}

func TestGetBoardHandler_GetBoardByID_NotFound(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Try to get a non-existent board
	w := makeGetBoardByIDRequest(r, "999")

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "board not found", response["error"])
}

func TestGetTasksByBoardHandler_EmptyTasks(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a board
	boardBody := marshalBoardBody("Test Board", "Test Description", "#FF5733")
	w1 := makePostBoardRequest(r, boardBody)

	var createdBoard types.Board
	json.Unmarshal(w1.Body.Bytes(), &createdBoard)

	// Get tasks for the board (should be empty)
	w2 := makeGetBoardTasksRequest(r, fmt.Sprintf("%d", createdBoard.ID))

	assert.Equal(t, http.StatusOK, w2.Code)

	var response []types.Task
	err := json.Unmarshal(w2.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, 0, len(response))
}
