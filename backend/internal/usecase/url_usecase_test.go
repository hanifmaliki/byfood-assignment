package usecase

import (
	"net/url"
	"testing"

	"library-management-system/internal/domain/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockURLRepository is a mock implementation of URLRepository
type MockURLRepository struct {
	mock.Mock
}

func (m *MockURLRepository) ProcessURL(request *entities.URLRequest) (*entities.URLResponse, error) {
	args := m.Called(request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.URLResponse), args.Error(1)
}

func TestNewURLUseCase(t *testing.T) {
	mockRepo := &MockURLRepository{}
	useCase := NewURLUseCase(mockRepo)

	assert.NotNil(t, useCase)
	assert.Equal(t, mockRepo, useCase.urlRepo)
}

func TestURLUseCase_ProcessURL(t *testing.T) {
	useCase := NewURLUseCase(&MockURLRepository{})

	tests := []struct {
		name           string
		request        *entities.URLRequest
		expectedResult *entities.URLResponse
		expectedError  string
	}{
		{
			name: "canonical operation - basic URL",
			request: &entities.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "canonical",
			},
			expectedResult: &entities.URLResponse{
				ProcessedURL: "https://BYFOOD.com/food-EXPeriences",
			},
			expectedError: "",
		},
		{
			name: "redirection operation - basic URL",
			request: &entities.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "redirection",
			},
			expectedResult: &entities.URLResponse{
				ProcessedURL: "https://www.byfood.com/food-experiences?query=abc/",
			},
			expectedError: "",
		},
		{
			name: "all operations - basic URL",
			request: &entities.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "all",
			},
			expectedResult: &entities.URLResponse{
				ProcessedURL: "https://www.byfood.com/food-experiences",
			},
			expectedError: "",
		},
		{
			name: "canonical operation - URL with multiple parameters",
			request: &entities.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc&sort=price&filter=available/",
				Operation: "canonical",
			},
			expectedResult: &entities.URLResponse{
				ProcessedURL: "https://BYFOOD.com/food-EXPeriences",
			},
			expectedError: "",
		},
		{
			name: "canonical operation - URL with special characters",
			request: &entities.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc%20def&sort=price%2Basc/",
				Operation: "canonical",
			},
			expectedResult: &entities.URLResponse{
				ProcessedURL: "https://BYFOOD.com/food-EXPeriences",
			},
			expectedError: "",
		},
		{
			name: "canonical operation - URL with hash",
			request: &entities.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/#section",
				Operation: "canonical",
			},
			expectedResult: &entities.URLResponse{
				ProcessedURL: "https://BYFOOD.com/food-EXPeriences#section",
			},
			expectedError: "",
		},
		{
			name: "canonical operation - URL with port",
			request: &entities.URLRequest{
				URL:       "https://BYFOOD.com:8080/food-EXPeriences?query=abc/",
				Operation: "canonical",
			},
			expectedResult: &entities.URLResponse{
				ProcessedURL: "https://BYFOOD.com:8080/food-EXPeriences",
			},
			expectedError: "",
		},
		{
			name: "canonical operation - URL with subdomain",
			request: &entities.URLRequest{
				URL:       "https://www.BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "canonical",
			},
			expectedResult: &entities.URLResponse{
				ProcessedURL: "https://www.BYFOOD.com/food-EXPeriences",
			},
			expectedError: "",
		},
		{
			name: "canonical operation - URL with path parameters",
			request: &entities.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences/123/details?query=abc/",
				Operation: "canonical",
			},
			expectedResult: &entities.URLResponse{
				ProcessedURL: "https://BYFOOD.com/food-EXPeriences/123/details",
			},
			expectedError: "",
		},
		{
			name: "invalid operation",
			request: &entities.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "invalid",
			},
			expectedResult: nil,
			expectedError:  "invalid operation type",
		},
		{
			name: "empty URL",
			request: &entities.URLRequest{
				URL:       "",
				Operation: "canonical",
			},
			expectedResult: nil,
			expectedError:  "URL is required",
		},
		{
			name: "empty operation",
			request: &entities.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "",
			},
			expectedResult: nil,
			expectedError:  "operation is required",
		},
		{
			name: "invalid URL format",
			request: &entities.URLRequest{
				URL:       "://invalid-url",
				Operation: "canonical",
			},
			expectedResult: nil,
			expectedError:  "invalid URL format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := useCase.ProcessURL(tt.request)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult.ProcessedURL, result.ProcessedURL)
			}
		})
	}
}

func TestURLUseCase_processCanonical(t *testing.T) {
	useCase := &URLUseCase{}

	tests := []struct {
		name           string
		url            string
		expectedResult string
	}{
		{
			name:           "basic URL with trailing slash",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc/",
			expectedResult: "https://BYFOOD.com/food-EXPeriences",
		},
		{
			name:           "URL without trailing slash",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc",
			expectedResult: "https://BYFOOD.com/food-EXPeriences",
		},
		{
			name:           "URL with multiple trailing slashes",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc///",
			expectedResult: "https://BYFOOD.com/food-EXPeriences",
		},
		{
			name:           "URL with mixed case in path",
			url:            "https://BYFOOD.com/Food-EXPeriences/Test-Path?query=abc/",
			expectedResult: "https://BYFOOD.com/Food-EXPeriences/Test-Path",
		},
		{
			name:           "URL with special characters in query",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc%20def&sort=price%2Basc/",
			expectedResult: "https://BYFOOD.com/food-EXPeriences",
		},
		{
			name:           "URL with hash fragment",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc/#section",
			expectedResult: "https://BYFOOD.com/food-EXPeriences#section",
		},
		{
			name:           "URL with port number",
			url:            "https://BYFOOD.com:8080/food-EXPeriences?query=abc/",
			expectedResult: "https://BYFOOD.com:8080/food-EXPeriences",
		},
		{
			name:           "URL with subdomain",
			url:            "https://www.BYFOOD.com/food-EXPeriences?query=abc/",
			expectedResult: "https://www.BYFOOD.com/food-EXPeriences",
		},
		{
			name:           "URL with path parameters",
			url:            "https://BYFOOD.com/food-EXPeriences/123/details?query=abc/",
			expectedResult: "https://BYFOOD.com/food-EXPeriences/123/details",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedURL, err := url.Parse(tt.url)
			assert.NoError(t, err)

			result := useCase.processCanonical(parsedURL)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestURLUseCase_processRedirection(t *testing.T) {
	useCase := &URLUseCase{}

	tests := []struct {
		name           string
		url            string
		expectedResult string
	}{
		{
			name:           "basic URL with trailing slash",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc/",
			expectedResult: "https://www.byfood.com/food-experiences?query=abc/",
		},
		{
			name:           "URL without trailing slash",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc",
			expectedResult: "https://www.byfood.com/food-experiences?query=abc",
		},
		{
			name:           "URL with multiple trailing slashes",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc///",
			expectedResult: "https://www.byfood.com/food-experiences?query=abc///",
		},
		{
			name:           "URL with mixed case in path",
			url:            "https://BYFOOD.com/Food-EXPeriences/Test-Path?query=abc/",
			expectedResult: "https://www.byfood.com/food-experiences/test-path?query=abc/",
		},
		{
			name:           "URL with special characters in query",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc%20def&sort=price%2Basc/",
			expectedResult: "https://www.byfood.com/food-experiences?query=abc%20def&sort=price%2basc/",
		},
		{
			name:           "URL with hash fragment",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc/#section",
			expectedResult: "https://www.byfood.com/food-experiences?query=abc/#section",
		},
		{
			name:           "URL with port number",
			url:            "https://BYFOOD.com:8080/food-EXPeriences?query=abc/",
			expectedResult: "https://www.byfood.com/food-experiences?query=abc/",
		},
		{
			name:           "URL with subdomain",
			url:            "https://www.BYFOOD.com/food-EXPeriences?query=abc/",
			expectedResult: "https://www.byfood.com/food-experiences?query=abc/",
		},
		{
			name:           "URL with path parameters",
			url:            "https://BYFOOD.com/food-EXPeriences/123/details?query=abc/",
			expectedResult: "https://www.byfood.com/food-experiences/123/details?query=abc/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedURL, err := url.Parse(tt.url)
			assert.NoError(t, err)

			result := useCase.processRedirection(parsedURL)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestURLUseCase_processAll(t *testing.T) {
	useCase := &URLUseCase{}

	tests := []struct {
		name           string
		url            string
		expectedResult string
	}{
		{
			name:           "basic URL with trailing slash",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc/",
			expectedResult: "https://www.byfood.com/food-experiences",
		},
		{
			name:           "URL without trailing slash",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc",
			expectedResult: "https://www.byfood.com/food-experiences",
		},
		{
			name:           "URL with multiple trailing slashes",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc///",
			expectedResult: "https://www.byfood.com/food-experiences",
		},
		{
			name:           "URL with mixed case in path",
			url:            "https://BYFOOD.com/Food-EXPeriences/Test-Path?query=abc/",
			expectedResult: "https://www.byfood.com/food-experiences/test-path",
		},
		{
			name:           "URL with special characters in query",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc%20def&sort=price%2Basc/",
			expectedResult: "https://www.byfood.com/food-experiences",
		},
		{
			name:           "URL with hash fragment",
			url:            "https://BYFOOD.com/food-EXPeriences?query=abc/#section",
			expectedResult: "https://www.byfood.com/food-experiences#section",
		},
		{
			name:           "URL with port number",
			url:            "https://BYFOOD.com:8080/food-EXPeriences?query=abc/",
			expectedResult: "https://www.byfood.com/food-experiences",
		},
		{
			name:           "URL with subdomain",
			url:            "https://www.BYFOOD.com/food-EXPeriences?query=abc/",
			expectedResult: "https://www.byfood.com/food-experiences",
		},
		{
			name:           "URL with path parameters",
			url:            "https://BYFOOD.com/food-EXPeriences/123/details?query=abc/",
			expectedResult: "https://www.byfood.com/food-experiences/123/details",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedURL, err := url.Parse(tt.url)
			assert.NoError(t, err)

			result := useCase.processAll(parsedURL)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}
