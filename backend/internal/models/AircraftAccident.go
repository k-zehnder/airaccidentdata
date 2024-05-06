// Package models defines data structures for aircraft, accidents, and related entities,
// supporting data storage and manipulation across the application.
package models

import (
	"time"
)

type Aircraft struct {
	ID                 int    `json:"id"`
	RegistrationNumber string `json:"registration_number"`
	AircraftMakeName   string `json:"aircraft_make_name"`
	AircraftModelName  string `json:"aircraft_model_name"`
	AircraftOperator   string `json:"aircraft_operator"`
}

type Location struct {
	ID          int     `json:"id"`
	CityName    string  `json:"city_name"`
	StateName   string  `json:"state_name"`
	CountryName string  `json:"country_name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type Accident struct {
	ID                        int       `json:"id"`
	Updated                   string    `json:"updated"`
	EntryDate                 time.Time `json:"entry_date"`
	EventLocalDate            time.Time `json:"event_local_date"`
	EventLocalTime            string    `json:"event_local_time"`
	RemarkText                string    `json:"remark_text"`
	EventTypeDescription      string    `json:"event_type_description"`
	FSDODescription           string    `json:"fsdo_description"`
	FlightNumber              string    `json:"flight_number"`
	AircraftMissingFlag       string    `json:"aircraft_missing_flag"`
	AircraftDamageDescription string    `json:"aircraft_damage_description"`
	FlightActivity            string    `json:"flight_activity"`
	FlightPhase               string    `json:"flight_phase"`
	FARPart                   string    `json:"far_part"`
	FatalFlag                 string    `json:"fatal_flag"`
	LocationID                int       `json:"location_id"`
	AircraftID                int       `json:"aircraft_id"`
}

type Injury struct {
	ID             int    `json:"id"`
	PersonType     string `json:"person_type"`
	InjurySeverity string `json:"injury_severity"`
	Count          int    `json:"count"`
	AccidentID     int    `json:"accident_id"`
}

type AircraftImage struct {
	ID         int    `json:"id"`
	AircraftID int    `json:"aircraft_id"`
	ImageURL   string `json:"image_url"`
	S3URL      string `json:"s3_url"`
}

type AircraftPaginatedResponse struct {
	Aircrafts []Aircraft `json:"aircrafts"`
	Total     int        `json:"total"`
	Page      int        `json:"page"`
	Limit     int        `json:"limit"`
}

type AccidentPaginatedResponse struct {
	Accidents []Accident `json:"accidents"`
	Total     int        `json:"total"`
	Page      int        `json:"page"`
	Limit     int        `json:"limit"`
}

type ImagesForAircraftResponse struct {
	AircraftID int             `json:"aircraft_id"`
	Images     []AircraftImage `json:"images"`
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

type ErrorResponse struct {
	Message string `json:"message"`
}
