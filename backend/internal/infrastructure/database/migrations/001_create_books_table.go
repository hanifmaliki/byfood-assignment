package migrations

import (
	"library-management-system/internal/domain/entities"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// CreateBooksTable creates the books table
func CreateBooksTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "001_create_books_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&entities.Book{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&entities.Book{})
		},
	}
}
