// Package config handles configuration management.

package config

import "os"

// AppConfig represents the application's configuraton.
type AppConfig struct {
	DataSourceName string // Database connection string
	Environment    string // Apllication environment (e.g., "dev", "prod")
	ServerAddress  string // Address on which the server should listen
	SwaggerHost    string // Host for Swagger documentation
}

// NewConfig initializes and returns a new AppConfig with default values obtained from environment variables.
func NewConfig() *AppConfig {
	return &AppConfig{
		DataSourceName: GetEnv("MYSQL_DSN", "default_dsn"),
		Environment:    GetEnv("GO_ENV", "dev"),
		ServerAddress:  GetEnv("SERVER_ADDRESS", "0.0.0.0:8080"),
		SwaggerHost:    GetDefaultSwaggerHost(GetEnv("GO_ENV", "dev")),
	}
}

// GetEnv retrieves the value of an environment variable or returns a fallback if it doesn't exist.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// GetDefaultSwaggerHost returns the default Swagger host based on the environment.
func GetDefaultSwaggerHost(env string) string {
	switch env {
	case "dev":
		return "localhost:8080"
	default:
		return ""
	}
}
