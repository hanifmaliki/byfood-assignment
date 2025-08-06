# Library Management System Makefile

.PHONY: help install setup test build run clean migrate rollback rollback-to status applied docker-up docker-down

# Default target
help:
	@echo "📚 Library Management System - Available Commands:"
	@echo ""
	@echo "🔧 Setup Commands:"
	@echo "  install     Install all dependencies (Go + Node.js)"
	@echo "  setup       Setup PostgreSQL database"
	@echo ""
	@echo "🚀 Development Commands:"
	@echo "  run         Run the full application (backend + frontend)"
	@echo "  backend     Run backend server only"
	@echo "  frontend    Run frontend server only"
	@echo ""
	@echo "🗄️  Database Commands:"
	@echo "  migrate     Run database migrations"
	@echo "  rollback    Rollback last migration"
	@echo "  rollback-to Rollback to specific migration (e.g., make rollback-to id=002_add_indexes_to_books)"
	@echo "  status      Show migration status"
	@echo "  applied     Show applied migrations"
	@echo "  db-reset    Reset database (rollback all + migrate)"
	@echo ""
	@echo "🐳 Docker Commands:"
	@echo "  docker-up   Start all services with Docker Compose"
	@echo "  docker-down Stop all Docker services"
	@echo ""
	@echo "🧪 Testing Commands:"
	@echo "  test        Run all tests"
	@echo "  test-backend Run backend tests only"
	@echo "  test-frontend Run frontend tests only"
	@echo ""
	@echo "🔨 Build Commands:"
	@echo "  build       Build backend binary"
	@echo "  clean       Clean build artifacts"
	@echo ""
	@echo "📊 Utility Commands:"
	@echo "  logs        Show application logs"
	@echo "  health      Check application health"
	@echo "  db-seed     Seed database with sample data"

# Install dependencies
install:
	@echo "📦 Installing dependencies..."
	@./setup.sh

# Setup PostgreSQL
setup:
	@echo "🐘 Setting up PostgreSQL..."
	@./setup-postgres.sh

# Run full application
run:
	@echo "🚀 Starting full application..."
	@echo "Backend: http://localhost:8080"
	@echo "Frontend: http://localhost:3000"
	@echo "Swagger: http://localhost:8080/swagger/index.html"
	@echo ""
	@echo "Press Ctrl+C to stop"
	@cd backend && go run cmd/main.go & \
	cd frontend && npm run dev & \
	wait

# Run backend only
backend:
	@echo "🔧 Starting backend server..."
	@cd backend && go run cmd/main.go

# Run frontend only
frontend:
	@echo "🎨 Starting frontend server..."
	@cd frontend && npm run dev

# Database migrations
migrate:
	@echo "🔄 Running database migrations..."
	@cd backend && go run cmd/migrate/main.go migrate

rollback:
	@echo "⏪ Rolling back last migration..."
	@cd backend && go run cmd/migrate/main.go rollback

rollback-to:
	@if [ -z "$(id)" ]; then \
		echo "❌ Migration ID is required. Usage: make rollback-to id=MIGRATION_ID"; \
		echo "Available migrations:"; \
		echo "  001_create_books_table"; \
		echo "  002_add_indexes_to_books"; \
		echo "  003_add_soft_delete_to_books"; \
		exit 1; \
	fi
	@echo "⏪ Rolling back to migration: $(id)"
	@cd backend && go run cmd/migrate/main.go rollback-to -id=$(id)

status:
	@echo "📊 Migration status:"
	@cd backend && go run cmd/migrate/main.go status

applied:
	@echo "📋 Applied migrations:"
	@cd backend && go run cmd/migrate/main.go applied

# Database reset
db-reset: rollback migrate
	@echo "🔄 Database reset complete!"

# Docker commands
docker-up:
	@echo "🐳 Starting services with Docker Compose..."
	@docker-compose up -d
	@echo "✅ Services started!"
	@echo "Backend: http://localhost:8080"
	@echo "Frontend: http://localhost:3000"
	@echo "PostgreSQL: localhost:5432"

docker-down:
	@echo "🛑 Stopping Docker services..."
	@docker-compose down
	@echo "✅ Services stopped!"

# Testing
test:
	@echo "🧪 Running all tests..."
	@cd backend && go test ./... -v
	@cd frontend && npm test

test-backend:
	@echo "🧪 Running backend tests..."
	@cd backend && go test ./... -v

test-frontend:
	@echo "🧪 Running frontend tests..."
	@cd frontend && npm test

# Build
build:
	@echo "🔨 Building backend binary..."
	@cd backend && go build -o bin/main cmd/main.go
	@echo "✅ Binary built: backend/bin/main"

clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf backend/bin
	@rm -rf frontend/.next
	@rm -rf frontend/node_modules/.cache
	@echo "✅ Cleaned!"

# Utility commands
logs:
	@echo "📋 Application logs:"
	@docker-compose logs -f

health:
	@echo "🏥 Checking application health..."
	@curl -s http://localhost:8080/health | jq . || echo "Backend not running"
	@curl -s http://localhost:3000 > /dev/null && echo "Frontend: ✅ Running" || echo "Frontend: ❌ Not running"

# Development helpers
dev-setup: install setup
	@echo "✅ Development environment ready!"

quick-start: dev-setup migrate
	@echo "🚀 Quick start complete!"
	@echo "Run 'make run' to start the application"

# Database seeding
db-seed:
	@echo "🌱 Seeding database with sample data..."
	@curl -X POST http://localhost:8080/api/books \
		-H "Content-Type: application/json" \
		-d '{"title":"The Great Gatsby","author":"F. Scott Fitzgerald","year":1925,"isbn":"978-0743273565"}' || echo "Backend not running"
	@curl -X POST http://localhost:8080/api/books \
		-H "Content-Type: application/json" \
		-d '{"title":"To Kill a Mockingbird","author":"Harper Lee","year":1960,"isbn":"978-0446310789"}' || echo "Backend not running"
	@curl -X POST http://localhost:8080/api/books \
		-H "Content-Type: application/json" \
		-d '{"title":"1984","author":"George Orwell","year":1949,"isbn":"978-0451524935"}' || echo "Backend not running"
	@echo "✅ Sample data added!"

# API testing
test-api:
	@echo "🧪 Testing API endpoints..."
	@echo "Creating a test book..."
	@curl -X POST http://localhost:8080/api/books \
		-H "Content-Type: application/json" \
		-d '{"title":"Test Book","author":"Test Author","year":2024,"isbn":"1234567890"}' || echo "Failed to create book"
	@echo ""
	@echo "Getting all books..."
	@curl -s http://localhost:8080/api/books | jq . || echo "Failed to get books"
	@echo ""
	@echo "Searching books..."
	@curl -s "http://localhost:8080/api/books/search?title=Test" | jq . || echo "Failed to search books"

# Migration helpers
migration-status: status
	@echo ""
	@echo "📋 Available migrations:"
	@echo "  001_create_books_table"
	@echo "  002_add_indexes_to_books"
	@echo "  003_add_soft_delete_to_books"

migration-help:
	@echo "🔄 Migration Commands:"
	@echo "  make migrate                    # Run all pending migrations"
	@echo "  make rollback                   # Rollback last migration"
	@echo "  make rollback-to id=MIGRATION_ID # Rollback to specific migration"
	@echo "  make status                     # Show migration status"
	@echo "  make applied                    # Show applied migrations"
	@echo "  make db-reset                   # Reset database (rollback all + migrate)"
	@echo ""
	@echo "📋 Available Migration IDs:"
	@echo "  001_create_books_table"
	@echo "  002_add_indexes_to_books"
	@echo "  003_add_soft_delete_to_books" 