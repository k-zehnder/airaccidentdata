package models

import "time"

type AircraftAccident struct {
	ID                        int       `csv:"id" json:"id"`
	Updated                   string    `csv:"updated" json:"updated"`
	EntryDate                 time.Time `csv:"entry_date" json:"entryDate"`
	EventLocalDate            time.Time `csv:"event_local_date" json:"eventLocalDate"`
	EventLocalTime            string    `csv:"event_local_time" json:"eventLocalTime"`
	LocationCityName          string    `csv:"location_city_name" json:"locationCityName"`
	LocationStateName         string    `csv:"location_state_name" json:"locationStateName"`
	LocationCountryName       string    `csv:"location_country_name" json:"locationCountryName"`
	RemarkText                string    `csv:"remark_text" json:"remarkText"`
	EventTypeDescription      string    `csv:"event_type_description" json:"eventTypeDescription"`
	FSDODescription           string    `csv:"fsdo_description" json:"fsdoDescription"`
	RegistrationNumber        string    `csv:"registration_number" json:"registrationNumber"`
	FlightNumber              string    `csv:"flight_number" json:"flightNumber"`
	AircraftOperator          string    `csv:"aircraft_operator" json:"aircraftOperator"`
	AircraftMakeName          string    `csv:"aircraft_make_name" json:"aircraftMakeName"`
	AircraftModelName         string    `csv:"aircraft_model_name" json:"aircraftModelName"`
	AircraftMissingFlag       string    `csv:"aircraft_missing_flag" json:"aircraftMissingFlag"`
	AircraftDamageDescription string    `csv:"aircraft_damage_description" json:"aircraftDamageDescription"`
	FlightActivity            string    `csv:"flight_activity" json:"flightActivity"`
	FlightPhase               string    `csv:"flight_phase" json:"flightPhase"`
	FARPart                   string    `csv:"far_part" json:"farPart"`
	MaxInjuryLevel            string    `csv:"max_injury_level" json:"maxInjuryLevel"`
	FatalFlag                 string    `csv:"fatal_flag" json:"fatalFlag"`
	FlightCrewInjuryNone      int       `csv:"flight_crew_injury_none" json:"flightCrewInjuryNone"`
	FlightCrewInjuryMinor     int       `csv:"flight_crew_injury_minor" json:"flightCrewInjuryMinor"`
	FlightCrewInjurySerious   int       `csv:"flight_crew_injury_serious" json:"flightCrewInjurySerious"`
	FlightCrewInjuryFatal     int       `csv:"flight_crew_injury_fatal" json:"flightCrewInjuryFatal"`
	FlightCrewInjuryUnknown   int       `csv:"flight_crew_injury_unknown" json:"flightCrewInjuryUnknown"`
	CabinCrewInjuryNone       int       `csv:"cabin_crew_injury_none" json:"cabinCrewInjuryNone"`
	CabinCrewInjuryMinor      int       `csv:"cabin_crew_injury_minor" json:"cabinCrewInjuryMinor"`
	CabinCrewInjurySerious    int       `csv:"cabin_crew_injury_serious" json:"cabinCrewInjurySerious"`
	CabinCrewInjuryFatal      int       `csv:"cabin_crew_injury_fatal" json:"cabinCrewInjuryFatal"`
	CabinCrewInjuryUnknown    int       `csv:"cabin_crew_injury_unknown" json:"cabinCrewInjuryUnknown"`
	PassengerInjuryNone       int       `csv:"passenger_injury_none" json:"passengerInjuryNone"`
	PassengerInjuryMinor      int       `csv:"passenger_injury_minor" json:"passengerInjuryMinor"`
	PassengerInjurySerious    int       `csv:"passenger_injury_serious" json:"passengerInjurySerious"`
	PassengerInjuryFatal      int       `csv:"passenger_injury_fatal" json:"passengerInjuryFatal"`
	PassengerInjuryUnknown    int       `csv:"passenger_injury_unknown" json:"passengerInjuryUnknown"`
	GroundInjuryNone          int       `csv:"ground_injury_none" json:"groundInjuryNone"`
	GroundInjuryMinor         int       `csv:"ground_injury_minor" json:"groundInjuryMinor"`
	GroundInjurySerious       int       `csv:"ground_injury_serious" json:"groundInjurySerious"`
	GroundInjuryFatal         int       `csv:"ground_injury_fatal" json:"groundInjuryFatal"`
	GroundInjuryUnknown       int       `csv:"ground_injury_unknown" json:"groundInjuryUnknown"`
}
