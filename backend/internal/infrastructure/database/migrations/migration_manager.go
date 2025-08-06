package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// MigrationManager handles database migrations using gormigrate
type MigrationManager struct {
	migrator   *gormigrate.Gormigrate
	db         *gorm.DB
	migrations []*gormigrate.Migration
}

// NewMigrationManager creates a new migration manager
func NewMigrationManager(db *gorm.DB) *MigrationManager {
	migrations := []*gormigrate.Migration{
		CreateBooksTable(),
		AddIndexesToBooks(),
		AddSoftDeleteToBooks(),
	}

	migrator := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	return &MigrationManager{
		migrator:   migrator,
		db:         db,
		migrations: migrations,
	}
}

// Migrate runs all pending migrations
func (m *MigrationManager) Migrate() error {
	log.Println("🔄 Starting database migrations...")

	if err := m.migrator.Migrate(); err != nil {
		log.Printf("❌ Migration failed: %v", err)
		return err
	}

	log.Println("✅ All migrations completed successfully")
	return nil
}

// Rollback rolls back the last migration
func (m *MigrationManager) Rollback() error {
	log.Println("⏪ Rolling back last migration...")

	if err := m.migrator.RollbackLast(); err != nil {
		log.Printf("❌ Rollback failed: %v", err)
		return err
	}

	log.Println("✅ Migration rolled back successfully")
	return nil
}

// RollbackTo rolls back to a specific migration
func (m *MigrationManager) RollbackTo(migrationID string) error {
	log.Printf("⏪ Rolling back to migration: %s", migrationID)

	if err := m.migrator.RollbackTo(migrationID); err != nil {
		log.Printf("❌ Rollback failed: %v", err)
		return err
	}

	log.Printf("✅ Rolled back to migration: %s", migrationID)
	return nil
}

// Status shows migration status
func (m *MigrationManager) Status() error {
	log.Println("📊 Migration Status:")

	// Get applied migrations from database
	var appliedMigrations []struct {
		ID string `gorm:"column:id"`
	}

	if err := m.db.Table("schema_migrations").Find(&appliedMigrations).Error; err != nil {
		// If table doesn't exist, no migrations have been applied
		log.Println("  No migrations have been applied yet")
		return nil
	}

	// Create a map of applied migrations for quick lookup
	appliedMap := make(map[string]bool)
	for _, migration := range appliedMigrations {
		appliedMap[migration.ID] = true
	}

	// Show status for each migration
	for _, migration := range m.migrations {
		status := "❌ Pending"
		if appliedMap[migration.ID] {
			status = "✅ Applied"
		}
		log.Printf("  %s - %s", migration.ID, status)
	}

	return nil
}

// GetMigrations returns all available migrations
func (m *MigrationManager) GetMigrations() []*gormigrate.Migration {
	return m.migrations
}

// GetAppliedMigrations returns all applied migrations
func (m *MigrationManager) GetAppliedMigrations() ([]string, error) {
	var appliedMigrations []struct {
		ID string `gorm:"column:id"`
	}

	if err := m.db.Table("schema_migrations").Find(&appliedMigrations).Error; err != nil {
		return nil, err
	}

	ids := make([]string, len(appliedMigrations))
	for i, migration := range appliedMigrations {
		ids[i] = migration.ID
	}

	return ids, nil
}
