#!/bin/bash

echo "🐘 Setting up PostgreSQL for Library Management System..."

# Check if PostgreSQL is installed
if ! command -v psql &> /dev/null; then
    echo "❌ PostgreSQL is not installed. Please install PostgreSQL first."
    echo "   On macOS: brew install postgresql"
    echo "   On Ubuntu: sudo apt-get install postgresql postgresql-contrib"
    echo "   On CentOS: sudo yum install postgresql postgresql-server"
    exit 1
fi

# Check if PostgreSQL service is running
if ! pg_isready -q; then
    echo "❌ PostgreSQL service is not running. Please start PostgreSQL first."
    echo "   On macOS: brew services start postgresql"
    echo "   On Ubuntu: sudo systemctl start postgresql"
    echo "   On CentOS: sudo systemctl start postgresql"
    exit 1
fi

echo "✅ PostgreSQL is installed and running"

# Get database configuration from environment or use defaults
DB_NAME=${DB_NAME:-library_db}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-password}

echo "📝 Creating database: $DB_NAME"
echo "👤 Database user: $DB_USER"

# Create database if it doesn't exist
if psql -lqt | cut -d \| -f 1 | grep -qw $DB_NAME; then
    echo "✅ Database '$DB_NAME' already exists"
else
    echo "🔧 Creating database '$DB_NAME'..."
    createdb $DB_NAME
    if [ $? -eq 0 ]; then
        echo "✅ Database '$DB_NAME' created successfully"
    else
        echo "❌ Failed to create database '$DB_NAME'"
        exit 1
    fi
fi

# Create user if it doesn't exist (only if not using default postgres user)
if [ "$DB_USER" != "postgres" ]; then
    echo "👤 Creating user '$DB_USER'..."
    psql -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';" 2>/dev/null || echo "User '$DB_USER' already exists or failed to create"
    psql -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;" 2>/dev/null || echo "Failed to grant privileges"
fi

echo ""
echo "🎉 PostgreSQL setup completed!"
echo ""
echo "📋 Next steps:"
echo "1. Update your .env file with the correct database credentials:"
echo "   DB_TYPE=postgres"
echo "   DB_HOST=localhost"
echo "   DB_PORT=5432"
echo "   DB_USER=$DB_USER"
echo "   DB_PASSWORD=$DB_PASSWORD"
echo "   DB_NAME=$DB_NAME"
echo "   DB_SSL_MODE=disable"
echo ""
echo "2. Start the backend:"
echo "   cd backend && go run cmd/main.go"
echo ""
echo "3. The application will automatically create the required tables"
echo ""
echo "🔗 Connection string: postgres://$DB_USER:$DB_PASSWORD@localhost:5432/$DB_NAME" 