// Package shared provides shared utility functions for database operations.
package shared

import (
	"database/sql"
	"fmt"
	"os"
)

// setupDatabase establishes a connection to the MySQL database.
func SetupDatabase() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database is not reachable: %w", err)
	}

	return db, nil
}
