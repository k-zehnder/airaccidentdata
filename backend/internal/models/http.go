package models

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Message string `json:"message"`
}

// AccidentResponse defines the structure for the paginated accident response.
// It includes the slice of accidents, total number of accidents, current page, and the limit (accidents per page).
type AccidentResponse struct {
	Accidents []AircraftAccident `json:"accidents"`
	Total     int                `json:"total"`
	Page      int                `json:"page"`
	Limit     int                `json:"limit"`
}

// AccidentResponse defines the structure for the response containing specific aircraft accidents.
type AccidentDetailResponse struct {
	ID                 int                `json:"id"`
	RegistrationNumber string             `json:"registration_number"`
	AircraftMakeName   string             `json:"aircraft_make_name"`
	AircraftModelName  string             `json:"aircraft_model_name"`
	AircraftOperator   string             `json:"aircraft_operator"`
	Accidents          []AircraftAccident `json:"accidents"`
}
