package repositories

import "library-management-system/internal/domain/entities"

// URLRepository defines the interface for URL processing
type URLRepository interface {
	ProcessURL(request *entities.URLRequest) (*entities.URLResponse, error)
}
