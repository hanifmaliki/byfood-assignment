package database

import (
	"log"

	"library-management-system/internal/domain/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database represents the database connection
type Database struct {
	DB *gorm.DB
}

// NewDatabase creates a new database connection
func NewDatabase() (*Database, error) {
	// Configure GORM logger
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("library.db"), config)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Auto migrate the database
	err = db.AutoMigrate(&entities.Book{})
	if err != nil {
		log.Printf("Failed to migrate database: %v", err)
		return nil, err
	}

	log.Println("Database connected and migrated successfully")
	return &Database{DB: db}, nil
}

// GetDB returns the database instance
func (d *Database) GetDB() *gorm.DB {
	return d.DB
}
