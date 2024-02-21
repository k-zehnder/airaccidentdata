package models

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Message string `json:"message"`
}

// AccidentResponse defines the structure for the paginated accident response.
// It includes the slice of accidents, total number of accidents, current page, and the limit (accidents per page).
type AccidentResponse struct {
	Accidents []AircraftAccident `json:"accidents"`
	Total     int                        `json:"total"`
	Page      int                        `json:"page"`
	Limit     int                        `json:"limit"`
}
