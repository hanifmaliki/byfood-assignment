package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Book represents a book entity
type Book struct {
	ID        string     `json:"id" gorm:"primaryKey;type:uuid"`
	Title     string     `json:"title" gorm:"not null;index"`
	Author    string     `json:"author" gorm:"not null;index"`
	Year      int        `json:"year" gorm:"not null;index"`
	ISBN      string     `json:"isbn" gorm:"uniqueIndex;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime;index"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

// BeforeCreate is called before creating a new book
func (b *Book) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return nil
}

// TableName returns the table name for the Book entity
func (Book) TableName() string {
	return "books"
}
