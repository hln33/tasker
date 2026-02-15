package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	types "tasker/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutBoardHandler_UpdateAllFields(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a board first
	boardBody := marshalBoardBody("Original Name", "Original Description", "#FF5733")
	w1 := makePostBoardRequest(r, boardBody)

	var createdBoard types.Board
	json.Unmarshal(w1.Body.Bytes(), &createdBoard)

	// Update the board
	updateBody := marshalBoardBody("Updated Name", "Updated Description", "#33FF57")
	w2 := makePutBoardRequest(r, fmt.Sprintf("%d", createdBoard.ID), updateBody)

	assert.Equal(t, http.StatusOK, w2.Code)

	var response types.Board
	err := json.Unmarshal(w2.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, createdBoard.ID, response.ID)
	assert.Equal(t, "Updated Name", response.Name)
	assert.Equal(t, "Updated Description", response.Description)
	assert.Equal(t, "#33FF57", response.Color)
}

func TestPutBoardHandler_PartialUpdate_Name(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a board first
	boardBody := marshalBoardBody("Original Name", "Original Description", "#FF5733")
	w1 := makePostBoardRequest(r, boardBody)

	var createdBoard types.Board
	json.Unmarshal(w1.Body.Bytes(), &createdBoard)

	// Update only the name
	updateBody := marshalBoardBody("Updated Name", "", "")
	w2 := makePutBoardRequest(r, fmt.Sprintf("%d", createdBoard.ID), updateBody)

	assert.Equal(t, http.StatusOK, w2.Code)

	var response types.Board
	json.Unmarshal(w2.Body.Bytes(), &response)

	assert.Equal(t, "Updated Name", response.Name)
	assert.Equal(t, "Original Description", response.Description) // Should remain unchanged
	assert.Equal(t, "#FF5733", response.Color)                     // Should remain unchanged
}

func TestPutBoardHandler_PartialUpdate_Color(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a board first
	boardBody := marshalBoardBody("Original Name", "Original Description", "#FF5733")
	w1 := makePostBoardRequest(r, boardBody)

	var createdBoard types.Board
	json.Unmarshal(w1.Body.Bytes(), &createdBoard)

	// Update only the color
	updateBody := marshalBoardBody("", "", "#123456")
	w2 := makePutBoardRequest(r, fmt.Sprintf("%d", createdBoard.ID), updateBody)

	assert.Equal(t, http.StatusOK, w2.Code)

	var response types.Board
	json.Unmarshal(w2.Body.Bytes(), &response)

	assert.Equal(t, "Original Name", response.Name)         // Should remain unchanged
	assert.Equal(t, "Original Description", response.Description) // Should remain unchanged
	assert.Equal(t, "#123456", response.Color)
}

func TestPutBoardHandler_BoardNotFound(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Try to update a non-existent board
	updateBody := marshalBoardBody("Updated Name", "", "#FF5733")
	w := makePutBoardRequest(r, "999", updateBody)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "board not found", response["error"])
}

func TestPutBoardHandler_EmptyName(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a board first
	boardBody := marshalBoardBody("Original Name", "Original Description", "#FF5733")
	w1 := makePostBoardRequest(r, boardBody)

	var createdBoard types.Board
	json.Unmarshal(w1.Body.Bytes(), &createdBoard)

	// Try to update with empty name
	updateBody := marshalBoardBody("   ", "", "#FF5733")
	w2 := makePutBoardRequest(r, fmt.Sprintf("%d", createdBoard.ID), updateBody)

	assert.Equal(t, http.StatusBadRequest, w2.Code)

	var response map[string]any
	json.Unmarshal(w2.Body.Bytes(), &response)

	assert.Equal(t, "validation failed", response["error"])

	details, ok := response["details"].(map[string]any)
	assert.True(t, ok)
	assert.NotNil(t, details["name"])
}

func TestPutBoardHandler_InvalidColorFormat(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create a board first
	boardBody := marshalBoardBody("Original Name", "Original Description", "#FF5733")
	w1 := makePostBoardRequest(r, boardBody)

	var createdBoard types.Board
	json.Unmarshal(w1.Body.Bytes(), &createdBoard)

	// Try to update with invalid color (too short)
	updateBody := marshalBoardBody("Updated Name", "", "#12345")
	w2 := makePutBoardRequest(r, fmt.Sprintf("%d", createdBoard.ID), updateBody)

	assert.Equal(t, http.StatusBadRequest, w2.Code)

	var response map[string]any
	json.Unmarshal(w2.Body.Bytes(), &response)

	details, ok := response["details"].(map[string]any)
	assert.True(t, ok)
	assert.NotNil(t, details["color"])
}

func TestPutBoardHandler_InvalidJSON(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Try to update with invalid JSON
	w := makePutBoardRequest(r, "1", []byte("invalid json"))

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "invalid JSON", response["error"])
}
