package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"library-management-system/internal/delivery/http/handlers"
	"library-management-system/internal/infrastructure/config"
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
	// Load configuration
	cfg := config.Load()

	// Initialize application with configuration
	app := NewApplication(cfg)

	// Start the server
	if err := app.Start(); err != nil {
		log.Fatal("Failed to start application:", err)
	}
}

// Application represents the main application
type Application struct {
	config *config.Config
	router *gin.Engine
}

// NewApplication creates a new application instance
func NewApplication(cfg *config.Config) *Application {
	// Set Gin mode based on environment
	if cfg.Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Log configuration
	log.Printf("Loading configuration for environment: %s", cfg.Server.Environment)
	log.Printf("Database type: %s", cfg.Database.Type)
	log.Printf("Database path: %s", cfg.Database.Path)
	log.Printf("API Prefix: %s", cfg.API.Prefix)
	log.Printf("Swagger enabled: %t", cfg.Swagger.Enabled)

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

	// Initialize router
	router := gin.Default()

	// Add CORS middleware
	router.Use(corsMiddleware(cfg.CORS))

	// Setup routes
	setupRoutes(router, cfg, bookHandler, urlHandler)

	return &Application{
		config: cfg,
		router: router,
	}
}

// Start starts the application server
func (app *Application) Start() error {
	serverAddr := fmt.Sprintf("%s:%s", app.config.Server.Host, app.config.Server.Port)
	log.Printf("Server starting on %s", serverAddr)
	return app.router.Run(serverAddr)
}

// corsMiddleware creates CORS middleware
func corsMiddleware(cors config.CORSConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" && contains(cors.AllowedOrigins, origin) {
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			c.Header("Access-Control-Allow-Origin", strings.Join(cors.AllowedOrigins, ","))
		}

		c.Header("Access-Control-Allow-Methods", strings.Join(cors.AllowedMethods, ","))
		c.Header("Access-Control-Allow-Headers", strings.Join(cors.AllowedHeaders, ","))

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

// setupRoutes sets up all application routes
func setupRoutes(router *gin.Engine, cfg *config.Config, bookHandler *handlers.BookHandler, urlHandler *handlers.URLHandler) {
	// API routes
	api := router.Group(cfg.API.Prefix)
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
	if cfg.Swagger.Enabled {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Health check
	healthEndpoint := "/health"
	router.GET(healthEndpoint, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": cfg.Swagger.Title,
			"version": cfg.Swagger.Version,
		})
	})
}

// contains checks if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
