package database

import (
	"fmt"
	"log"

	"library-management-system/internal/infrastructure/config"
	"library-management-system/internal/infrastructure/database/migrations"

	"gorm.io/driver/postgres"
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
	// Load configuration
	cfg := config.Load()

	// Configure GORM logger based on log level
	var gormLogLevel logger.LogLevel
	switch cfg.Logging.Level {
	case "debug":
		gormLogLevel = logger.Info
	case "info":
		gormLogLevel = logger.Info
	case "warn":
		gormLogLevel = logger.Warn
	case "error":
		gormLogLevel = logger.Error
	default:
		gormLogLevel = logger.Info
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	}

	var db *gorm.DB
	var err error

	// Connect to database based on type
	switch cfg.Database.Type {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.Database.Path), gormConfig)
		if err != nil {
			log.Printf("Failed to connect to SQLite database: %v", err)
			return nil, err
		}
		log.Printf("Connected to SQLite database: %s", cfg.Database.Path)
	case "postgres":
		// Build PostgreSQL connection string
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Name,
			cfg.Database.SSLMode,
		)

		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
		if err != nil {
			log.Printf("Failed to connect to PostgreSQL database: %v", err)
			return nil, err
		}
		log.Printf("Connected to PostgreSQL database: %s:%s/%s",
			cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Database.Type)
	}

	// Run migrations
	if err := runMigrations(db); err != nil {
		log.Printf("Failed to run migrations: %v", err)
		return nil, err
	}

	log.Println("Database connected and migrated successfully")
	return &Database{DB: db}, nil
}

// runMigrations runs database migrations using gormigrate
func runMigrations(db *gorm.DB) error {
	migrationManager := migrations.NewMigrationManager(db)
	return migrationManager.Migrate()
}

// GetDB returns the database instance
func (d *Database) GetDB() *gorm.DB {
	return d.DB
}

// RunMigrations runs migrations manually (for CLI commands)
func (d *Database) RunMigrations() error {
	migrationManager := migrations.NewMigrationManager(d.DB)
	return migrationManager.Migrate()
}

// RollbackMigration rolls back the last migration
func (d *Database) RollbackMigration() error {
	migrationManager := migrations.NewMigrationManager(d.DB)
	return migrationManager.Rollback()
}

// RollbackToMigration rolls back to a specific migration
func (d *Database) RollbackToMigration(migrationID string) error {
	migrationManager := migrations.NewMigrationManager(d.DB)
	return migrationManager.RollbackTo(migrationID)
}

// MigrationStatus shows migration status
func (d *Database) MigrationStatus() error {
	migrationManager := migrations.NewMigrationManager(d.DB)
	return migrationManager.Status()
}

// GetAppliedMigrations returns all applied migrations
func (d *Database) GetAppliedMigrations() ([]string, error) {
	migrationManager := migrations.NewMigrationManager(d.DB)
	return migrationManager.GetAppliedMigrations()
}
