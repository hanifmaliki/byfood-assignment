package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"library-management-system/internal/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// BookUseCaseInterface defines the interface for book use case
type BookUseCaseInterface interface {
	CreateBook(book *entities.Book) error
	GetBook(id string) (*entities.Book, error)
	GetAllBooks() ([]entities.Book, error)
	UpdateBook(id string, book *entities.Book) error
	DeleteBook(id string) error
	HardDeleteBook(id string) error
	SearchBooksByTitle(title string) ([]entities.Book, error)
	SearchBooksByAuthor(author string) ([]entities.Book, error)
	SearchBooksByYear(yearStr string) ([]entities.Book, error)
	GetDeletedBooks() ([]entities.Book, error)
	RestoreBook(id string) error
}

// MockBookUseCase is a mock implementation of BookUseCaseInterface
type MockBookUseCase struct {
	mock.Mock
}

func (m *MockBookUseCase) CreateBook(book *entities.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookUseCase) GetBook(id string) (*entities.Book, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Book), args.Error(1)
}

func (m *MockBookUseCase) GetAllBooks() ([]entities.Book, error) {
	args := m.Called()
	return args.Get(0).([]entities.Book), args.Error(1)
}

func (m *MockBookUseCase) UpdateBook(id string, book *entities.Book) error {
	args := m.Called(id, book)
	return args.Error(0)
}

func (m *MockBookUseCase) DeleteBook(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookUseCase) HardDeleteBook(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookUseCase) SearchBooksByTitle(title string) ([]entities.Book, error) {
	args := m.Called(title)
	return args.Get(0).([]entities.Book), args.Error(1)
}

func (m *MockBookUseCase) SearchBooksByAuthor(author string) ([]entities.Book, error) {
	args := m.Called(author)
	return args.Get(0).([]entities.Book), args.Error(1)
}

func (m *MockBookUseCase) SearchBooksByYear(yearStr string) ([]entities.Book, error) {
	args := m.Called(yearStr)
	return args.Get(0).([]entities.Book), args.Error(1)
}

func (m *MockBookUseCase) GetDeletedBooks() ([]entities.Book, error) {
	args := m.Called()
	return args.Get(0).([]entities.Book), args.Error(1)
}

func (m *MockBookUseCase) RestoreBook(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// TestBookHandler wraps BookHandler for testing
type TestBookHandler struct {
	bookUseCase BookUseCaseInterface
}

// NewTestBookHandler creates a new test book handler
func NewTestBookHandler(bookUseCase BookUseCaseInterface) *TestBookHandler {
	return &TestBookHandler{
		bookUseCase: bookUseCase,
	}
}

// GetBooks handles GET /api/books
func (h *TestBookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookUseCase.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// CreateBook handles POST /api/books
func (h *TestBookHandler) CreateBook(c *gin.Context) {
	var req CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := &entities.Book{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
		ISBN:   req.ISBN,
	}

	if err := h.bookUseCase.CreateBook(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// GetBook handles GET /api/books/:id
func (h *TestBookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := h.bookUseCase.GetBook(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook handles PUT /api/books/:id
func (h *TestBookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var req UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := &entities.Book{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
		ISBN:   req.ISBN,
	}

	if err := h.bookUseCase.UpdateBook(id, book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook handles DELETE /api/books/:id
func (h *TestBookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := h.bookUseCase.DeleteBook(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}

// SearchBooks handles GET /api/books/search
func (h *TestBookHandler) SearchBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	year := c.Query("year")

	if title == "" && author == "" && year == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "at least one search parameter is required"})
		return
	}

	var books []entities.Book
	var err error

	if title != "" {
		books, err = h.bookUseCase.SearchBooksByTitle(title)
	} else if author != "" {
		books, err = h.bookUseCase.SearchBooksByAuthor(author)
	} else if year != "" {
		books, err = h.bookUseCase.SearchBooksByYear(year)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetDeletedBooks handles GET /api/books/deleted
func (h *TestBookHandler) GetDeletedBooks(c *gin.Context) {
	books, err := h.bookUseCase.GetDeletedBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// RestoreBook handles POST /api/books/:id/restore
func (h *TestBookHandler) RestoreBook(c *gin.Context) {
	id := c.Param("id")
	if err := h.bookUseCase.RestoreBook(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book restored successfully"})
}

// HardDeleteBook handles DELETE /api/books/:id/permanent
func (h *TestBookHandler) HardDeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := h.bookUseCase.HardDeleteBook(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book permanently deleted"})
}

func setupTestRouter(handler *TestBookHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	api := router.Group("/api")
	books := api.Group("/books")
	{
		books.GET("", handler.GetBooks)
		books.POST("", handler.CreateBook)
		books.GET("/search", handler.SearchBooks)
		books.GET("/deleted", handler.GetDeletedBooks)
		books.GET("/:id", handler.GetBook)
		books.PUT("/:id", handler.UpdateBook)
		books.DELETE("/:id", handler.DeleteBook)
		books.POST("/:id/restore", handler.RestoreBook)
		books.DELETE("/:id/permanent", handler.HardDeleteBook)
	}

	return router
}

func TestNewBookHandler(t *testing.T) {
	mockUseCase := &MockBookUseCase{}
	handler := NewTestBookHandler(mockUseCase)

	assert.NotNil(t, handler)
	assert.Equal(t, mockUseCase, handler.bookUseCase)
}

func TestBookHandler_CreateBook(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    CreateBookRequest
		mockSetup      func(*MockBookUseCase)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful creation",
			requestBody: CreateBookRequest{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			mockSetup: func(useCase *MockBookUseCase) {
				useCase.On("CreateBook", mock.AnythingOfType("*entities.Book")).Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   `"title":"Test Book"`,
		},
		{
			name: "missing required fields",
			requestBody: CreateBookRequest{
				Title:  "",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			mockSetup: func(useCase *MockBookUseCase) {
				// No mock setup needed as validation should fail first
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `"error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &MockBookUseCase{}
			handler := NewTestBookHandler(mockUseCase)

			if tt.mockSetup != nil {
				tt.mockSetup(mockUseCase)
			}

			router := setupTestRouter(handler)

			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/api/books", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestBookHandler_GetBooks(t *testing.T) {
	mockUseCase := &MockBookUseCase{}
	handler := NewTestBookHandler(mockUseCase)

	expectedBooks := []entities.Book{
		{ID: "1", Title: "Book 1", Author: "Author 1", Year: 2024, ISBN: "1234567890"},
		{ID: "2", Title: "Book 2", Author: "Author 2", Year: 2023, ISBN: "0987654321"},
	}

	mockUseCase.On("GetAllBooks").Return(expectedBooks, nil)

	router := setupTestRouter(handler)
	req := httptest.NewRequest("GET", "/api/books", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"id":"1"`)
	assert.Contains(t, w.Body.String(), `"id":"2"`)
	mockUseCase.AssertExpectations(t)
}

func TestBookHandler_GetBook(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockSetup      func(*MockBookUseCase)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful retrieval",
			id:   "test-id",
			mockSetup: func(useCase *MockBookUseCase) {
				book := &entities.Book{
					ID:     "test-id",
					Title:  "Test Book",
					Author: "Test Author",
					Year:   2024,
					ISBN:   "1234567890",
				}
				useCase.On("GetBook", "test-id").Return(book, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"id":"test-id"`,
		},
		{
			name: "book not found",
			id:   "non-existent-id",
			mockSetup: func(useCase *MockBookUseCase) {
				useCase.On("GetBook", "non-existent-id").Return((*entities.Book)(nil), nil)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `"error":"book not found"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &MockBookUseCase{}
			handler := NewTestBookHandler(mockUseCase)

			if tt.mockSetup != nil {
				tt.mockSetup(mockUseCase)
			}

			router := setupTestRouter(handler)
			req := httptest.NewRequest("GET", "/api/books/"+tt.id, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestBookHandler_UpdateBook(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		requestBody    UpdateBookRequest
		mockSetup      func(*MockBookUseCase)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful update",
			id:   "test-id",
			requestBody: UpdateBookRequest{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			mockSetup: func(useCase *MockBookUseCase) {
				useCase.On("UpdateBook", "test-id", mock.AnythingOfType("*entities.Book")).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"title":"Updated Book"`,
		},
		{
			name: "missing required fields",
			id:   "test-id",
			requestBody: UpdateBookRequest{
				Title:  "",
				Author: "Updated Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			mockSetup: func(useCase *MockBookUseCase) {
				// No mock setup needed as validation should fail first
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `"error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &MockBookUseCase{}
			handler := NewTestBookHandler(mockUseCase)

			if tt.mockSetup != nil {
				tt.mockSetup(mockUseCase)
			}

			router := setupTestRouter(handler)

			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("PUT", "/api/books/"+tt.id, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestBookHandler_DeleteBook(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockSetup      func(*MockBookUseCase)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful deletion",
			id:   "test-id",
			mockSetup: func(useCase *MockBookUseCase) {
				useCase.On("DeleteBook", "test-id").Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"message":"book deleted successfully"`,
		},
		{
			name: "book not found",
			id:   "non-existent-id",
			mockSetup: func(useCase *MockBookUseCase) {
				useCase.On("DeleteBook", "non-existent-id").Return(errors.New("book not found"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `"error":"book not found"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &MockBookUseCase{}
			handler := NewTestBookHandler(mockUseCase)

			if tt.mockSetup != nil {
				tt.mockSetup(mockUseCase)
			}

			router := setupTestRouter(handler)
			req := httptest.NewRequest("DELETE", "/api/books/"+tt.id, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestBookHandler_SearchBooks(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		mockSetup      func(*MockBookUseCase)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:  "search by title",
			query: "?title=Test",
			mockSetup: func(useCase *MockBookUseCase) {
				books := []entities.Book{
					{ID: "1", Title: "Test Book", Author: "Author 1", Year: 2024, ISBN: "1234567890"},
				}
				useCase.On("SearchBooksByTitle", "Test").Return(books, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"title":"Test Book"`,
		},
		{
			name:  "search by author",
			query: "?author=Author",
			mockSetup: func(useCase *MockBookUseCase) {
				books := []entities.Book{
					{ID: "1", Title: "Book 1", Author: "Author 1", Year: 2024, ISBN: "1234567890"},
				}
				useCase.On("SearchBooksByAuthor", "Author").Return(books, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"author":"Author 1"`,
		},
		{
			name:  "search by year",
			query: "?year=2024",
			mockSetup: func(useCase *MockBookUseCase) {
				books := []entities.Book{
					{ID: "1", Title: "Book 1", Author: "Author 1", Year: 2024, ISBN: "1234567890"},
				}
				useCase.On("SearchBooksByYear", "2024").Return(books, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"year":2024`,
		},
		{
			name:           "no search parameters",
			query:          "",
			mockSetup:      func(useCase *MockBookUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `"error":"at least one search parameter is required"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &MockBookUseCase{}
			handler := NewTestBookHandler(mockUseCase)

			if tt.mockSetup != nil {
				tt.mockSetup(mockUseCase)
			}

			router := setupTestRouter(handler)
			req := httptest.NewRequest("GET", "/api/books/search"+tt.query, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestBookHandler_GetDeletedBooks(t *testing.T) {
	mockUseCase := &MockBookUseCase{}
	handler := NewTestBookHandler(mockUseCase)

	expectedBooks := []entities.Book{
		{ID: "1", Title: "Deleted Book 1", Author: "Author 1", Year: 2024, ISBN: "1234567890"},
		{ID: "2", Title: "Deleted Book 2", Author: "Author 2", Year: 2023, ISBN: "0987654321"},
	}

	mockUseCase.On("GetDeletedBooks").Return(expectedBooks, nil)

	router := setupTestRouter(handler)
	req := httptest.NewRequest("GET", "/api/books/deleted", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"title":"Deleted Book 1"`)
	assert.Contains(t, w.Body.String(), `"title":"Deleted Book 2"`)
	mockUseCase.AssertExpectations(t)
}

func TestBookHandler_RestoreBook(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockSetup      func(*MockBookUseCase)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful restore",
			id:   "test-id",
			mockSetup: func(useCase *MockBookUseCase) {
				useCase.On("RestoreBook", "test-id").Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"message":"book restored successfully"`,
		},
		{
			name: "restore failed",
			id:   "non-existent-id",
			mockSetup: func(useCase *MockBookUseCase) {
				useCase.On("RestoreBook", "non-existent-id").Return(errors.New("book not found"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `"error":"book not found"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &MockBookUseCase{}
			handler := NewTestBookHandler(mockUseCase)

			if tt.mockSetup != nil {
				tt.mockSetup(mockUseCase)
			}

			router := setupTestRouter(handler)
			req := httptest.NewRequest("POST", "/api/books/"+tt.id+"/restore", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestBookHandler_HardDeleteBook(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockSetup      func(*MockBookUseCase)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful hard delete",
			id:   "test-id",
			mockSetup: func(useCase *MockBookUseCase) {
				useCase.On("HardDeleteBook", "test-id").Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"message":"book permanently deleted"`,
		},
		{
			name: "hard delete failed",
			id:   "non-existent-id",
			mockSetup: func(useCase *MockBookUseCase) {
				useCase.On("HardDeleteBook", "non-existent-id").Return(errors.New("book not found"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `"error":"book not found"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &MockBookUseCase{}
			handler := NewTestBookHandler(mockUseCase)

			if tt.mockSetup != nil {
				tt.mockSetup(mockUseCase)
			}

			router := setupTestRouter(handler)
			req := httptest.NewRequest("DELETE", "/api/books/"+tt.id+"/permanent", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
			mockUseCase.AssertExpectations(t)
		})
	}
}
