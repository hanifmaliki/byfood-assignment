package migrations

import (
	"fmt"
	"log"
	"time"

	"library-management-system/internal/domain/entities"

	"gorm.io/gorm"
)

// Migration represents a database migration
type Migration struct {
	ID        uint      `gorm:"primaryKey"`
	Version   string    `gorm:"uniqueIndex;not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

// MigrationFunc represents a migration function
type MigrationFunc func(*gorm.DB) error

// MigrationRecord represents a migration with its function
type MigrationRecord struct {
	Version string
	Name    string
	Up      MigrationFunc
	Down    MigrationFunc
}

// Migrator handles database migrations
type Migrator struct {
	db         *gorm.DB
	migrations []MigrationRecord
}

// NewMigrator creates a new migrator instance
func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{
		db:         db,
		migrations: getMigrations(),
	}
}

// getMigrations returns all available migrations
func getMigrations() []MigrationRecord {
	return []MigrationRecord{
		{
			Version: "001",
			Name:    "create_books_table",
			Up:      createBooksTable,
			Down:    dropBooksTable,
		},
		{
			Version: "002",
			Name:    "add_indexes_to_books",
			Up:      addIndexesToBooks,
			Down:    removeIndexesFromBooks,
		},
		{
			Version: "003",
			Name:    "add_soft_delete_to_books",
			Up:      addSoftDeleteToBooks,
			Down:    removeSoftDeleteFromBooks,
		},
	}
}

// Migrate runs all pending migrations
func (m *Migrator) Migrate() error {
	log.Println("Starting database migrations...")

	// Create migrations table if it doesn't exist
	if err := m.db.AutoMigrate(&Migration{}); err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Get applied migrations
	var appliedMigrations []Migration
	if err := m.db.Find(&appliedMigrations).Error; err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	// Create a map of applied migrations for quick lookup
	appliedMap := make(map[string]bool)
	for _, migration := range appliedMigrations {
		appliedMap[migration.Version] = true
	}

	// Run pending migrations
	for _, migration := range m.migrations {
		if !appliedMap[migration.Version] {
			log.Printf("Running migration %s: %s", migration.Version, migration.Name)

			if err := migration.Up(m.db); err != nil {
				return fmt.Errorf("failed to run migration %s: %w", migration.Version, err)
			}

			// Record the migration
			record := Migration{
				Version:   migration.Version,
				Name:      migration.Name,
				CreatedAt: time.Now(),
			}

			if err := m.db.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to record migration %s: %w", migration.Version, err)
			}

			log.Printf("✅ Migration %s completed successfully", migration.Version)
		} else {
			log.Printf("⏭️  Migration %s already applied", migration.Version)
		}
	}

	log.Println("🎉 All migrations completed successfully")
	return nil
}

// Rollback rolls back the last migration
func (m *Migrator) Rollback() error {
	log.Println("Rolling back last migration...")

	var lastMigration Migration
	if err := m.db.Order("version DESC").First(&lastMigration).Error; err != nil {
		return fmt.Errorf("no migrations to rollback: %w", err)
	}

	// Find the migration function
	var migrationFunc MigrationFunc
	for _, migration := range m.migrations {
		if migration.Version == lastMigration.Version {
			migrationFunc = migration.Down
			break
		}
	}

	if migrationFunc == nil {
		return fmt.Errorf("migration function not found for version %s", lastMigration.Version)
	}

	// Run the rollback
	if err := migrationFunc(m.db); err != nil {
		return fmt.Errorf("failed to rollback migration %s: %w", lastMigration.Version, err)
	}

	// Remove the migration record
	if err := m.db.Delete(&lastMigration).Error; err != nil {
		return fmt.Errorf("failed to remove migration record: %w", err)
	}

	log.Printf("✅ Rolled back migration %s", lastMigration.Version)
	return nil
}

// Status shows migration status
func (m *Migrator) Status() error {
	log.Println("📊 Migration Status:")

	var appliedMigrations []Migration
	if err := m.db.Order("version").Find(&appliedMigrations).Error; err != nil {
		return fmt.Errorf("failed to get migration status: %w", err)
	}

	appliedMap := make(map[string]bool)
	for _, migration := range appliedMigrations {
		appliedMap[migration.Version] = true
	}

	for _, migration := range m.migrations {
		status := "❌ Pending"
		if appliedMap[migration.Version] {
			status = "✅ Applied"
		}
		log.Printf("  %s - %s: %s", migration.Version, migration.Name, status)
	}

	return nil
}

// Migration functions

func createBooksTable(db *gorm.DB) error {
	return db.AutoMigrate(&entities.Book{})
}

func dropBooksTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&entities.Book{})
}

func addIndexesToBooks(db *gorm.DB) error {
	// Add indexes for better performance
	indexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_books_title ON books(title)",
		"CREATE INDEX IF NOT EXISTS idx_books_author ON books(author)",
		"CREATE INDEX IF NOT EXISTS idx_books_year ON books(year)",
		"CREATE INDEX IF NOT EXISTS idx_books_isbn ON books(isbn)",
		"CREATE INDEX IF NOT EXISTS idx_books_created_at ON books(created_at)",
	}

	for _, index := range indexes {
		if err := db.Exec(index).Error; err != nil {
			return fmt.Errorf("failed to create index: %w", err)
		}
	}

	return nil
}

func removeIndexesFromBooks(db *gorm.DB) error {
	indexes := []string{
		"DROP INDEX IF EXISTS idx_books_title",
		"DROP INDEX IF EXISTS idx_books_author",
		"DROP INDEX IF EXISTS idx_books_year",
		"DROP INDEX IF EXISTS idx_books_isbn",
		"DROP INDEX IF EXISTS idx_books_created_at",
	}

	for _, index := range indexes {
		if err := db.Exec(index).Error; err != nil {
			return fmt.Errorf("failed to drop index: %w", err)
		}
	}

	return nil
}

func addSoftDeleteToBooks(db *gorm.DB) error {
	// Add deleted_at column for soft deletes
	return db.Exec("ALTER TABLE books ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP").Error
}

func removeSoftDeleteFromBooks(db *gorm.DB) error {
	// Remove deleted_at column
	return db.Exec("ALTER TABLE books DROP COLUMN IF EXISTS deleted_at").Error
}
