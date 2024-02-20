package store

import (
	"database/sql"
	"fmt"

	// Import the MySQL driver with a blank identifier to ensure its `init()` function is executed.
	"github.com/computers33333/airaccidentdata/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

// StoreInterface defines the methods that our store implementations must have.
type StoreInterface interface {
	GetAccidents([]*models.AircraftAccident, error)
}

// Store satisfies the StoreInterface.
type Store struct {
	db *sql.DB // db represents a connection to a MySQL database
}

// NewStore establishes a new MYSQL database connection.
func NewStore(dataSourceName string) (*Store, error) {
	// Attempt to open a database connection.
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Ping the database to ensure the connection is active and the server is reachable.
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Return a new Store instance with the established database connection.
	return &Store{db: db}, nil
}

// Method on the Store type that returns a slice of pointers and an error.
// This method can modify state because it has a pointer receiver.
func (s *Store) GetAccidents() ([]*models.AircraftAccident, error) {
	var incidents []*models.AircraftAccident

	// Query to fetch all incidents from the database
	rows, err := s.db.Query("SELECT * FROM AircraftAccidents ORDER BY id DESC LIMIT 30;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var incident models.AircraftAccident

		// Scan each column in the row into the corresponding field of the AircraftIncident struct
		if err := rows.Scan(
			&incident.ID, &incident.Updated, &incident.EntryDate, &incident.EventLocalDate,
			&incident.EventLocalTime, &incident.LocationCityName, &incident.LocationStateName,
			&incident.LocationCountryName, &incident.RemarkText, &incident.EventTypeDescription,
			&incident.FSDODescription, &incident.RegistrationNumber, &incident.FlightNumber,
			&incident.AircraftOperator, &incident.AircraftMakeName, &incident.AircraftModelName,
			&incident.AircraftMissingFlag, &incident.AircraftDamageDescription, &incident.FlightActivity,
			&incident.FlightPhase, &incident.FARPart, &incident.MaxInjuryLevel, &incident.FatalFlag,
			&incident.FlightCrewInjuryNone, &incident.FlightCrewInjuryMinor, &incident.FlightCrewInjurySerious,
			&incident.FlightCrewInjuryFatal, &incident.FlightCrewInjuryUnknown, &incident.CabinCrewInjuryNone,
			&incident.CabinCrewInjuryMinor, &incident.CabinCrewInjurySerious, &incident.CabinCrewInjuryFatal,
			&incident.CabinCrewInjuryUnknown, &incident.PassengerInjuryNone, &incident.PassengerInjuryMinor,
			&incident.PassengerInjurySerious, &incident.PassengerInjuryFatal, &incident.PassengerInjuryUnknown,
			&incident.GroundInjuryNone, &incident.GroundInjuryMinor, &incident.GroundInjurySerious,
			&incident.GroundInjuryFatal, &incident.GroundInjuryUnknown,
		); err != nil {
			return nil, err
		}

		incidents = append(incidents, &incident)
	}

	// Handle any iteration errors.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Return a slice of AircraftIncident pointers and an error.
	return incidents, nil
}