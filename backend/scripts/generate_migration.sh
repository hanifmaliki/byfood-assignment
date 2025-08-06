#!/bin/bash

# Migration Generator Script
# Usage: ./scripts/generate_migration.sh <migration_name>
# Example: ./scripts/generate_migration.sh add_user_table

if [ $# -eq 0 ]; then
    echo "❌ Migration name is required"
    echo "Usage: ./scripts/generate_migration.sh <migration_name>"
    echo "Example: ./scripts/generate_migration.sh add_user_table"
    exit 1
fi

MIGRATION_NAME=$1
TIMESTAMP=$(date +"%Y%m%d%H%M%S")
FILENAME="${TIMESTAMP}_${MIGRATION_NAME}.go"
FILEPATH="internal/infrastructure/database/migrations/${FILENAME}"

# Convert migration name to function name (snake_case to PascalCase)
FUNCTION_NAME=$(echo $MIGRATION_NAME | sed 's/_\([a-z]\)/\U\1/g' | sed 's/^[a-z]/\U&/')

echo "🔄 Generating migration: ${FILENAME}"
echo "📁 File path: ${FILEPATH}"
echo "🔧 Function name: ${FUNCTION_NAME}"

# Create the migration file
cat > "$FILEPATH" << EOF
package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// ${FUNCTION_NAME} ${MIGRATION_NAME}
func ${FUNCTION_NAME}() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "${TIMESTAMP}_${MIGRATION_NAME}",
		Migrate: func(tx *gorm.DB) error {
			// TODO: Implement migration logic
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// TODO: Implement rollback logic
			return nil
		},
	}
}
EOF

echo "✅ Migration file created: ${FILEPATH}"
echo ""
echo "📝 Next steps:"
echo "1. Edit ${FILEPATH} to implement your migration logic"
echo "2. Add the migration to migration_manager.go:"
echo "   migrations := []*gormigrate.Migration{"
echo "       CreateBooksTable(),"
echo "       AddIndexesToBooks(),"
echo "       AddSoftDeleteToBooks(),"
echo "       ${FUNCTION_NAME}(), // Add this line"
echo "   }"
echo ""
echo "3. Test your migration:"
echo "   make migrate"
echo ""
echo "4. If needed, rollback:"
echo "   make rollback" 