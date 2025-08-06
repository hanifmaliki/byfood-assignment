package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// AddIndexesToBooks adds performance indexes to the books table
func AddIndexesToBooks() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "002_add_indexes_to_books",
		Migrate: func(tx *gorm.DB) error {
			indexes := []string{
				"CREATE INDEX IF NOT EXISTS idx_books_title ON books(title)",
				"CREATE INDEX IF NOT EXISTS idx_books_author ON books(author)",
				"CREATE INDEX IF NOT EXISTS idx_books_year ON books(year)",
				"CREATE INDEX IF NOT EXISTS idx_books_isbn ON books(isbn)",
				"CREATE INDEX IF NOT EXISTS idx_books_created_at ON books(created_at)",
			}

			for _, index := range indexes {
				if err := tx.Exec(index).Error; err != nil {
					return err
				}
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			indexes := []string{
				"DROP INDEX IF EXISTS idx_books_title",
				"DROP INDEX IF EXISTS idx_books_author",
				"DROP INDEX IF EXISTS idx_books_year",
				"DROP INDEX IF EXISTS idx_books_isbn",
				"DROP INDEX IF EXISTS idx_books_created_at",
			}

			for _, index := range indexes {
				if err := tx.Exec(index).Error; err != nil {
					return err
				}
			}

			return nil
		},
	}
}
