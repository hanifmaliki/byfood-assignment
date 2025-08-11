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
	return &BookRepositoryImpl{db: db}
}

// Create creates a new book
func (r *BookRepositoryImpl) Create(book *entities.Book) error {
	return r.db.Create(book).Error
}

// GetByID retrieves a book by ID
func (r *BookRepositoryImpl) GetByID(id string) (*entities.Book, error) {
	var book entities.Book
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

// GetAll retrieves all books
func (r *BookRepositoryImpl) GetAll() ([]entities.Book, error) {
	var books []entities.Book
	err := r.db.Find(&books).Error
	return books, err
}

// Update updates a book
func (r *BookRepositoryImpl) Update(book *entities.Book) error {
	// Use Updates with a struct to only update non-zero fields and preserve timestamps
	// This automatically sets updated_at without affecting created_at
	return r.db.Model(book).Updates(book).Error
}

// Delete deletes a book (soft delete)
func (r *BookRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entities.Book{}, "id = ?", id).Error
}

// HardDelete permanently deletes a book
func (r *BookRepositoryImpl) HardDelete(id string) error {
	return r.db.Unscoped().Delete(&entities.Book{}, "id = ?", id).Error
}

// FindByTitle finds books by title (case-insensitive)
func (r *BookRepositoryImpl) FindByTitle(title string) ([]entities.Book, error) {
	var books []entities.Book
	err := r.db.Where("LOWER(title) LIKE LOWER(?)", "%"+title+"%").Find(&books).Error
	return books, err
}

// FindByAuthor finds books by author (case-insensitive)
func (r *BookRepositoryImpl) FindByAuthor(author string) ([]entities.Book, error) {
	var books []entities.Book
	err := r.db.Where("LOWER(author) LIKE LOWER(?)", "%"+author+"%").Find(&books).Error
	return books, err
}

// FindByYear finds books by year
func (r *BookRepositoryImpl) FindByYear(year int) ([]entities.Book, error) {
	var books []entities.Book
	err := r.db.Where("year = ?", year).Find(&books).Error
	return books, err
}

// FindByISBN finds a book by ISBN
func (r *BookRepositoryImpl) FindByISBN(isbn string) (*entities.Book, error) {
	var book entities.Book
	err := r.db.Where("isbn = ?", isbn).First(&book).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

// GetDeletedBooks retrieves all soft-deleted books
func (r *BookRepositoryImpl) GetDeletedBooks() ([]entities.Book, error) {
	var books []entities.Book
	err := r.db.Unscoped().Where("deleted_at IS NOT NULL").Find(&books).Error
	return books, err
}

// Restore restores a soft-deleted book
func (r *BookRepositoryImpl) Restore(id string) error {
	return r.db.Unscoped().Model(&entities.Book{}).Where("id = ?", id).Update("deleted_at", nil).Error
}
