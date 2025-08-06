package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// AddSoftDeleteToBooks adds soft delete functionality to the books table
func AddSoftDeleteToBooks() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "003_add_soft_delete_to_books",
		Migrate: func(tx *gorm.DB) error {
			// Add deleted_at column for soft deletes
			return tx.Exec("ALTER TABLE books ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP").Error
		},
		Rollback: func(tx *gorm.DB) error {
			// Remove deleted_at column
			return tx.Exec("ALTER TABLE books DROP COLUMN IF EXISTS deleted_at").Error
		},
	}
}
