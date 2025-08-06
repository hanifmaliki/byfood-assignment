package usecase

import (
	"errors"
	"time"

	"library-management-system/internal/domain/entities"
	"library-management-system/internal/domain/repositories"
)

// BookUseCase handles book business logic
type BookUseCase struct {
	bookRepo repositories.BookRepository
}

// NewBookUseCase creates a new book use case
func NewBookUseCase(bookRepo repositories.BookRepository) *BookUseCase {
	return &BookUseCase{
		bookRepo: bookRepo,
	}
}

// CreateBook creates a new book
func (uc *BookUseCase) CreateBook(title, author, isbn string, year int) (*entities.Book, error) {
	// Validate input
	if title == "" {
		return nil, errors.New("title is required")
	}
	if author == "" {
		return nil, errors.New("author is required")
	}
	if isbn == "" {
		return nil, errors.New("isbn is required")
	}
	if year < 1800 || year > 2024 {
		return nil, errors.New("year must be between 1800 and 2024")
	}

	book := entities.NewBook(title, author, isbn, year)
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	err := uc.bookRepo.Create(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// GetBook retrieves a book by ID
func (uc *BookUseCase) GetBook(id string) (*entities.Book, error) {
	if id == "" {
		return nil, errors.New("book ID is required")
	}

	return uc.bookRepo.GetByID(id)
}

// GetAllBooks retrieves all books
func (uc *BookUseCase) GetAllBooks() ([]*entities.Book, error) {
	return uc.bookRepo.GetAll()
}

// UpdateBook updates an existing book
func (uc *BookUseCase) UpdateBook(id, title, author, isbn string, year int) (*entities.Book, error) {
	if id == "" {
		return nil, errors.New("book ID is required")
	}

	book, err := uc.bookRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if title != "" {
		book.Title = title
	}
	if author != "" {
		book.Author = author
	}
	if isbn != "" {
		book.ISBN = isbn
	}
	if year > 0 {
		if year < 1800 || year > 2024 {
			return nil, errors.New("year must be between 1800 and 2024")
		}
		book.Year = year
	}

	book.UpdatedAt = time.Now()

	err = uc.bookRepo.Update(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// DeleteBook deletes a book by ID
func (uc *BookUseCase) DeleteBook(id string) error {
	if id == "" {
		return errors.New("book ID is required")
	}

	return uc.bookRepo.Delete(id)
}
