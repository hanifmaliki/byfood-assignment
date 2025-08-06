package usecase

import (
	"testing"

	"library-management-system/internal/domain/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBookRepository is a mock implementation of BookRepository
type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) Create(book *entities.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) GetByID(id string) (*entities.Book, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Book), args.Error(1)
}

func (m *MockBookRepository) GetAll() ([]entities.Book, error) {
	args := m.Called()
	return args.Get(0).([]entities.Book), args.Error(1)
}

func (m *MockBookRepository) Update(book *entities.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookRepository) HardDelete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookRepository) FindByTitle(title string) ([]entities.Book, error) {
	args := m.Called(title)
	return args.Get(0).([]entities.Book), args.Error(1)
}

func (m *MockBookRepository) FindByAuthor(author string) ([]entities.Book, error) {
	args := m.Called(author)
	return args.Get(0).([]entities.Book), args.Error(1)
}

func (m *MockBookRepository) FindByYear(year int) ([]entities.Book, error) {
	args := m.Called(year)
	return args.Get(0).([]entities.Book), args.Error(1)
}

func (m *MockBookRepository) FindByISBN(isbn string) (*entities.Book, error) {
	args := m.Called(isbn)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Book), args.Error(1)
}

func (m *MockBookRepository) GetDeletedBooks() ([]entities.Book, error) {
	args := m.Called()
	return args.Get(0).([]entities.Book), args.Error(1)
}

func (m *MockBookRepository) Restore(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestNewBookUseCase(t *testing.T) {
	mockRepo := &MockBookRepository{}
	useCase := NewBookUseCase(mockRepo)

	assert.NotNil(t, useCase)
	assert.Equal(t, mockRepo, useCase.bookRepo)
}

func TestBookUseCase_CreateBook(t *testing.T) {
	tests := []struct {
		name          string
		book          *entities.Book
		mockSetup     func(*MockBookRepository)
		expectedError string
	}{
		{
			name: "successful creation",
			book: &entities.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			mockSetup: func(repo *MockBookRepository) {
				repo.On("FindByISBN", "1234567890").Return((*entities.Book)(nil), nil)
				repo.On("Create", mock.AnythingOfType("*entities.Book")).Return(nil)
			},
			expectedError: "",
		},
		{
			name: "ISBN already exists",
			book: &entities.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			mockSetup: func(repo *MockBookRepository) {
				existingBook := &entities.Book{ID: "existing-id", ISBN: "1234567890"}
				repo.On("FindByISBN", "1234567890").Return(existingBook, nil)
			},
			expectedError: "book with this ISBN already exists",
		},
		{
			name: "invalid book data",
			book: &entities.Book{
				Title:  "", // Empty title
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			mockSetup: func(repo *MockBookRepository) {
				// No mock setup needed as validation should fail first
			},
			expectedError: "book title is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockBookRepository{}
			useCase := NewBookUseCase(mockRepo)

			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			err := useCase.CreateBook(tt.book)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBookUseCase_GetBook(t *testing.T) {
	tests := []struct {
		name          string
		id            string
		mockSetup     func(*MockBookRepository)
		expectedBook  *entities.Book
		expectedError string
	}{
		{
			name: "successful retrieval",
			id:   "test-id",
			mockSetup: func(repo *MockBookRepository) {
				book := &entities.Book{
					ID:     "test-id",
					Title:  "Test Book",
					Author: "Test Author",
					Year:   2024,
					ISBN:   "1234567890",
				}
				repo.On("GetByID", "test-id").Return(book, nil)
			},
			expectedBook: &entities.Book{
				ID:     "test-id",
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			expectedError: "",
		},
		{
			name: "empty ID",
			id:   "",
			mockSetup: func(repo *MockBookRepository) {
				// No mock setup needed as validation should fail first
			},
			expectedBook:  nil,
			expectedError: "book ID is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockBookRepository{}
			useCase := NewBookUseCase(mockRepo)

			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			book, err := useCase.GetBook(tt.id)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, book)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBook, book)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBookUseCase_GetAllBooks(t *testing.T) {
	mockRepo := &MockBookRepository{}
	useCase := NewBookUseCase(mockRepo)

	expectedBooks := []entities.Book{
		{ID: "1", Title: "Book 1", Author: "Author 1", Year: 2024, ISBN: "1234567890"},
		{ID: "2", Title: "Book 2", Author: "Author 2", Year: 2023, ISBN: "0987654321"},
	}

	mockRepo.On("GetAll").Return(expectedBooks, nil)

	books, err := useCase.GetAllBooks()

	assert.NoError(t, err)
	assert.Equal(t, expectedBooks, books)
	mockRepo.AssertExpectations(t)
}

func TestBookUseCase_UpdateBook(t *testing.T) {
	tests := []struct {
		name          string
		id            string
		book          *entities.Book
		mockSetup     func(*MockBookRepository)
		expectedError string
	}{
		{
			name: "successful update",
			id:   "test-id",
			book: &entities.Book{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			mockSetup: func(repo *MockBookRepository) {
				existingBook := &entities.Book{
					ID:     "test-id",
					Title:  "Original Book",
					Author: "Original Author",
					Year:   2023,
					ISBN:   "1234567890",
				}
				repo.On("GetByID", "test-id").Return(existingBook, nil)
				repo.On("Update", mock.AnythingOfType("*entities.Book")).Return(nil)
			},
			expectedError: "",
		},
		{
			name: "book not found",
			id:   "non-existent-id",
			book: &entities.Book{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			mockSetup: func(repo *MockBookRepository) {
				repo.On("GetByID", "non-existent-id").Return((*entities.Book)(nil), nil)
			},
			expectedError: "book not found",
		},
		{
			name: "empty ID",
			id:   "",
			book: &entities.Book{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			mockSetup: func(repo *MockBookRepository) {
				// No mock setup needed as validation should fail first
			},
			expectedError: "book ID is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockBookRepository{}
			useCase := NewBookUseCase(mockRepo)

			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			err := useCase.UpdateBook(tt.id, tt.book)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBookUseCase_DeleteBook(t *testing.T) {
	tests := []struct {
		name          string
		id            string
		mockSetup     func(*MockBookRepository)
		expectedError string
	}{
		{
			name: "successful deletion",
			id:   "test-id",
			mockSetup: func(repo *MockBookRepository) {
				existingBook := &entities.Book{
					ID:     "test-id",
					Title:  "Test Book",
					Author: "Test Author",
					Year:   2024,
					ISBN:   "1234567890",
				}
				repo.On("GetByID", "test-id").Return(existingBook, nil)
				repo.On("Delete", "test-id").Return(nil)
			},
			expectedError: "",
		},
		{
			name: "book not found",
			id:   "non-existent-id",
			mockSetup: func(repo *MockBookRepository) {
				repo.On("GetByID", "non-existent-id").Return((*entities.Book)(nil), nil)
			},
			expectedError: "book not found",
		},
		{
			name: "empty ID",
			id:   "",
			mockSetup: func(repo *MockBookRepository) {
				// No mock setup needed as validation should fail first
			},
			expectedError: "book ID is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockBookRepository{}
			useCase := NewBookUseCase(mockRepo)

			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			err := useCase.DeleteBook(tt.id)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBookUseCase_SearchBooksByTitle(t *testing.T) {
	tests := []struct {
		name          string
		title         string
		mockSetup     func(*MockBookRepository)
		expectedBooks []entities.Book
		expectedError string
	}{
		{
			name:  "successful search",
			title: "Test",
			mockSetup: func(repo *MockBookRepository) {
				books := []entities.Book{
					{ID: "1", Title: "Test Book 1", Author: "Author 1", Year: 2024, ISBN: "1234567890"},
					{ID: "2", Title: "Test Book 2", Author: "Author 2", Year: 2023, ISBN: "0987654321"},
				}
				repo.On("FindByTitle", "Test").Return(books, nil)
			},
			expectedBooks: []entities.Book{
				{ID: "1", Title: "Test Book 1", Author: "Author 1", Year: 2024, ISBN: "1234567890"},
				{ID: "2", Title: "Test Book 2", Author: "Author 2", Year: 2023, ISBN: "0987654321"},
			},
			expectedError: "",
		},
		{
			name:  "empty title",
			title: "",
			mockSetup: func(repo *MockBookRepository) {
				// No mock setup needed as validation should fail first
			},
			expectedBooks: nil,
			expectedError: "title is required for search",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockBookRepository{}
			useCase := NewBookUseCase(mockRepo)

			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			books, err := useCase.SearchBooksByTitle(tt.title)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, books)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBooks, books)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBookUseCase_validateBook(t *testing.T) {
	useCase := &BookUseCase{}

	tests := []struct {
		name          string
		book          *entities.Book
		expectedError string
	}{
		{
			name: "valid book",
			book: &entities.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			expectedError: "",
		},
		{
			name: "empty title",
			book: &entities.Book{
				Author: "Test Author",
				Year:   2024,
				ISBN:   "1234567890",
			},
			expectedError: "book title is required",
		},
		{
			name: "empty author",
			book: &entities.Book{
				Title: "Test Book",
				Year:  2024,
				ISBN:  "1234567890",
			},
			expectedError: "book author is required",
		},
		{
			name: "year too old",
			book: &entities.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   999,
				ISBN:   "1234567890",
			},
			expectedError: "book year must be between 1000 and 2100",
		},
		{
			name: "year too new",
			book: &entities.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2101,
				ISBN:   "1234567890",
			},
			expectedError: "book year must be between 1000 and 2100",
		},
		{
			name: "empty ISBN",
			book: &entities.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
			},
			expectedError: "book ISBN is required",
		},
		{
			name: "ISBN too short",
			book: &entities.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "123",
			},
			expectedError: "book ISBN must be between 10 and 13 characters",
		},
		{
			name: "ISBN too long",
			book: &entities.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2024,
				ISBN:   "12345678901234",
			},
			expectedError: "book ISBN must be between 10 and 13 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := useCase.validateBook(tt.book)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
