package models

import "time"

type Aircraft struct {
	ID                 int                 `json:"id"`
	RegistrationNumber string              `json:"registration_number"`
	AircraftMakeName   string              `json:"aircraft_make_name"`
	AircraftModelName  string              `json:"aircraft_model_name"`
	AircraftOperator   string              `json:"aircraft_operator"`
	Accidents          []*AircraftAccident `json:"accidents"`
}

type AircraftAccident struct {
	ID                        int       `json:"id"`
	Updated                   string    `json:"updated"`
	EntryDate                 time.Time `json:"entry_date"`
	EventLocalDate            time.Time `json:"event_local_date"`
	EventLocalTime            string    `json:"event_local_time"`
	LocationCityName          string    `json:"location_city_name"`
	LocationStateName         string    `json:"location_state_name"`
	LocationCountryName       string    `json:"location_country_name"`
	RemarkText                string    `json:"remark_text"`
	EventTypeDescription      string    `json:"event_type_description"`
	FSDODescription           string    `json:"fsdo_description"`
	FlightNumber              string    `json:"flight_number"`
	AircraftMissingFlag       string    `json:"aircraft_missing_flag"`
	AircraftDamageDescription string    `json:"aircraft_damage_description"`
	FlightActivity            string    `json:"flight_activity"`
	FlightPhase               string    `json:"flight_phase"`
	FARPart                   string    `json:"far_part"`
	MaxInjuryLevel            string    `json:"max_injury_level"`
	FatalFlag                 string    `json:"fatal_flag"`
	FlightCrewInjuryNone      int       `json:"flight_crew_injury_none"`
	FlightCrewInjuryMinor     int       `json:"flight_crew_injury_minor"`
	FlightCrewInjurySerious   int       `json:"flight_crew_injury_serious"`
	FlightCrewInjuryFatal     int       `json:"flight_crew_injury_fatal"`
	FlightCrewInjuryUnknown   int       `json:"flight_crew_injury_unknown"`
	CabinCrewInjuryNone       int       `json:"cabin_crew_injury_none"`
	CabinCrewInjuryMinor      int       `json:"cabin_crew_injury_minor"`
	CabinCrewInjurySerious    int       `json:"cabin_crew_injury_serious"`
	CabinCrewInjuryFatal      int       `json:"cabin_crew_injury_fatal"`
	CabinCrewInjuryUnknown    int       `json:"cabin_crew_injury_unknown"`
	PassengerInjuryNone       int       `json:"passenger_injury_none"`
	PassengerInjuryMinor      int       `json:"passenger_injury_minor"`
	PassengerInjurySerious    int       `json:"passenger_injury_serious"`
	PassengerInjuryFatal      int       `json:"passenger_injury_fatal"`
	PassengerInjuryUnknown    int       `json:"passenger_injury_unknown"`
	GroundInjuryNone          int       `json:"ground_injury_none"`
	GroundInjuryMinor         int       `json:"ground_injury_minor"`
	GroundInjurySerious       int       `json:"ground_injury_serious"`
	GroundInjuryFatal         int       `json:"ground_injury_fatal"`
	GroundInjuryUnknown       int       `json:"ground_injury_unknown"`
	AircraftID                int       `json:"aircraft_id"`
}
