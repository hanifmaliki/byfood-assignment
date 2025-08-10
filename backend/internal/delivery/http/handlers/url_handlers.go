package handlers

import (
	"net/http"

	"library-management-system/internal/domain/entities"
	"library-management-system/internal/usecase"

	"github.com/gin-gonic/gin"
)

// URLHandler handles HTTP requests for URL processing operations
type URLHandler struct {
	urlUseCase *usecase.URLUseCase
}

// NewURLHandler creates a new URL handler
func NewURLHandler(urlUseCase *usecase.URLUseCase) *URLHandler {
	return &URLHandler{
		urlUseCase: urlUseCase,
	}
}

// ProcessURL handles POST /api/url/process
// @Summary Process URL
// @Description Process a URL according to the specified operation (canonical, redirection, or all)
// @Tags url
// @Accept json
// @Produce json
// @Param request body entities.URLRequest true "URL processing request"
// @Success 200 {object} entities.URLResponse
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /url/process [post]
func (h *URLHandler) ProcessURL(c *gin.Context) {
	var req entities.URLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.urlUseCase.ProcessURL(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
