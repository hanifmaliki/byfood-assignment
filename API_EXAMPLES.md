# 📚 Library Management System API - Usage Examples

This document provides comprehensive examples of how to use the Library Management System API endpoints.

## 🌐 Base URL
```
http://localhost:8080/api
```

## 📖 Book Management Endpoints

### 1. Get All Books
**GET** `/books`

**Response:**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "year": 1925,
    "isbn": "978-0743273565",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "deleted_at": null
  },
  {
    "id": "7c9e6679-7425-40de-944b-e07fc1f90ae7",
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "year": 1960,
    "isbn": "978-0446310789",
    "created_at": "2024-01-15T11:15:00Z",
    "updated_at": "2024-01-15T11:15:00Z",
    "deleted_at": null
  }
]
```

### 2. Create a New Book
**POST** `/books`

**Request Body:**
```json
{
  "title": "1984",
  "author": "George Orwell",
  "year": 1949,
  "isbn": "978-0451524935"
}
```

**Response (201 Created):**
```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "title": "1984",
  "author": "George Orwell",
  "year": 1949,
  "isbn": "978-0451524935",
  "created_at": "2024-01-15T14:20:00Z",
  "updated_at": "2024-01-15T14:20:00Z",
  "deleted_at": null
}
```

**Validation Errors (400 Bad Request):**
```json
{
  "error": "book ISBN must be between 10 and 13 characters"
}
```

```json
{
  "error": "book with this ISBN already exists"
}
```

### 3. Get Book by ID
**GET** `/books/{id}`

**Example:** `GET /books/550e8400-e29b-41d4-a716-446655440000`

**Response (200 OK):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "year": 1925,
  "isbn": "978-0743273565",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z",
  "deleted_at": null
}
```

**Not Found (404):**
```json
{
  "error": "book not found"
}
```

### 4. Update Book
**PUT** `/books/{id}`

**Example:** `PUT /books/550e8400-e29b-41d4-a716-446655440000`

**Request Body:**
```json
{
  "title": "The Great Gatsby (Updated Edition)",
  "author": "F. Scott Fitzgerald",
  "year": 1925,
  "isbn": "978-0743273565"
}
```

**Response (200 OK):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "The Great Gatsby (Updated Edition)",
  "author": "F. Scott Fitzgerald",
  "year": 1925,
  "isbn": "978-0743273565",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T15:45:00Z",
  "deleted_at": null
}
```

**Note:** The `updated_at` timestamp is automatically updated.

### 5. Search Books
**GET** `/books/search?title={title}&author={author}&year={year}`

**Examples:**

**Search by title:**
```
GET /books/search?title=Gatsby
```

**Search by author:**
```
GET /books/search?author=Fitzgerald
```

**Search by year:**
```
GET /books/search?year=1925
```

**Combined search:**
```
GET /books/search?title=Gatsby&author=Fitzgerald&year=1925
```

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "year": 1925,
    "isbn": "978-0743273565",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "deleted_at": null
  }
]
```

### 6. Soft Delete Book
**DELETE** `/books/{id}`

**Example:** `DELETE /books/550e8400-e29b-41d4-a716-446655440000`

**Response (200 OK):**
```json
{
  "message": "Book deleted successfully"
}
```

**Note:** The book is marked as deleted but remains in the database.

### 7. Get Deleted Books
**GET** `/books/deleted`

**Response (200 OK):**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "year": 1925,
    "isbn": "978-0743273565",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "deleted_at": "2024-01-15T16:00:00Z"
  }
]
```

### 8. Restore Deleted Book
**POST** `/books/{id}/restore`

**Example:** `POST /books/550e8400-e29b-41d4-a716-446655440000/restore`

**Response (200 OK):**
```json
{
  "message": "Book restored successfully"
}
```

### 9. Permanent Delete Book
**DELETE** `/books/{id}/permanent`

**Example:** `DELETE /books/550e8400-e29b-41d4-a716-446655440000/permanent`

**Response (200 OK):**
```json
{
  "message": "Book permanently deleted successfully"
}
```

**Warning:** This operation cannot be undone!

## 🔗 URL Processing Endpoints

### Process URL
**POST** `/url/process`

**Request Body:**
```json
{
  "url": "https://example.com/page?utm_source=google&utm_medium=cpc&ref=123",
  "operation": "canonical"
}
```

**Available Operations:**
- `"canonical"` - Removes tracking parameters
- `"redirection"` - Follows redirects
- `"all"` - Combines both operations

**Response (200 OK):**
```json
{
  "processed_url": "https://example.com/page"
}
```

**Validation Error (400 Bad Request):**
```json
{
  "error": "invalid URL format"
}
```

## 🏥 Health Check

### Health Status
**GET** `/health`

**Response (200 OK):**
```json
{
  "status": "ok",
  "service": "Library Management System API",
  "version": "1.0"
}
```

## 🚨 Error Handling

All endpoints return consistent error responses:

**400 Bad Request:**
```json
{
  "error": "validation error message"
}
```

**404 Not Found:**
```json
{
  "error": "resource not found"
}
```

**500 Internal Server Error:**
```json
{
  "error": "internal server error"
}
```

## 📝 cURL Examples

### Create a Book
```bash
curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Hobbit",
    "author": "J.R.R. Tolkien",
    "year": 1937,
    "isbn": "978-0547928241"
  }'
```

### Update a Book
```bash
curl -X PUT http://localhost:8080/api/books/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Hobbit (Updated)",
    "author": "J.R.R. Tolkien",
    "year": 1937,
    "isbn": "978-0547928241"
  }'
```

### Search Books
```bash
curl "http://localhost:8080/api/books/search?title=Hobbit&author=Tolkien"
```

### Process URL
```bash
curl -X POST http://localhost:8080/api/url/process \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/page?utm_source=google",
    "operation": "canonical"
  }'
```

## 🔍 Swagger Documentation

Access the interactive API documentation at:
```
http://localhost:8080/swagger/index.html
```

## 📋 Notes

- **ISBN Validation**: Must be between 10-13 characters
- **Year Validation**: Cannot exceed current year
- **Timestamps**: Automatically managed by the system
- **Soft Delete**: Books are marked as deleted but can be restored
- **Search**: Case-insensitive partial matching for title and author
- **URL Processing**: Supports canonical, redirection, and combined operations 