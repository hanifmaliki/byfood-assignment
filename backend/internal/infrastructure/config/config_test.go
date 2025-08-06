package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	// Save original environment variables
	originalEnv := make(map[string]string)
	envVars := []string{
		"BACKEND_PORT", "BACKEND_HOST", "BACKEND_ENVIRONMENT",
		"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_SSL_MODE",
		"API_VERSION", "API_PREFIX", "API_TIMEOUT",
		"CORS_ALLOWED_ORIGINS", "CORS_ALLOWED_METHODS", "CORS_ALLOWED_HEADERS",
		"LOG_LEVEL", "LOG_FORMAT", "LOG_FILE",
		"SWAGGER_ENABLED", "SWAGGER_TITLE", "SWAGGER_DESCRIPTION", "SWAGGER_VERSION",
		"JWT_SECRET", "JWT_EXPIRY",
	}

	for _, envVar := range envVars {
		if value := os.Getenv(envVar); value != "" {
			originalEnv[envVar] = value
		}
	}

	// Clean up after test
	defer func() {
		for envVar, value := range originalEnv {
			os.Setenv(envVar, value)
		}
		for _, envVar := range envVars {
			if _, exists := originalEnv[envVar]; !exists {
				os.Unsetenv(envVar)
			}
		}
	}()

	// Test default values
	config := Load()

	assert.NotNil(t, config)
	assert.Equal(t, "8080", config.Server.Port)
	assert.Equal(t, "localhost", config.Server.Host)
	assert.Equal(t, "development", config.Server.Environment)

	assert.Equal(t, "localhost", config.Database.Host)
	assert.Equal(t, "5432", config.Database.Port)
	assert.Equal(t, "postgres", config.Database.User)
	assert.Equal(t, "password", config.Database.Password)
	assert.Equal(t, "library_db", config.Database.Name)
	assert.Equal(t, "disable", config.Database.SSLMode)

	assert.Equal(t, "v1", config.API.Version)
	assert.Equal(t, "/api", config.API.Prefix)
	assert.Equal(t, "30s", config.API.Timeout)

	assert.Equal(t, []string{"http://localhost:3000", "http://localhost:3001"}, config.CORS.AllowedOrigins)
	assert.Equal(t, []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, config.CORS.AllowedMethods)
	assert.Equal(t, []string{"Content-Type", "Authorization"}, config.CORS.AllowedHeaders)

	assert.Equal(t, "info", config.Logging.Level)
	assert.Equal(t, "json", config.Logging.Format)
	assert.Equal(t, "./logs/app.log", config.Logging.File)

	assert.True(t, config.Swagger.Enabled)
	assert.Equal(t, "Library Management System API", config.Swagger.Title)
	assert.Equal(t, "A RESTful API for managing books and URL processing", config.Swagger.Description)
	assert.Equal(t, "1.0", config.Swagger.Version)

	assert.Equal(t, "your-super-secret-jwt-key-change-this-in-production", config.Security.JWTSecret)
	assert.Equal(t, "24h", config.Security.JWTExpiry)
}

func TestLoad_WithEnvironmentVariables(t *testing.T) {
	// Save original environment variables
	originalEnv := make(map[string]string)
	envVars := []string{
		"BACKEND_PORT", "BACKEND_HOST", "BACKEND_ENVIRONMENT",
		"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_SSL_MODE",
		"API_VERSION", "API_PREFIX", "API_TIMEOUT",
		"CORS_ALLOWED_ORIGINS", "CORS_ALLOWED_METHODS", "CORS_ALLOWED_HEADERS",
		"LOG_LEVEL", "LOG_FORMAT", "LOG_FILE",
		"SWAGGER_ENABLED", "SWAGGER_TITLE", "SWAGGER_DESCRIPTION", "SWAGGER_VERSION",
		"JWT_SECRET", "JWT_EXPIRY",
	}

	for _, envVar := range envVars {
		if value := os.Getenv(envVar); value != "" {
			originalEnv[envVar] = value
		}
	}

	// Clean up after test
	defer func() {
		for envVar, value := range originalEnv {
			os.Setenv(envVar, value)
		}
		for _, envVar := range envVars {
			if _, exists := originalEnv[envVar]; !exists {
				os.Unsetenv(envVar)
			}
		}
	}()

	// Set custom environment variables
	os.Setenv("BACKEND_PORT", "9090")
	os.Setenv("BACKEND_HOST", "0.0.0.0")
	os.Setenv("BACKEND_ENVIRONMENT", "production")
	os.Setenv("DB_HOST", "db.example.com")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_USER", "custom_user")
	os.Setenv("DB_PASSWORD", "custom_password")
	os.Setenv("DB_NAME", "custom_db")
	os.Setenv("DB_SSL_MODE", "require")
	os.Setenv("API_VERSION", "v2")
	os.Setenv("API_PREFIX", "/api/v2")
	os.Setenv("API_TIMEOUT", "60s")
	os.Setenv("CORS_ALLOWED_ORIGINS", "https://example.com,https://app.example.com")
	os.Setenv("CORS_ALLOWED_METHODS", "GET,POST")
	os.Setenv("CORS_ALLOWED_HEADERS", "Content-Type")
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("LOG_FORMAT", "text")
	os.Setenv("LOG_FILE", "/var/log/app.log")
	os.Setenv("SWAGGER_ENABLED", "false")
	os.Setenv("SWAGGER_TITLE", "Custom API")
	os.Setenv("SWAGGER_DESCRIPTION", "Custom API Description")
	os.Setenv("SWAGGER_VERSION", "2.0")
	os.Setenv("JWT_SECRET", "custom-jwt-secret")
	os.Setenv("JWT_EXPIRY", "12h")

	config := Load()

	assert.NotNil(t, config)
	assert.Equal(t, "9090", config.Server.Port)
	assert.Equal(t, "0.0.0.0", config.Server.Host)
	assert.Equal(t, "production", config.Server.Environment)

	assert.Equal(t, "db.example.com", config.Database.Host)
	assert.Equal(t, "5433", config.Database.Port)
	assert.Equal(t, "custom_user", config.Database.User)
	assert.Equal(t, "custom_password", config.Database.Password)
	assert.Equal(t, "custom_db", config.Database.Name)
	assert.Equal(t, "require", config.Database.SSLMode)

	assert.Equal(t, "v2", config.API.Version)
	assert.Equal(t, "/api/v2", config.API.Prefix)
	assert.Equal(t, "60s", config.API.Timeout)

	assert.Equal(t, []string{"https://example.com", "https://app.example.com"}, config.CORS.AllowedOrigins)
	assert.Equal(t, []string{"GET", "POST"}, config.CORS.AllowedMethods)
	assert.Equal(t, []string{"Content-Type"}, config.CORS.AllowedHeaders)

	assert.Equal(t, "debug", config.Logging.Level)
	assert.Equal(t, "text", config.Logging.Format)
	assert.Equal(t, "/var/log/app.log", config.Logging.File)

	assert.False(t, config.Swagger.Enabled)
	assert.Equal(t, "Custom API", config.Swagger.Title)
	assert.Equal(t, "Custom API Description", config.Swagger.Description)
	assert.Equal(t, "2.0", config.Swagger.Version)

	assert.Equal(t, "custom-jwt-secret", config.Security.JWTSecret)
	assert.Equal(t, "12h", config.Security.JWTExpiry)
}

func TestGetEnv(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("TEST_ENV_VAR")

	// Clean up after test
	defer func() {
		if originalValue != "" {
			os.Setenv("TEST_ENV_VAR", originalValue)
		} else {
			os.Unsetenv("TEST_ENV_VAR")
		}
	}()

	// Test with environment variable set
	os.Setenv("TEST_ENV_VAR", "test_value")
	assert.Equal(t, "test_value", getEnv("TEST_ENV_VAR", "default_value"))

	// Test with environment variable not set
	os.Unsetenv("TEST_ENV_VAR")
	assert.Equal(t, "default_value", getEnv("TEST_ENV_VAR", "default_value"))
}

func TestGetEnvBool(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("TEST_BOOL_VAR")

	// Clean up after test
	defer func() {
		if originalValue != "" {
			os.Setenv("TEST_BOOL_VAR", originalValue)
		} else {
			os.Unsetenv("TEST_BOOL_VAR")
		}
	}()

	// Test with "true" value
	os.Setenv("TEST_BOOL_VAR", "true")
	assert.True(t, getEnvBool("TEST_BOOL_VAR", false))

	// Test with "false" value
	os.Setenv("TEST_BOOL_VAR", "false")
	assert.False(t, getEnvBool("TEST_BOOL_VAR", true))

	// Test with invalid value
	os.Setenv("TEST_BOOL_VAR", "invalid")
	assert.True(t, getEnvBool("TEST_BOOL_VAR", true))

	// Test with environment variable not set
	os.Unsetenv("TEST_BOOL_VAR")
	assert.True(t, getEnvBool("TEST_BOOL_VAR", true))
	assert.False(t, getEnvBool("TEST_BOOL_VAR", false))
}

func TestGetEnvInt(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("TEST_INT_VAR")

	// Clean up after test
	defer func() {
		if originalValue != "" {
			os.Setenv("TEST_INT_VAR", originalValue)
		} else {
			os.Unsetenv("TEST_INT_VAR")
		}
	}()

	// Test with valid integer value
	os.Setenv("TEST_INT_VAR", "123")
	assert.Equal(t, 123, getEnvInt("TEST_INT_VAR", 0))

	// Test with invalid value
	os.Setenv("TEST_INT_VAR", "invalid")
	assert.Equal(t, 456, getEnvInt("TEST_INT_VAR", 456))

	// Test with environment variable not set
	os.Unsetenv("TEST_INT_VAR")
	assert.Equal(t, 789, getEnvInt("TEST_INT_VAR", 789))
}

func TestConfig_StringRepresentation(t *testing.T) {
	config := Load()

	// Test that config can be converted to string (for logging purposes)
	// This is an indirect test that the struct is properly formed
	assert.NotEmpty(t, config.Server.Port)
	assert.NotEmpty(t, config.Database.Host)
	assert.NotEmpty(t, config.API.Prefix)
	assert.NotEmpty(t, config.Logging.Level)
	assert.NotEmpty(t, config.Swagger.Title)
	assert.NotEmpty(t, config.Security.JWTSecret)
}

func TestConfig_Validation(t *testing.T) {
	config := Load()

	// Test that required fields are not empty
	assert.NotEmpty(t, config.Server.Port)
	assert.NotEmpty(t, config.Server.Host)
	assert.NotEmpty(t, config.Database.Host)
	assert.NotEmpty(t, config.Database.Port)
	assert.NotEmpty(t, config.Database.User)
	assert.NotEmpty(t, config.Database.Name)
	assert.NotEmpty(t, config.API.Prefix)
	assert.NotEmpty(t, config.Logging.Level)
	assert.NotEmpty(t, config.Swagger.Title)
	assert.NotEmpty(t, config.Security.JWTSecret)

	// Test that arrays are properly initialized
	assert.NotNil(t, config.CORS.AllowedOrigins)
	assert.NotNil(t, config.CORS.AllowedMethods)
	assert.NotNil(t, config.CORS.AllowedHeaders)
	assert.Greater(t, len(config.CORS.AllowedOrigins), 0)
	assert.Greater(t, len(config.CORS.AllowedMethods), 0)
	assert.Greater(t, len(config.CORS.AllowedHeaders), 0)
}
