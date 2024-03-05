package models

import "time"

type Aircraft struct {
	ID                 int
	RegistrationNumber string
	AircraftMakeName   string
	AircraftModelName  string
	AircraftOperator   string
	Accidents          []*Accident
}

type Accident struct {
	ID                        int
	Updated                   string
	EntryDate                 time.Time
	EventLocalDate            time.Time
	EventLocalTime            string
	LocationCityName          string
	LocationStateName         string
	LocationCountryName       string
	RemarkText                string
	EventTypeDescription      string
	FSDODescription           string
	FlightNumber              string
	AircraftMissingFlag       string
	AircraftDamageDescription string
	FlightActivity            string
	FlightPhase               string
	FARPart                   string
	MaxInjuryLevel            string
	FatalFlag                 string
	FlightCrewInjuryNone      int
	FlightCrewInjuryMinor     int
	FlightCrewInjurySerious   int
	FlightCrewInjuryFatal     int
	FlightCrewInjuryUnknown   int
	CabinCrewInjuryNone       int
	CabinCrewInjuryMinor      int
	CabinCrewInjurySerious    int
	CabinCrewInjuryFatal      int
	CabinCrewInjuryUnknown    int
	PassengerInjuryNone       int
	PassengerInjuryMinor      int
	PassengerInjurySerious    int
	PassengerInjuryFatal      int
	PassengerInjuryUnknown    int
	GroundInjuryNone          int
	GroundInjuryMinor         int
	GroundInjurySerious       int
	GroundInjuryFatal         int
	GroundInjuryUnknown       int
	AircraftID                int
}
