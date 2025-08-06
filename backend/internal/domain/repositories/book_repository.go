package repositories

import "library-management-system/internal/domain/entities"

// BookRepository defines the interface for book data access
type BookRepository interface {
	Create(book *entities.Book) error
	GetByID(id string) (*entities.Book, error)
	GetAll() ([]entities.Book, error)
	Update(book *entities.Book) error
	Delete(id string) error
	HardDelete(id string) error
	FindByTitle(title string) ([]entities.Book, error)
	FindByAuthor(author string) ([]entities.Book, error)
	FindByYear(year int) ([]entities.Book, error)
	FindByISBN(isbn string) (*entities.Book, error)
	GetDeletedBooks() ([]entities.Book, error)
	Restore(id string) error
}
