package usecase

import (
	"errors"
	"net/url"
	"strings"

	"library-management-system/internal/domain/entities"
	"library-management-system/internal/domain/repositories"
)

// URLUseCase handles URL processing business logic
type URLUseCase struct {
	urlRepo repositories.URLRepository
}

// NewURLUseCase creates a new URL use case
func NewURLUseCase(urlRepo repositories.URLRepository) *URLUseCase {
	return &URLUseCase{
		urlRepo: urlRepo,
	}
}

// ProcessURL processes a URL according to the specified operation
func (uc *URLUseCase) ProcessURL(request *entities.URLRequest) (*entities.URLResponse, error) {
	// Validate input
	if request.URL == "" {
		return nil, errors.New("URL is required")
	}

	if request.Operation == "" {
		return nil, errors.New("operation is required")
	}

	// Validate operation type
	validOperations := []string{"canonical", "redirection", "all"}
	isValidOperation := false
	for _, op := range validOperations {
		if request.Operation == op {
			isValidOperation = true
			break
		}
	}
	if !isValidOperation {
		return nil, errors.New("invalid operation type")
	}

	// Parse the URL
	parsedURL, err := url.Parse(request.URL)
	if err != nil {
		return nil, errors.New("invalid URL format")
	}

	var processedURL string

	switch request.Operation {
	case "canonical":
		processedURL = uc.processCanonical(parsedURL)
	case "redirection":
		processedURL = uc.processRedirection(parsedURL)
	case "all":
		processedURL = uc.processAll(parsedURL)
	}

	return &entities.URLResponse{
		ProcessedURL: processedURL,
	}, nil
}

// processCanonical removes query parameters and trailing slashes
func (uc *URLUseCase) processCanonical(parsedURL *url.URL) string {
	// Remove query parameters
	parsedURL.RawQuery = ""

	// Remove trailing slashes from path
	path := strings.TrimRight(parsedURL.Path, "/")
	if path == "" {
		path = "/"
	}
	parsedURL.Path = path

	return parsedURL.String()
}

// processRedirection ensures domain is www.byfood.com and converts to lowercase
func (uc *URLUseCase) processRedirection(parsedURL *url.URL) string {
	// Set domain to www.byfood.com
	parsedURL.Host = "www.byfood.com"

	// Convert entire URL to lowercase
	return strings.ToLower(parsedURL.String())
}

// processAll applies both canonical and redirection processing
func (uc *URLUseCase) processAll(parsedURL *url.URL) string {
	// First apply canonical processing
	canonicalURL := uc.processCanonical(parsedURL)

	// Parse the canonical URL for redirection processing
	canonicalParsedURL, _ := url.Parse(canonicalURL)

	// Then apply redirection processing
	return uc.processRedirection(canonicalParsedURL)
}
