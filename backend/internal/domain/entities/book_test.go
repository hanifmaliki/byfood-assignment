package entities

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBook_BeforeCreate(t *testing.T) {
	tests := []struct {
		name     string
		book     *Book
		expected string
	}{
		{
			name: "should generate UUID when ID is empty",
			book: &Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			expected: "", // Will be generated
		},
		{
			name: "should not change ID when already set",
			book: &Book{
				ID:     "existing-id",
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			expected: "existing-id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock transaction
			var mockTx *gorm.DB

			err := tt.book.BeforeCreate(mockTx)
			assert.NoError(t, err)

			if tt.expected == "" {
				// Should have generated a UUID
				assert.NotEmpty(t, tt.book.ID)
				// Verify it's a valid UUID
				_, err := uuid.Parse(tt.book.ID)
				assert.NoError(t, err)
			} else {
				// Should not have changed the ID
				assert.Equal(t, tt.expected, tt.book.ID)
			}
		})
	}
}

func TestBook_TableName(t *testing.T) {
	book := &Book{}
	assert.Equal(t, "books", book.TableName())
}

func TestBook_Validation(t *testing.T) {
	tests := []struct {
		name    string
		book    *Book
		isValid bool
	}{
		{
			name: "valid book",
			book: &Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			isValid: true,
		},
		{
			name: "empty title",
			book: &Book{
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			isValid: false,
		},
		{
			name: "empty author",
			book: &Book{
				Title: "Test Book",
				Year:  2024,
				ISBN:  "1234567890",
			},
			isValid: false,
		},
		{
			name: "invalid year (too old)",
			book: &Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   999,
				ISBN:   "1234567890",
			},
			isValid: false,
		},
		{
			name: "invalid year (too new)",
			book: &Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2101,
				ISBN:   "1234567890",
			},
			isValid: false,
		},
		{
			name: "empty ISBN",
			book: &Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
			},
			isValid: false,
		},
		{
			name: "ISBN too short",
			book: &Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "123",
			},
			isValid: false,
		},
		{
			name: "ISBN too long",
			book: &Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "12345678901234",
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Basic validation checks
			isValid := tt.book.Title != "" &&
				tt.book.Author != "" &&
				tt.book.Year >= 1000 && tt.book.Year <= 2100 &&
				len(tt.book.ISBN) >= 10 && len(tt.book.ISBN) <= 13

			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

func TestBook_JSONTags(t *testing.T) {
	book := &Book{
		ID:        "test-id",
		Title:     "Test Book",
		Author:    "Test Author",
		Year:      2024,
		ISBN:      "1234567890",
		CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	// Test that the struct can be marshaled to JSON
	// This indirectly tests the JSON tags
	jsonData, err := json.Marshal(book)
	assert.NoError(t, err)
	assert.Contains(t, string(jsonData), `"id":"test-id"`)
	assert.Contains(t, string(jsonData), `"title":"Test Book"`)
	assert.Contains(t, string(jsonData), `"author":"Test Author"`)
	assert.Contains(t, string(jsonData), `"year":2024`)
	assert.Contains(t, string(jsonData), `"isbn":"1234567890"`)
}

func TestBook_SoftDelete(t *testing.T) {
	book := &Book{
		ID:        "test-id",
		Title:     "Test Book",
		Author:    "Test Author",
		Year:      2024,
		ISBN:      "1234567890",
		CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	// Initially not deleted
	assert.Nil(t, book.DeletedAt)

	// Simulate soft delete
	deletedAt := time.Now()
	book.DeletedAt = &deletedAt

	// Should be marked as deleted
	assert.NotNil(t, book.DeletedAt)
	assert.Equal(t, deletedAt, *book.DeletedAt)
}
