package usecase

import (
	"errors"
	"strconv"

	"library-management-system/internal/domain/entities"
	"library-management-system/internal/domain/repositories"
)

// BookUseCase implements book business logic
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
func (uc *BookUseCase) CreateBook(book *entities.Book) error {
	// Validate book data
	if err := uc.validateBook(book); err != nil {
		return err
	}

	// Check if ISBN already exists
	existingBook, err := uc.bookRepo.FindByISBN(book.ISBN)
	if err != nil {
		return err
	}
	if existingBook != nil {
		return errors.New("book with this ISBN already exists")
	}

	return uc.bookRepo.Create(book)
}

// GetBook retrieves a book by ID
func (uc *BookUseCase) GetBook(id string) (*entities.Book, error) {
	if id == "" {
		return nil, errors.New("book ID is required")
	}

	return uc.bookRepo.GetByID(id)
}

// GetAllBooks retrieves all books
func (uc *BookUseCase) GetAllBooks() ([]entities.Book, error) {
	return uc.bookRepo.GetAll()
}

// UpdateBook updates an existing book
func (uc *BookUseCase) UpdateBook(id string, book *entities.Book) error {
	if id == "" {
		return errors.New("book ID is required")
	}

	// Validate book data
	if err := uc.validateBook(book); err != nil {
		return err
	}

	// Check if book exists
	existingBook, err := uc.bookRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingBook == nil {
		return errors.New("book not found")
	}

	// Check if ISBN is being changed and if it already exists
	if book.ISBN != existingBook.ISBN {
		bookWithISBN, err := uc.bookRepo.FindByISBN(book.ISBN)
		if err != nil {
			return err
		}
		if bookWithISBN != nil {
			return errors.New("book with this ISBN already exists")
		}
	}

	// Preserve existing data and update only the provided fields
	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.Year = book.Year
	existingBook.ISBN = book.ISBN

	return uc.bookRepo.Update(existingBook)
}

// DeleteBook deletes a book (soft delete)
func (uc *BookUseCase) DeleteBook(id string) error {
	if id == "" {
		return errors.New("book ID is required")
	}

	// Check if book exists
	existingBook, err := uc.bookRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingBook == nil {
		return errors.New("book not found")
	}

	return uc.bookRepo.Delete(id)
}

// HardDeleteBook permanently deletes a book
func (uc *BookUseCase) HardDeleteBook(id string) error {
	if id == "" {
		return errors.New("book ID is required")
	}

	// Check if book exists
	existingBook, err := uc.bookRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingBook == nil {
		return errors.New("book not found")
	}

	return uc.bookRepo.HardDelete(id)
}

// SearchBooksByTitle searches books by title
func (uc *BookUseCase) SearchBooksByTitle(title string) ([]entities.Book, error) {
	if title == "" {
		return nil, errors.New("title is required for search")
	}

	return uc.bookRepo.FindByTitle(title)
}

// SearchBooksByAuthor searches books by author
func (uc *BookUseCase) SearchBooksByAuthor(author string) ([]entities.Book, error) {
	if author == "" {
		return nil, errors.New("author is required for search")
	}

	return uc.bookRepo.FindByAuthor(author)
}

// SearchBooksByYear searches books by year
func (uc *BookUseCase) SearchBooksByYear(yearStr string) ([]entities.Book, error) {
	if yearStr == "" {
		return nil, errors.New("year is required for search")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return nil, errors.New("invalid year format")
	}

	return uc.bookRepo.FindByYear(year)
}

// GetDeletedBooks retrieves all soft-deleted books
func (uc *BookUseCase) GetDeletedBooks() ([]entities.Book, error) {
	return uc.bookRepo.GetDeletedBooks()
}

// RestoreBook restores a soft-deleted book
func (uc *BookUseCase) RestoreBook(id string) error {
	if id == "" {
		return errors.New("book ID is required")
	}

	return uc.bookRepo.Restore(id)
}

// validateBook validates book data
func (uc *BookUseCase) validateBook(book *entities.Book) error {
	if book.Title == "" {
		return errors.New("book title is required")
	}
	if book.Author == "" {
		return errors.New("book author is required")
	}
	if book.Year < 1000 || book.Year > 2100 {
		return errors.New("book year must be between 1000 and 2100")
	}
	if book.ISBN == "" {
		return errors.New("book ISBN is required")
	}
	if len(book.ISBN) < 10 || len(book.ISBN) > 13 {
		return errors.New("book ISBN must be between 10 and 13 characters")
	}

	return nil
}
