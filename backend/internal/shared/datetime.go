// Package shared provides shared utility functions for database operations.
package shared

import (
	"fmt"
	"time"
)

// Helper function to parse a date string into time.Time, returns time.Time and error.
func ParseDate(dateStr string) (time.Time, error) {
	layout := "02-Jan-06"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing date '%s': %v", dateStr, err)
	}
	return t, nil
}

// Helper function to format a time string into time.Time, returns time.Time and error.
func ParseTime(timeStr string) (string, error) {
	layout := "15:04:05Z"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return "", fmt.Errorf("error parsing time '%s': %v", timeStr, err)
	}
	return t.Format("15:04:05"), nil
}
