// Package shared provides shared utility functions for database operations.
package shared

import (
	"log"
	"strconv"
)

// atoiSafe converts string to int, returns 0 if conversion fails or the string is empty.
func AtoiSafe(s string) int {
	if s == "" {
		return 0
	}
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		return 0
	}
	return value
}
