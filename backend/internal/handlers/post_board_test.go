package handlers

import (
	"encoding/json"
	"net/http"
	types "tasker/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostBoardHandler_Success(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	jsonBody := marshalBoardBody("Test Board", "Test Description", "#FF5733")
	w := makePostBoardRequest(r, jsonBody)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response types.Board
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, 1, response.ID)
	assert.Equal(t, "Test Board", response.Name)
	assert.Equal(t, "Test Description", response.Description)
	assert.Equal(t, "#FF5733", response.Color)
}

func TestPostBoardHandler_WithDefaults(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	jsonBody := marshalBoardBody("Minimal Board", "", "")
	w := makePostBoardRequest(r, jsonBody)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response types.Board
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "#3B82F6", response.Color) // Default blue color
	assert.Equal(t, "", response.Description)
}

func TestPostBoardHandler_MissingName(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	jsonBody := marshalBoardBody("", "Board without name", "#FF5733")
	w := makePostBoardRequest(r, jsonBody)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "validation failed", response["error"])

	details, ok := response["details"].(map[string]any)
	assert.True(t, ok)
	assert.NotNil(t, details["name"])
}

func TestPostBoardHandler_InvalidColor(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	jsonBody := marshalBoardBody("Test Board", "", "#12345")
	w := makePostBoardRequest(r, jsonBody)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	details, ok := response["details"].(map[string]any)
	assert.True(t, ok)
	assert.NotNil(t, details["color"])
}

func TestPostBoardHandler_InvalidJSON(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()
	w := makePostBoardRequest(r, []byte("invalid json"))

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "invalid JSON", response["error"])
}

func TestPostBoardHandler_IDIncrementing(t *testing.T) {
	setupTest()
	defer tearDownTest()

	r := setupTestRouter()

	// Create first board
	jsonBody1 := marshalBoardBody("Board 1", "", "#FF5733")
	w1 := makePostBoardRequest(r, jsonBody1)

	var response1 types.Board
	json.Unmarshal(w1.Body.Bytes(), &response1)
	assert.Equal(t, 1, response1.ID)

	// Create second board
	jsonBody2 := marshalBoardBody("Board 2", "", "#33FF57")
	w2 := makePostBoardRequest(r, jsonBody2)

	var response2 types.Board
	json.Unmarshal(w2.Body.Bytes(), &response2)
	assert.Equal(t, 2, response2.ID)
}

func TestValidateBoard(t *testing.T) {
	tests := []struct {
		name        string
		board       types.Board
		expectError bool
		errorFields []string
	}{
		{
			name: "Valid board",
			board: types.Board{
				Name:  "Valid Board Name",
				Color: "#3B82F6",
			},
			expectError: false,
		},
		{
			name: "Empty name",
			board: types.Board{
				Name: "",
			},
			expectError: true,
			errorFields: []string{"name"},
		},
		{
			name: "Whitespace only name",
			board: types.Board{
				Name: "   ",
			},
			expectError: true,
			errorFields: []string{"name"},
		},
		{
			name: "Invalid color format (too short)",
			board: types.Board{
				Name:  "Test",
				Color: "#12345",
			},
			expectError: true,
			errorFields: []string{"color"},
		},
		{
			name: "Invalid color format (too long)",
			board: types.Board{
				Name:  "Test",
				Color: "#1234567",
			},
			expectError: true,
			errorFields: []string{"color"},
		},
		{
			name: "Valid color (hex)",
			board: types.Board{
				Name:  "Test",
				Color: "#FF5733",
			},
			expectError: false,
		},
		{
			name: "Valid color (named - no validation yet)",
			board: types.Board{
				Name:  "Test",
				Color: "blue",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := validateBoard(tt.board)

			if tt.expectError && len(errors) == 0 {
				t.Error("Expected validation errors, got none")
			}

			if !tt.expectError && len(errors) > 0 {
				t.Errorf("Expected no errors, got %v", errors)
			}

			for _, field := range tt.errorFields {
				if errors[field] == "" {
					t.Errorf("Expected error for field %s, got none", field)
				}
			}
		})
	}
}
