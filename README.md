# Library Management System

A full-stack CRUD application for managing a small library of books, built with Next.js (TypeScript) frontend and Golang backend with clean architecture and URL cleanup service.

## Project Structure

```
byfood-assignment/
├── frontend/                 # Next.js frontend
│   ├── src/
│   │   ├── app/             # Next.js 14 app directory
│   │   ├── components/      # React components
│   │   ├── contexts/        # Context API for state management
│   │   ├── types/          # TypeScript type definitions
│   │   └── utils/          # Utility functions
│   ├── public/             # Static assets
│   └── package.json
├── backend/                 # Golang backend (Clean Architecture)
│   ├── cmd/                # Application entry points
│   ├── internal/           # Internal packages
│   │   ├── domain/         # Business entities and interfaces
│   │   ├── usecase/        # Business logic and use cases
│   │   ├── repository/     # Data access layer
│   │   ├── delivery/       # HTTP handlers and routing
│   │   └── infrastructure/ # External concerns (DB, etc.)
│   ├── pkg/                # Public packages
│   └── go.mod
└── README.md
```

## Features

### Frontend (Next.js 14 & TypeScript)
- 📚 Dashboard with book listing
- ➕ Add new books with form validation
- ✏️ Edit existing books
- 👁️ View book details
- 🗑️ Delete books with confirmation
- 🎨 Modern UI with modal dialogs
- 📱 Responsive design
- 🔄 Real-time state management with Context API

### Backend (Golang - Clean Architecture)
- 📖 RESTful API for book management
- 🔗 URL cleanup and redirection service
- 🗄️ PostgreSQL/SQLite database integration
- ✅ Input validation
- 📝 Comprehensive logging
- 📚 Swagger API documentation
- 🏗️ Clean Architecture with layers:
  - Domain (entities, interfaces)
  - Use Case (business logic)
  - Repository (data access)
  - Delivery (HTTP handlers)

## Prerequisites

- Node.js (v18 or higher)
- Go (v1.21 or higher)
- Git
- PostgreSQL
- Docker and Docker Compose (for containerized setup)

## Quick Start

### Option 1: Docker Compose (Recommended)

The fastest way to get started:

```bash
# 1. Clone the repository
git clone <repository-url>
cd byfood-assignment

# 2. Start all services
docker-compose -f docker-compose.dev.yml up -d

# 3. Access the application
# Frontend: http://localhost:3000
# Backend: http://localhost:8080
# Swagger: http://localhost:8080/swagger/index.html
```

### Option 2: Manual Setup

If you prefer to run services locally:

#### 1. Clone the repository
```bash
git clone <repository-url>
cd byfood-assignment
```

#### 2. Set up environment variables
```bash
cp env.example .env
# Edit .env with your configuration
```

#### 3. Set up database
```bash
./setup-postgres.sh
```

#### 4. Backend Setup
```bash
cd backend
cp env.example .env
go mod tidy
go run cmd/main.go
```

The backend will start on `http://localhost:8080`

#### 5. Frontend Setup
```bash
cd frontend
cp env.example .env.local
npm install
npm run dev
```

The frontend will start on `http://localhost:3000`

## Docker Configuration

### Docker Compose Files

The project includes two Docker Compose configurations:

#### Development Environment (`docker-compose.dev.yml`)
- **Hot Reload**: Source code is mounted for live development
- **Development Ports**: 
  - Frontend: 3000
  - Backend: 8080
  - PostgreSQL: 5432
- **Volume Mounts**: Source code, Go modules, Node modules
- **Environment**: Development settings with debug enabled

#### Production Environment (`docker-compose.yml`)
- **Optimized Builds**: Multi-stage builds for smaller images
- **Production Ports**: Same as development
- **No Volume Mounts**: Uses built binaries
- **Environment**: Production settings with security optimizations

### Docker Images

#### Backend Images
- **Development**: `Dockerfile.dev` - Includes build tools for live development
- **Production**: `Dockerfile` - Multi-stage build with minimal runtime image

#### Frontend Images
- **Development**: `Dockerfile.dev` - Includes development dependencies
- **Production**: `Dockerfile` - Multi-stage build with optimized Next.js output

### Environment Variables for Docker

The Docker Compose files use the following environment structure:

```bash
# Backend (.env.example)
DB_HOST=postgres          # Docker service name
DB_PORT=5432
DB_NAME=library_management
DB_USER=postgres
DB_PASSWORD=postgres
BACKEND_HOST=0.0.0.0      # Required for Docker networking
BACKEND_PORT=8080

# Frontend (.env.example)
NEXT_PUBLIC_API_URL=http://localhost:8080
NODE_ENV=development
```

### Docker Commands Reference

```bash
# Development
docker-compose -f docker-compose.dev.yml up -d          # Start services
docker-compose -f docker-compose.dev.yml down           # Stop services
docker-compose -f docker-compose.dev.yml logs -f        # Follow logs
docker-compose -f docker-compose.dev.yml restart backend # Restart service

# Production
docker-compose -f docker-compose.yml up -d              # Start services
docker-compose -f docker-compose.yml down               # Stop services
docker-compose -f docker-compose.yml logs -f            # Follow logs

# Build images
docker-compose -f docker-compose.dev.yml build          # Build dev images
docker-compose -f docker-compose.yml build              # Build prod images

# Clean up
docker-compose -f docker-compose.dev.yml down -v        # Remove volumes
docker system prune -a                                  # Clean unused images
```

### Development Workflow with Docker

```bash
# Start services in background
docker-compose -f docker-compose.dev.yml up -d

# Make code changes (they auto-reload)
# Backend: Edit files in ./backend/
# Frontend: Edit files in ./frontend/

# View logs for live debugging
docker-compose -f docker-compose.dev.yml logs -f backend

# Restart a specific service
docker-compose -f docker-compose.dev.yml restart backend

# Stop all services
docker-compose -f docker-compose.dev.yml down
```

## Database Setup

The application uses **gormigrate** with **timestamp-based migration naming** for better collaboration and chronological ordering.

#### Migration Naming Convention

```
Format: YYYYMMDDHHMMSS_descriptive_name
Example: 20241201000000_create_books_table
```

**Benefits:**
- ✅ Prevents conflicts when multiple developers create migrations
- ✅ Ensures chronological order
- ✅ Clear timestamp for when migration was created
- ✅ Descriptive names for easy identification

#### Available Migrations

| Timestamp | Name | Description |
|-----------|------|-------------|
| `20241201000000` | `create_books_table` | Creates the books table with basic structure |
| `20241201000001` | `add_indexes_to_books` | Adds performance indexes for title, author, year, ISBN, created_at |
| `20241201000002` | `add_soft_delete_to_books` | Adds `deleted_at` column for soft deletes |

#### Migration Commands

```bash
# Run migrations
make migrate

# Rollback last migration
make rollback

# Rollback to specific migration
make rollback-to id=20241201000001_add_indexes_to_books

# Check migration status
make status

# Show applied migrations
make applied

# Reset database (rollback all + migrate)
make db-reset
```

#### Creating New Migrations

Use the migration generator script:

```bash
# Generate a new migration
./scripts/generate_migration.sh add_user_table

# This creates: 20241201143000_add_user_table.go
```

**Manual Steps:**
1. Edit the generated migration file
2. Add the migration to `migration_manager.go`
3. Test with `make migrate`
4. Rollback if needed with `make rollback`

#### Migration File Structure

```go
// 20241201143000_add_user_table.go
package migrations

import (
    "github.com/go-gormigrate/gormigrate/v2"
    "gorm.io/gorm"
)

// AddUserTable adds user table
func AddUserTable() *gormigrate.Migration {
    return &gormigrate.Migration{
        ID: "20241201143000_add_user_table",
        Migrate: func(tx *gorm.DB) error {
            // Migration logic
            return nil
        },
        Rollback: func(tx *gorm.DB) error {
            // Rollback logic
            return nil
        },
    }
}
```

### Manual Database Operations

#### PostgreSQL
```bash
# Connect to database
psql -h localhost -U postgres -d library_db

# View tables
\dt

# View table structure
\d books

# View migration history
SELECT * FROM schema_migrations ORDER BY id;
```

## Environment Variables

### Setting Up Environment Variables

1. **Root Level Configuration:**
   ```bash
   cp env.example .env
   # Edit .env with your specific values
   ```

2. **Backend Configuration:**
   ```bash
   cd backend
   cp env.example .env
   # Edit .env with your backend-specific values
   ```

3. **Frontend Configuration:**
   ```bash
   cd frontend
   cp env.example .env.local
   # Edit .env.local with your frontend-specific values
   ```

### Key Environment Variables

#### PostgreSQL Configuration (.env)
```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_NAME=library_db
DB_SSL_MODE=disable

# Server Configuration
BACKEND_PORT=8080
BACKEND_HOST=localhost

# Security Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# CORS Configuration
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
```

#### Frontend (.env.local)
```bash
# API Configuration
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_API_VERSION=v1

# Application Configuration
NEXT_PUBLIC_APP_NAME="Library Management System"
NEXT_PUBLIC_DEBUG=true
```

## API Endpoints

### Book Management API

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/books` | Get all books |
| POST | `/api/books` | Create a new book |
| GET | `/api/books/{id}` | Get book by ID |
| PUT | `/api/books/{id}` | Update book by ID |
| DELETE | `/api/books/{id}` | Delete book by ID |

### URL Cleanup Service

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/url/process` | Process URL cleanup and redirection |

## Usage Examples

### Book API Examples

#### Create a book
```bash
curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "year": 1925,
    "isbn": "978-0743273565"
  }'
```

#### Get all books
```bash
curl http://localhost:8080/api/books
```

### URL Processing Examples

#### Canonical URL processing
```bash
curl -X POST http://localhost:8080/api/url/process \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://BYFOOD.com/food-EXPeriences?query=abc/",
    "operation": "canonical"
  }'
```

#### Redirection processing
```bash
curl -X POST http://localhost:8080/api/url/process \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://BYFOOD.com/food-EXPeriences?query=abc/",
    "operation": "redirection"
  }'
```

#### All operations
```bash
curl -X POST http://localhost:8080/api/url/process \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://BYFOOD.com/food-EXPeriences?query=abc/",
    "operation": "all"
  }'
```

## Testing

### Backend Tests
```bash
cd backend
go test ./...
```

### Frontend Tests
```bash
cd frontend
npm test
```

## API Documentation

Once the backend is running, you can access the Swagger documentation at:
`http://localhost:8080/swagger/index.html`

## Development

### Backend Development (Clean Architecture)
- **Domain Layer**: Contains business entities and interfaces
- **Use Case Layer**: Contains business logic and orchestration
- **Repository Layer**: Handles data access and persistence
- **Delivery Layer**: Manages HTTP requests and responses
- **Infrastructure Layer**: External concerns like database, logging

### Frontend Development
- Built with Next.js 14 and TypeScript
- Uses Context API for state management
- Form validation with visual feedback
- Responsive design with Tailwind CSS

## Environment Configuration

### Development
- Set `NODE_ENV=development`
- Enable debug mode: `NEXT_PUBLIC_DEBUG=true`
- Use local PostgreSQL database

### Production
- Set `NODE_ENV=production`
- Disable debug mode: `NEXT_PUBLIC_DEBUG=false`
- Use production PostgreSQL database
- Set proper CORS origins
- Configure rate limiting
- Enable SSL for PostgreSQL: `DB_SSL_MODE=require`

### Security Notes
- **Always change default JWT secrets** in production
- **Use strong passwords** for database connections
- **Enable HTTPS** in production
- **Configure proper CORS** origins
- **Set up rate limiting** for API endpoints
- **Use SSL connections** for PostgreSQL in production

## Database Migration

The application uses **gormigrate** with **timestamp-based migration naming** for better collaboration and chronological ordering.

#### Migration Naming Convention

```
Format: YYYYMMDDHHMMSS_descriptive_name
Example: 20241201000000_create_books_table
```

**Benefits:**
- ✅ Prevents conflicts when multiple developers create migrations
- ✅ Ensures chronological order
- ✅ Clear timestamp for when migration was created
- ✅ Descriptive names for easy identification

#### Available Migrations

| Timestamp | Name | Description |
|-----------|------|-------------|
| `20241201000000` | `create_books_table` | Creates the books table with basic structure |
| `20241201000001` | `add_indexes_to_books` | Adds performance indexes for title, author, year, ISBN, created_at |
| `20241201000002` | `add_soft_delete_to_books` | Adds `deleted_at` column for soft deletes |

#### Migration Commands

```bash
# Run migrations
make migrate

# Rollback last migration
make rollback

# Rollback to specific migration
make rollback-to id=20241201000001_add_indexes_to_books

# Check migration status
make status

# Show applied migrations
make applied

# Reset database (rollback all + migrate)
make db-reset
```

#### Creating New Migrations

Use the migration generator script:

```bash
# Generate a new migration
./scripts/generate_migration.sh add_user_table

# This creates: 20241201143000_add_user_table.go
```

**Manual Steps:**
1. Edit the generated migration file
2. Add the migration to `migration_manager.go`
3. Test with `make migrate`
4. Rollback if needed with `make rollback`

#### Migration File Structure

```go
// 20241201143000_add_user_table.go
package migrations

import (
    "github.com/go-gormigrate/gormigrate/v2"
    "gorm.io/gorm"
)

// AddUserTable adds user table
func AddUserTable() *gormigrate.Migration {
    return &gormigrate.Migration{
        ID: "20241201143000_add_user_table",
        Migrate: func(tx *gorm.DB) error {
            // Migration logic
            return nil
        },
        Rollback: func(tx *gorm.DB) error {
            // Rollback logic
            return nil
        },
    }
}
```

### Manual Database Operations

#### PostgreSQL
```bash
# Connect to database
psql -h localhost -U postgres -d library_db

# View tables
\dt

# View table structure
\d books

# View migration history
SELECT * FROM schema_migrations ORDER BY id;
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is for assignment purposes. 