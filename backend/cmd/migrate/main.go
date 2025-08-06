package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"library-management-system/internal/infrastructure/database"
)

func main() {
	// Define command line flags
	var (
		command     = flag.String("command", "migrate", "Migration command: migrate, rollback, rollback-to, status")
		migrationID = flag.String("id", "", "Migration ID for rollback-to command")
		help        = flag.Bool("help", false, "Show help")
	)
	flag.Parse()

	if *help {
		showHelp()
		return
	}

	// Initialize database
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Execute command
	switch *command {
	case "migrate":
		if err := db.RunMigrations(); err != nil {
			log.Fatal("Failed to run migrations:", err)
		}
		fmt.Println("✅ Migrations completed successfully")

	case "rollback":
		if err := db.RollbackMigration(); err != nil {
			log.Fatal("Failed to rollback migration:", err)
		}
		fmt.Println("✅ Migration rolled back successfully")

	case "rollback-to":
		if *migrationID == "" {
			fmt.Println("❌ Migration ID is required for rollback-to command")
			fmt.Println("Usage: go run cmd/migrate/main.go -command=rollback-to -id=MIGRATION_ID")
			os.Exit(1)
		}
		if err := db.RollbackToMigration(*migrationID); err != nil {
			log.Fatal("Failed to rollback to migration:", err)
		}
		fmt.Printf("✅ Rolled back to migration: %s\n", *migrationID)

	case "status":
		if err := db.MigrationStatus(); err != nil {
			log.Fatal("Failed to get migration status:", err)
		}

	case "applied":
		applied, err := db.GetAppliedMigrations()
		if err != nil {
			log.Fatal("Failed to get applied migrations:", err)
		}
		fmt.Println("📋 Applied migrations:")
		for _, migrationID := range applied {
			fmt.Printf("  ✅ %s\n", migrationID)
		}

	default:
		fmt.Printf("❌ Unknown command: %s\n", *command)
		showHelp()
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Println("🔄 Database Migration Tool (gormigrate)")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run cmd/migrate/main.go [command] [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  migrate      Run pending migrations (default)")
	fmt.Println("  rollback     Rollback the last migration")
	fmt.Println("  rollback-to  Rollback to a specific migration")
	fmt.Println("  status       Show migration status")
	fmt.Println("  applied      Show applied migrations")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -command     Migration command")
	fmt.Println("  -id          Migration ID (for rollback-to)")
	fmt.Println("  -help        Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run cmd/migrate/main.go migrate")
	fmt.Println("  go run cmd/migrate/main.go rollback")
	fmt.Println("  go run cmd/migrate/main.go rollback-to -id=20241201000001_add_indexes_to_books")
	fmt.Println("  go run cmd/migrate/main.go status")
	fmt.Println("  go run cmd/migrate/main.go applied")
	fmt.Println()
	fmt.Println("Available Migrations:")
	fmt.Println("  20241201000000_create_books_table")
	fmt.Println("  20241201000001_add_indexes_to_books")
	fmt.Println("  20241201000002_add_soft_delete_to_books")
	fmt.Println()
	fmt.Println("📝 Migration Naming Convention:")
	fmt.Println("  Format: YYYYMMDDHHMMSS_descriptive_name")
	fmt.Println("  Example: 20241201000000_create_books_table")
	fmt.Println("  This ensures chronological order and prevents conflicts")
}
