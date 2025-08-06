package entities

import (
	"time"

	"github.com/google/uuid"
)

// Book represents a book in the library domain
type Book struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	ISBN      string    `json:"isbn"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewBook creates a new book with a generated ID
func NewBook(title, author, isbn string, year int) *Book {
	return &Book{
		ID:     uuid.New().String(),
		Title:  title,
		Author: author,
		Year:   year,
		ISBN:   isbn,
	}
}
