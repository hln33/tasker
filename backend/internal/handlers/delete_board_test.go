package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	types "tasker/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteBoardHandler_Success(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a dummy board first to get ID counter past 1
	dummyBody := marshalBoardBody("Dummy Board", "This is a dummy board", "#000000")
	makePostBoardRequest(r, dummyBody)

	// Create the actual board we want to delete
	boardBody := marshalBoardBody("Board to Delete", "This board will be deleted", "#FF5733")
	w1 := makePostBoardRequest(r, boardBody)

	var createdBoard types.Board
	json.Unmarshal(w1.Body.Bytes(), &createdBoard)

	// The created board should have ID 2
	assert.Equal(t, 2, createdBoard.ID)

	// Delete the board
	w2 := makeDeleteBoardRequest(r, fmt.Sprintf("%d", createdBoard.ID))

	assert.Equal(t, http.StatusNoContent, w2.Code)

	// Verify board is deleted
	w3 := makeGetBoardByIDRequest(r, fmt.Sprintf("%d", createdBoard.ID))
	assert.Equal(t, http.StatusNotFound, w3.Code)
}

func TestDeleteBoardHandler_BoardNotFound(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Try to delete a non-existent board
	w := makeDeleteBoardRequest(r, "999")

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "board not found", response["error"])
}

func TestDeleteBoardHandler_CannotDeleteDefaultBoard(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Try to delete the default board
	w := makeDeleteBoardRequest(r, "1")

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "cannot delete the default board", response["error"])
}

func TestDeleteBoardHandler_InvalidBoardID(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Try to delete with invalid ID (non-numeric)
	w := makeDeleteBoardRequest(r, "invalid")

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "invalid board ID", response["error"])
}
