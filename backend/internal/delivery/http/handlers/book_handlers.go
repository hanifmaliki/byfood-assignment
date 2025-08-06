package handlers

import (
	"net/http"

	"library-management-system/internal/usecase"

	"github.com/gin-gonic/gin"
)

// BookHandler handles HTTP requests for book operations
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
	Year   int    `json:"year" binding:"required,min=1800,max=2024"`
	ISBN   string `json:"isbn" binding:"required"`
}

// UpdateBookRequest represents the request body for updating a book
type UpdateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	ISBN   string `json:"isbn"`
}

// GetBooks handles GET /api/books
// @Summary Get all books
// @Description Retrieve all books from the library
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} entities.Book
// @Router /api/books [get]
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
// @Description Add a new book to the library
// @Tags books
// @Accept json
// @Produce json
// @Param book body CreateBookRequest true "Book information"
// @Success 201 {object} entities.Book
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookUseCase.CreateBook(req.Title, req.Author, req.ISBN, req.Year)
	if err != nil {
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
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book ID is required"})
		return
	}

	book, err := h.bookUseCase.GetBook(id)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook handles PUT /api/books/:id
// @Summary Update a book
// @Description Update an existing book's information
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body UpdateBookRequest true "Updated book information"
// @Success 200 {object} entities.Book
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/books/{id} [put]
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

	book, err := h.bookUseCase.UpdateBook(id, req.Title, req.Author, req.ISBN, req.Year)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook handles DELETE /api/books/:id
// @Summary Delete a book
// @Description Remove a book from the library
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 204 "No Content"
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book ID is required"})
		return
	}

	err := h.bookUseCase.DeleteBook(id)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
