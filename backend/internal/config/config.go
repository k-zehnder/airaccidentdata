package config

import (
	"fmt"
	"log"
	"os"

	"github.com/computers33333/airaccidentdata/docs"
	"github.com/joho/godotenv"
)

// AppConfig represents the application's configuration.
type AppConfig struct {
	DataSourceName   string // Database connection string
	Environment      string // Application environment (e.g., "development", "production")
	ServerAddress    string // Address on which the server should listen
	SwaggerHost      string // Host for Swagger documentation
	PageURL          string // URL to fetch the FAA accident data CSV file
	CSVFilePath      string // Path to save the downloaded FAA accident data CSV file
	GoogleMapsAPIKey string // API Key for Google Maps
}

// NewConfig initializes and returns a new AppConfig with default values obtained from environment variables.
func NewConfig() *AppConfig {
	// Load environment variables
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Warning: Failed to load environment variables: %v", err)
	}

	config := &AppConfig{
		DataSourceName:   GetDataSourceName(),
		Environment:      GetEnv("GO_ENV", "development"),
		ServerAddress:    GetEnv("SERVER_ADDRESS", "0.0.0.0:8080"),
		SwaggerHost:      GetDefaultSwaggerHost(GetEnv("GO_ENV", "development")),
		PageURL:          "https://www.asias.faa.gov/apex/f?p=100:93:::NO:::",
		CSVFilePath:      "downloaded_file.csv",
		GoogleMapsAPIKey: GetEnv("GOOGLE_MAPS_API_KEY", ""),
	}

	// Configure Swagger host
	docs.SwaggerInfo.Host = config.SwaggerHost

	return config
}

// GetEnv retrieves the value of an environment variable or returns a fallback value if the variable is not set.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// GetDefaultSwaggerHost returns the default Swagger host based on the environment.
func GetDefaultSwaggerHost(env string) string {
	switch env {
	case "development":
		return "localhost:8080"
	default:
		return "airaccidentdata.com"
	}
}

// GetDataSourceName constructs the MySQL Data Source Name (DSN) from individual environment variables.
func GetDataSourceName() string {
	user := GetEnv("MYSQL_USER", "user")
	password := GetEnv("MYSQL_PASSWORD", "password")
	host := GetEnv("MYSQL_HOST", "mysql")
	port := GetEnv("MYSQL_PORT", "3306")
	database := GetEnv("MYSQL_DATABASE", "airaccidentdata")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
}
