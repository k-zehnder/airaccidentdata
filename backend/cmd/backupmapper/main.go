package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Pull database details from environment variables
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	// Load the CA certificate
	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile("./ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}

	// Register the custom TLS config
	err = mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs:    rootCertPool,
		MinVersion: tls.VersionTLS12,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=custom", user, pass, host, port, database)

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Execute a query
	rows, err := db.Query("SELECT VERSION()")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var version string
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("MySQL version: %s\n", version)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// loadEnv searches for the .env file starting in the current directory and moving up.
func loadEnv() error {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		// Check if .env exists in this directory.
		if _, err := os.Stat(filepath.Join(dir, ".env")); err == nil {
			// Load the .env file.
			return godotenv.Load(filepath.Join(dir, ".env"))
		}

		// Move up to the parent directory.
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// Root of the filesystem reached, .env not found
			return fmt.Errorf("root directory reached, .env file not found")
		}
		dir = parentDir
	}
}
