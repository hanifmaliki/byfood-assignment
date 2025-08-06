package main

import (
	"log"
	"net/http"

	"library-management-system/internal/delivery/http/handlers"
	"library-management-system/internal/infrastructure/database"
	"library-management-system/internal/repository"
	"library-management-system/internal/usecase"

	_ "library-management-system/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Library Management System API
// @version 1.0
// @description A RESTful API for managing books and URL processing with clean architecture
// @host localhost:8080
// @BasePath /api
func main() {
	// Initialize database
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repositories
	bookRepo := repository.NewBookRepository(db.GetDB())
	urlRepo := repository.NewURLRepository()

	// Initialize use cases
	bookUseCase := usecase.NewBookUseCase(bookRepo)
	urlUseCase := usecase.NewURLUseCase(urlRepo)

	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookUseCase)
	urlHandler := handlers.NewURLHandler(urlUseCase)

	// Initialize Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// API routes
	api := r.Group("/api")
	{
		// Book management routes
		books := api.Group("/books")
		{
			books.GET("", bookHandler.GetBooks)
			books.POST("", bookHandler.CreateBook)
			books.GET("/:id", bookHandler.GetBook)
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)
		}

		// URL processing routes
		url := api.Group("/url")
		{
			url.POST("/process", urlHandler.ProcessURL)
		}
	}

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
