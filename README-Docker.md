# Docker Setup for Library Management System

This project includes separate Docker configurations for development and production environments.

## 🏗️ Docker Files Structure

```
├── backend/
│   ├── Dockerfile.dev      # Development: bind mounts, live reload
│   └── Dockerfile          # Production: multi-stage build, optimized
├── frontend/
│   ├── Dockerfile.dev      # Development: bind mounts, live reload
│   └── Dockerfile          # Production: multi-stage build, optimized
├── docker-compose.yml      # Production environment (default)
├── docker-compose.dev.yml  # Development environment
└── docker-compose.yml.backup # Original backup
```

## 🚀 Development Environment

### Features
- **Live Reload**: Code changes automatically restart services
- **Bind Mounts**: Source code is mounted from host to container
- **Hot Reload**: Go and Next.js hot reloading enabled
- **Debugging**: Full source maps and debugging capabilities

### Commands

```bash
# Start development environment
docker-compose -f docker-compose.dev.yml up -d

# Start with rebuild
docker-compose -f docker-compose.dev.yml up -d --build

# View logs
docker-compose -f docker-compose.dev.yml logs -f

# Stop development environment
docker-compose -f docker-compose.dev.yml down

# Stop and remove volumes
docker-compose -f docker-compose.dev.yml down -v
```

### Development URLs
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Swagger Docs**: http://localhost:8080/swagger/index.html
- **Database**: localhost:5432

## 🏭 Production Environment

### Features
- **Multi-stage Builds**: Optimized, smaller production images
- **Pre-built Binaries**: Go binary and Next.js build artifacts
- **Security**: Minimal runtime dependencies
- **Performance**: Optimized for production workloads

### Commands

```bash
# Build and start production environment (default)
docker-compose up -d --build

# Start production environment
docker-compose up -d

# View logs
docker-compose logs -f

# Stop production environment
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

### Environment Variables

Create a `.env` file for production:

```bash
# Database
POSTGRES_PASSWORD=your_secure_password

# Backend
SWAGGER_ENABLED=false
CORS_ALLOWED_ORIGINS=https://yourdomain.com

# Force rebuild (optional)
REBUILD=$(date +%s)
```

## 🔧 Customization

### Development Environment Variables

```bash
# Override development settings
export POSTGRES_PASSWORD=custom_password
export CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001

docker-compose -f docker-compose.dev.yml up -d
```

### Production Build Arguments

```bash
# Force rebuild of backend
docker-compose build --build-arg REBUILD=$(date +%s) backend

# Build specific service
docker-compose build backend
```

## 📊 Performance Comparison

| Aspect | Development | Production |
|--------|-------------|------------|
| **Image Size** | Larger (includes dev tools) | Smaller (optimized) |
| **Startup Time** | Faster (no build step) | Slower (build required) |
| **Live Reload** | ✅ Yes | ❌ No |
| **Resource Usage** | Higher | Lower |
| **Debugging** | ✅ Full | ❌ Limited |

## 🛠️ Troubleshooting

### Common Issues

1. **Port Conflicts**
   ```bash
   # Check what's using the ports
   lsof -i :3000
   lsof -i :8080
   lsof -i :5432
   ```

2. **Permission Issues**
   ```bash
   # Fix file permissions
   sudo chown -R $USER:$USER .
   ```

3. **Build Cache Issues**
   ```bash
   # Clear build cache
   docker system prune -a
   docker volume prune
   ```

4. **Database Connection Issues**
   ```bash
   # Check database logs
   docker-compose -f docker-compose.dev.yml logs postgres
   ```

### Reset Everything

```bash
# Stop all containers and remove everything
docker-compose -f docker-compose.dev.yml down -v
docker-compose down -v
docker system prune -a --volumes

# Rebuild from scratch
docker-compose -f docker-compose.dev.yml up -d --build
```

## 🔄 Switching Between Environments

```bash
# Stop current environment
docker-compose -f docker-compose.dev.yml down

# Start production environment (default)
docker-compose up -d

# Or vice versa
docker-compose down
docker-compose -f docker-compose.dev.yml up -d
```

## 📝 Notes

- **Development**: Use for coding, testing, and debugging
- **Production**: Use for staging, testing, and deployment (default)
- **Database**: Data persists between environment switches
- **Networks**: Separate networks prevent conflicts
- **Volumes**: Development uses bind mounts, production uses named volumes 