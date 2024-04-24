package models

import "time"

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Message string `json:"message"`
}

// AccidentResponse defines the structure for the paginated accident response.
type AccidentPaginatedResponse struct {
	Accidents []AircraftAccident `json:"accidents"`
	Total     int                `json:"total"`
	Page      int                `json:"page"`
	Limit     int                `json:"limit"`
}

// AircraftPaginatedResponse defines the structure for the paginated aircraft response.
type AircraftPaginatedResponse struct {
	Aircrafts []AircraftResponse `json:"aircrafts"`
	Total     int                `json:"total"`
	Page      int                `json:"page"`
	Limit     int                `json:"limit"`
}

// AircraftResponse defines the structure for the response containing aircraft details.
type AircraftResponse struct {
	ID                 int    `json:"id"`
	RegistrationNumber string `json:"registration_number"`
	AircraftMakeName   string `json:"aircraft_make_name"`
	AircraftModelName  string `json:"aircraft_model_name"`
	AircraftOperator   string `json:"aircraft_operator"`
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

// AircraftAccident represents details of an aviation accident.
type AircraftAccidentResponse struct {
	ID                      int       `json:"id"`
	Updated                 string    `json:"updated"`
	EntryDate               time.Time `json:"entry_date"`
	EventLocalDate          time.Time `json:"event_local_date"`
	EventLocalTime          string    `json:"event_local_time"`
	LocationCityName        string    `json:"location_city_name"`
	LocationStateName       string    `json:"location_state_name"`
	LocationCountryName     string    `json:"location_country_name"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	RemarkText              string    `json:"remark_text"`
	EventTypeDescription    string    `json:"event_type_description"`
	FSDODescription         string    `json:"fsdo_description"`
	FlightNumber            string    `json:"flight_number"`
	AircraftMissingFlag     string    `json:"aircraft_missing_flag"`
	AircraftDamageDesc      string    `json:"aircraft_damage_description"`
	FlightActivity          string    `json:"flight_activity"`
	FlightPhase             string    `json:"flight_phase"`
	FARPart                 string    `json:"far_part"`
	MaxInjuryLevel          string    `json:"max_injury_level"`
	FatalFlag               string    `json:"fatal_flag"`
	FlightCrewInjuryNone    int       `json:"flight_crew_injury_none"`
	FlightCrewInjuryMinor   int       `json:"flight_crew_injury_minor"`
	FlightCrewInjurySerious int       `json:"flight_crew_injury_serious"`
	FlightCrewInjuryFatal   int       `json:"flight_crew_injury_fatal"`
	FlightCrewInjuryUnknown int       `json:"flight_crew_injury_unknown"`
	CabinCrewInjuryNone     int       `json:"cabin_crew_injury_none"`
	CabinCrewInjuryMinor    int       `json:"cabin_crew_injury_minor"`
	CabinCrewInjurySerious  int       `json:"cabin_crew_injury_serious"`
	CabinCrewInjuryFatal    int       `json:"cabin_crew_injury_fatal"`
	CabinCrewInjuryUnknown  int       `json:"cabin_crew_injury_unknown"`
	PassengerInjuryNone     int       `json:"passenger_injury_none"`
	PassengerInjuryMinor    int       `json:"passenger_injury_minor"`
	PassengerInjurySerious  int       `json:"passenger_injury_serious"`
	PassengerInjuryFatal    int       `json:"passenger_injury_fatal"`
	PassengerInjuryUnknown  int       `json:"passenger_injury_unknown"`
	GroundInjuryNone        int       `json:"ground_injury_none"`
	GroundInjuryMinor       int       `json:"ground_injury_minor"`
	GroundInjurySerious     int       `json:"ground_injury_serious"`
	GroundInjuryFatal       int       `json:"ground_injury_fatal"`
	GroundInjuryUnknown     int       `json:"ground_injury_unknown"`
	AircraftID              int       `json:"aircraft_id"`
}

// ImageResponse represents the response format for one image associated with an aircraft.
type ImageResponse struct {
	ID         int    `json:"id"`
	AircraftID int    `json:"aircraft_id"`
	ImageURL   string `json:"image_url"`
	S3URL      string `json:"s3_url"`
}

// ImagesForAircraftResponse represents the response format for all images associated with an aircraft.
type ImagesForAircraftResponse struct {
	AircraftID int             `json:"aircraft_id"`
	Images     []ImageResponse `json:"images"`
}

type GeoResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}
