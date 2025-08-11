package handlers

import (
	"net/http"

	"library-management-system/internal/domain/entities"
	"library-management-system/internal/usecase"

	"github.com/gin-gonic/gin"
)

// BookHandler handles HTTP requests for books
type BookHandler struct {
	bookUseCase *usecase.BookUseCase
}

// NewBookHandler creates a new book handler
func NewBookHandler(bookUseCase *usecase.BookUseCase) *BookHandler {
	return &BookHandler{
		bookUseCase: bookUseCase,
	}
}

// CreateBookRequest represents the request body for creating a book
type CreateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required"`
	ISBN   string `json:"isbn" binding:"required"`
}

// UpdateBookRequest represents the request body for updating a book
type UpdateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required"`
	ISBN   string `json:"isbn" binding:"required"`
}

// GetBooks handles GET /api/books
// @Summary Get all books
// @Description Retrieve all books from the library
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} entities.Book
// @Router /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookUseCase.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// CreateBook handles POST /api/books
// @Summary Create a new book
// @Description Create a new book in the library
// @Tags books
// @Accept json
// @Produce json
// @Param book body CreateBookRequest true "Book information"
// @Success 201 {object} entities.Book
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
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
// @Summary Get a book by ID
// @Description Retrieve a specific book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} entities.Book
// @Failure 404 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book ID is required"})
		return
	}

	book, err := h.bookUseCase.GetBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook handles PUT /api/books/:id
// @Summary Update a book
// @Description Update an existing book in the library
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body UpdateBookRequest true "Updated book information"
// @Success 200 {object} entities.Book
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 404 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book ID is required"})
		return
	}

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

	// Get the updated book to return with proper timestamps
	updatedBook, err := h.bookUseCase.GetBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated book"})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

// DeleteBook handles DELETE /api/books/:id
// @Summary Delete a book
// @Description Soft delete a book from the library
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} handlers.MessageResponse
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 404 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book ID is required"})
		return
	}

	if err := h.bookUseCase.DeleteBook(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}

// SearchBooks handles GET /api/books/search
// @Summary Search books
// @Description Search books by title, author, or year
// @Tags books
// @Accept json
// @Produce json
// @Param title query string false "Search by title"
// @Param author query string false "Search by author"
// @Param year query int false "Search by year"
// @Success 200 {array} entities.Book
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /books/search [get]
func (h *BookHandler) SearchBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	yearStr := c.Query("year")

	var books []entities.Book
	var err error

	switch {
	case title != "":
		books, err = h.bookUseCase.SearchBooksByTitle(title)
	case author != "":
		books, err = h.bookUseCase.SearchBooksByAuthor(author)
	case yearStr != "":
		books, err = h.bookUseCase.SearchBooksByYear(yearStr)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "at least one search parameter is required"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetDeletedBooks handles GET /api/books/deleted
// @Summary Get deleted books
// @Description Retrieve all soft-deleted books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} entities.Book
// @Failure 500 {object} handlers.ErrorResponse
// @Router /books/deleted [get]
func (h *BookHandler) GetDeletedBooks(c *gin.Context) {
	books, err := h.bookUseCase.GetDeletedBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// RestoreBook handles POST /api/books/:id/restore
// @Summary Restore a deleted book
// @Description Restore a soft-deleted book
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} handlers.MessageResponse
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /books/{id}/restore [post]
func (h *BookHandler) RestoreBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book ID is required"})
		return
	}

	if err := h.bookUseCase.RestoreBook(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book restored successfully"})
}

// HardDeleteBook handles DELETE /api/books/:id/permanent
// @Summary Permanently delete a book
// @Description Permanently delete a book from the library
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} handlers.MessageResponse
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 404 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /books/{id}/permanent [delete]
func (h *BookHandler) HardDeleteBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book ID is required"})
		return
	}

	if err := h.bookUseCase.HardDeleteBook(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book permanently deleted"})
}
