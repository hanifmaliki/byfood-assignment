#!/bin/bash

echo "🚀 Setting up Library Management System..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    exit 1
fi

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js 18 or higher."
    exit 1
fi

# Check if npm is installed
if ! command -v npm &> /dev/null; then
    echo "❌ npm is not installed. Please install npm."
    exit 1
fi

echo "✅ Prerequisites check passed"

# Setup Backend
echo "📦 Setting up backend..."
cd backend

# Initialize Go modules
go mod tidy

# Install dependencies
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/swag
go get -u github.com/swaggo/files
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u github.com/google/uuid

echo "✅ Backend setup completed"

# Setup Frontend
echo "📦 Setting up frontend..."
cd ../frontend

# Install dependencies
npm install

echo "✅ Frontend setup completed"

echo ""
echo "🎉 Setup completed successfully!"
echo ""
echo "To start the application:"
echo "1. Start the backend: cd backend && go run cmd/main.go"
echo "2. Start the frontend: cd frontend && npm run dev"
echo ""
echo "Backend will be available at: http://localhost:8080"
echo "Frontend will be available at: http://localhost:3000"
echo "API Documentation will be available at: http://localhost:8080/swagger/index.html" 