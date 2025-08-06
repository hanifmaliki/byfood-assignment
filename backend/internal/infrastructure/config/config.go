package config

import (
	"os"
	"strconv"
	"strings"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	API      APIConfig
	CORS     CORSConfig
	Logging  LoggingConfig
	Swagger  SwaggerConfig
	Security SecurityConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port        string
	Host        string
	Environment string
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Type string
	Path string
}

// APIConfig holds API configuration
type APIConfig struct {
	Version string
	Prefix  string
	Timeout string
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string
	Format string
	File   string
}

// SwaggerConfig holds Swagger configuration
type SwaggerConfig struct {
	Enabled     bool
	Title       string
	Description string
	Version     string
}

// SecurityConfig holds security configuration
type SecurityConfig struct {
	JWTSecret string
	JWTExpiry string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:        getEnv("BACKEND_PORT", "8080"),
			Host:        getEnv("BACKEND_HOST", "localhost"),
			Environment: getEnv("BACKEND_ENVIRONMENT", "development"),
		},
		Database: DatabaseConfig{
			Type: getEnv("DB_TYPE", "sqlite"),
			Path: getEnv("DB_PATH", "library.db"),
		},
		API: APIConfig{
			Version: getEnv("API_VERSION", "v1"),
			Prefix:  getEnv("API_PREFIX", "/api"),
			Timeout: getEnv("API_TIMEOUT", "30s"),
		},
		CORS: CORSConfig{
			AllowedOrigins: strings.Split(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:3001"), ","),
			AllowedMethods: strings.Split(getEnv("CORS_ALLOWED_METHODS", "GET,POST,PUT,DELETE,OPTIONS"), ","),
			AllowedHeaders: strings.Split(getEnv("CORS_ALLOWED_HEADERS", "Content-Type,Authorization"), ","),
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
			File:   getEnv("LOG_FILE", "./logs/app.log"),
		},
		Swagger: SwaggerConfig{
			Enabled:     getEnvBool("SWAGGER_ENABLED", true),
			Title:       getEnv("SWAGGER_TITLE", "Library Management System API"),
			Description: getEnv("SWAGGER_DESCRIPTION", "A RESTful API for managing books and URL processing"),
			Version:     getEnv("SWAGGER_VERSION", "1.0"),
		},
		Security: SecurityConfig{
			JWTSecret: getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-this-in-production"),
			JWTExpiry: getEnv("JWT_EXPIRY", "24h"),
		},
	}
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// getEnvBool gets environment variable as boolean with fallback
func getEnvBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}

// getEnvInt gets environment variable as integer with fallback
func getEnvInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
