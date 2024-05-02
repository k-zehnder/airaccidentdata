// Package shared provides shared utility functions for database operations.
package shared

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// loadEnv searches for the .env file starting in the current directory and moving up.
func LoadEnv() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".env")); err == nil {
			return godotenv.Load(filepath.Join(dir, ".env"))
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return fmt.Errorf("root directory reached, .env file not found")
		}
		dir = parentDir
	}
}
