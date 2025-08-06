package repository

import (
	"errors"

	"library-management-system/internal/domain/entities"
	"library-management-system/internal/domain/repositories"

	"gorm.io/gorm"
)

// BookRepositoryImpl implements the BookRepository interface
type BookRepositoryImpl struct {
	db *gorm.DB
}

// NewBookRepository creates a new book repository
func NewBookRepository(db *gorm.DB) repositories.BookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}

// Create saves a new book to the database
func (r *BookRepositoryImpl) Create(book *entities.Book) error {
	return r.db.Create(book).Error
}

// GetByID retrieves a book by ID
func (r *BookRepositoryImpl) GetByID(id string) (*entities.Book, error) {
	var book entities.Book
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
}

// GetAll retrieves all books
func (r *BookRepositoryImpl) GetAll() ([]*entities.Book, error) {
	var books []*entities.Book
	err := r.db.Order("created_at desc").Find(&books).Error
	return books, err
}

// Update updates an existing book
func (r *BookRepositoryImpl) Update(book *entities.Book) error {
	return r.db.Save(book).Error
}

// Delete removes a book by ID
func (r *BookRepositoryImpl) Delete(id string) error {
	result := r.db.Where("id = ?", id).Delete(&entities.Book{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("book not found")
	}
	return nil
}
