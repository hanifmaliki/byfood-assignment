# Project Structure Documentation

## Overview

This is a full-stack library management system built with Next.js frontend and Go backend using clean architecture principles.

## Directory Structure

```
byfood-assignment/
├── README.md                    # Main project documentation
├── requirements.txt             # Original assignment requirements
├── setup.sh                    # Automated setup script
├── PROJECT_STRUCTURE.md        # This file - detailed structure
│
├── backend/                    # Go backend (Clean Architecture)
│   ├── go.mod                  # Go module definition
│   ├── cmd/
│   │   └── main.go            # Application entry point
│   ├── internal/               # Internal packages
│   │   ├── domain/            # Business entities and interfaces
│   │   │   ├── entities/
│   │   │   │   ├── book.go    # Book entity
│   │   │   │   └── url.go     # URL processing entities
│   │   │   └── repositories/
│   │   │       ├── book_repository.go    # Book repository interface
│   │   │       └── url_repository.go     # URL repository interface
│   │   ├── usecase/           # Business logic layer
│   │   │   ├── book_usecase.go    # Book business logic
│   │   │   └── url_usecase.go     # URL processing business logic
│   │   ├── repository/         # Data access layer
│   │   │   ├── book_repository_impl.go   # Book repository implementation
│   │   │   └── url_repository_impl.go    # URL repository implementation
│   │   ├── delivery/           # HTTP handlers and routing
│   │   │   └── http/
│   │   │       └── handlers/
│   │   │           ├── book_handlers.go   # Book HTTP handlers
│   │   │           └── url_handlers.go    # URL HTTP handlers
│   │   └── infrastructure/     # External concerns
│   │       └── database/
│   │           └── database.go # Database connection and setup
│   ├── docs/                   # Swagger documentation
│   │   └── docs.go            # Auto-generated Swagger docs
│   └── test/                   # Backend tests
│       └── main_test.go       # Basic test example
│
└── frontend/                   # Next.js frontend
    ├── package.json            # Node.js dependencies
    ├── next.config.js         # Next.js configuration
    ├── tailwind.config.js     # Tailwind CSS configuration
    ├── postcss.config.js      # PostCSS configuration
    ├── tsconfig.json          # TypeScript configuration
    ├── jest.config.js         # Jest testing configuration
    ├── jest.setup.js          # Jest setup file
    └── src/                   # Source code
        ├── app/               # Next.js 14 app directory
        │   ├── globals.css    # Global styles
        │   ├── layout.tsx     # Root layout
        │   ├── page.tsx       # Main dashboard page
        │   └── test/          # Frontend tests
        │       └── page.test.tsx
        ├── components/        # React components
        │   ├── BookCard.tsx   # Individual book display
        │   └── BookForm.tsx   # Add/edit book form
        ├── contexts/          # Context API for state management
        │   └── BookContext.tsx # Book state management
        ├── types/             # TypeScript type definitions
        │   ├── book.ts        # Book-related types
        │   └── url.ts         # URL processing types
        └── utils/             # Utility functions
            └── api.ts         # API client functions
```

## Architecture Overview

### Backend (Clean Architecture)

The backend follows clean architecture principles with clear separation of concerns:

1. **Domain Layer** (`internal/domain/`)
   - Contains business entities and repository interfaces
   - Defines the core business rules and entities
   - No dependencies on external frameworks

2. **Use Case Layer** (`internal/usecase/`)
   - Contains business logic and orchestration
   - Implements application-specific business rules
   - Depends only on domain layer

3. **Repository Layer** (`internal/repository/`)
   - Implements data access logic
   - Handles database operations
   - Implements domain repository interfaces

4. **Delivery Layer** (`internal/delivery/`)
   - Handles HTTP requests and responses
   - Manages routing and request/response formatting
   - Depends on use case layer

5. **Infrastructure Layer** (`internal/infrastructure/`)
   - Contains external concerns like database setup
   - Handles configuration and external service integration

### Frontend (Next.js 14)

The frontend is built with modern React patterns:

1. **App Directory** (`src/app/`)
   - Uses Next.js 14 app directory structure
   - Server-side rendering capabilities
   - File-based routing

2. **Components** (`src/components/`)
   - Reusable React components
   - Form validation and error handling
   - Modal dialogs and responsive design

3. **Context API** (`src/contexts/`)
   - Global state management
   - Real-time updates across components
   - Error handling and loading states

4. **Type Safety** (`src/types/`)
   - TypeScript interfaces for all data structures
   - API request/response type definitions
   - Strict type checking

## Key Features Implemented

### Backend Features
- ✅ RESTful API for book management (CRUD operations)
- ✅ URL cleanup and redirection service
- ✅ SQLite database integration
- ✅ Input validation and error handling
- ✅ Comprehensive logging
- ✅ Swagger API documentation
- ✅ Clean Architecture implementation
- ✅ CORS middleware for frontend integration

### Frontend Features
- ✅ Modern UI with Tailwind CSS
- ✅ Responsive design for all screen sizes
- ✅ Form validation with visual feedback
- ✅ Modal dialogs for better UX
- ✅ Real-time state management with Context API
- ✅ TypeScript for type safety
- ✅ Error handling and loading states
- ✅ Delete confirmation dialogs

### Integration Features
- ✅ API proxy configuration in Next.js
- ✅ CORS setup for cross-origin requests
- ✅ Comprehensive error handling
- ✅ Loading states and user feedback

## API Endpoints

### Book Management
- `GET /api/books` - Get all books
- `POST /api/books` - Create a new book
- `GET /api/books/{id}` - Get book by ID
- `PUT /api/books/{id}` - Update book by ID
- `DELETE /api/books/{id}` - Delete book by ID

### URL Processing
- `POST /api/url/process` - Process URL cleanup and redirection

## Testing

### Backend Testing
- Unit tests for business logic
- Integration tests for API endpoints
- Database testing with SQLite

### Frontend Testing
- Component testing with React Testing Library
- Jest configuration for modern React
- Mock implementations for API calls

## Development Workflow

1. **Setup**: Run `./setup.sh` to install dependencies
2. **Backend**: `cd backend && go run cmd/main.go`
3. **Frontend**: `cd frontend && npm run dev`
4. **Testing**: 
   - Backend: `cd backend && go test ./...`
   - Frontend: `cd frontend && npm test`

## Technologies Used

### Backend
- **Go 1.21+** - Programming language
- **Gin** - HTTP web framework
- **GORM** - ORM for database operations
- **SQLite** - Database
- **Swagger** - API documentation
- **Clean Architecture** - Design pattern

### Frontend
- **Next.js 14** - React framework
- **TypeScript** - Type safety
- **Tailwind CSS** - Styling
- **Lucide React** - Icons
- **Axios** - HTTP client
- **Context API** - State management

## Deployment Considerations

### Backend
- Can be containerized with Docker
- SQLite can be replaced with PostgreSQL/MySQL
- Environment variables for configuration
- Health check endpoint available

### Frontend
- Static export capability
- Environment variables for API configuration
- Optimized for production builds
- Responsive design for mobile devices 