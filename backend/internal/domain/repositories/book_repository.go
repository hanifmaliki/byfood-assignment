package repositories

import "library-management-system/internal/domain/entities"

// BookRepository defines the interface for book data access
type BookRepository interface {
	Create(book *entities.Book) error
	GetByID(id string) (*entities.Book, error)
	GetAll() ([]*entities.Book, error)
	Update(book *entities.Book) error
	Delete(id string) error
}
