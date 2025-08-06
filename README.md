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
- 🗄️ SQLite database integration
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

#### Backend (.env)
```bash
# Server Configuration
PORT=8080
HOST=localhost

# Database Configuration
DB_TYPE=sqlite
DB_PATH=library.db

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

## Quick Start

### 1. Clone the repository
```bash
git clone <repository-url>
cd byfood-assignment
```

### 2. Set up environment variables
```bash
cp env.example .env
# Edit .env with your configuration
```

### 3. Backend Setup
```bash
cd backend
cp env.example .env
go mod tidy
go run cmd/main.go
```

The backend will start on `http://localhost:8080`

### 4. Frontend Setup
```bash
cd frontend
cp env.example .env.local
npm install
npm run dev
```

The frontend will start on `http://localhost:3000`

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
- Use local database: `DB_PATH=library.db`

### Production
- Set `NODE_ENV=production`
- Disable debug mode: `NEXT_PUBLIC_DEBUG=false`
- Use production database
- Set proper CORS origins
- Configure rate limiting

### Security Notes
- **Always change default JWT secrets** in production
- **Use strong passwords** for database connections
- **Enable HTTPS** in production
- **Configure proper CORS** origins
- **Set up rate limiting** for API endpoints

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is for assignment purposes. 