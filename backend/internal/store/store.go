package store

import (
	"database/sql"
	"fmt"

	"github.com/computers33333/airaccidentdata/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

// StoreInterface defines the methods that our store implementations must have.
type StoreInterface interface {
	GetAccidents(page, limit int) ([]*models.AircraftAccident, int, error)
	GetAccidentsByRegistration(registrationNumber string) ([]*models.AircraftAccident, error)
	GetAircraftById(id int) ([]*models.Aircraft, int, error)
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

// GetAccidents fetches a specific page of aircraft accidents from the database.
func (s *Store) GetAccidents(page, limit int) ([]*models.AircraftAccident, int, error) {
	var incidents []*models.AircraftAccident

	offset := (page - 1) * limit
	query := `
		SELECT 
			id, updated, entry_date, event_local_date, event_local_time,
			location_city_name, location_state_name, location_country_name, 
			remark_text, event_type_description, fsdo_description,
			flight_number, aircraft_missing_flag, aircraft_damage_description,
			flight_activity, flight_phase, far_part, max_injury_level, fatal_flag, aircraft_id
		FROM Accidents
		ORDER BY id LIMIT ? OFFSET ?;
	`

	// Fetch the accidents from the database
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("query execution error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var incident models.AircraftAccident
		// Scan each column in the row into the corresponding field of the AircraftIncident struct
		if err := rows.Scan(
			&incident.ID, &incident.Updated, &incident.EntryDate, &incident.EventLocalDate,
			&incident.EventLocalTime, &incident.LocationCityName, &incident.LocationStateName,
			&incident.LocationCountryName, &incident.RemarkText, &incident.EventTypeDescription,
			&incident.FSDODescription, &incident.FlightNumber, &incident.AircraftMissingFlag,
			&incident.AircraftDamageDescription, &incident.FlightActivity, &incident.FlightPhase,
			&incident.FARPart, &incident.MaxInjuryLevel, &incident.FatalFlag, &incident.AircraftID,
		); err != nil {
			return nil, 0, err
		}

		incidents = append(incidents, &incident)
	}

	// Handle any iteration errors.
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	// Fetch total count of incidents for pagination
	var totalCount int
	countQuery := "SELECT COUNT(*) FROM Accidents;"
	err = s.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("count query error: %w", err)
	}

	return incidents, totalCount, nil
}

// GetAccidentByRegistrationHandler fetches an aircraft and its accidents by registration number
func (s *Store) GetAccidentByRegistrationHandler(registrationNumber string) (*models.Aircraft, error) {
	// Fetch Aircraft
	aircraftQuery := `SELECT id, registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator FROM Aircrafts WHERE registration_number = ?`
	var aircraft models.Aircraft
	err := s.db.QueryRow(aircraftQuery, registrationNumber).Scan(&aircraft.ID, &aircraft.RegistrationNumber, &aircraft.AircraftMakeName, &aircraft.AircraftModelName, &aircraft.AircraftOperator)
	if err != nil {
		return nil, fmt.Errorf("error fetching aircraft: %w", err)
	}

	// Fetch Accidents
	accidentsQuery := `SELECT * FROM Accidents WHERE aircraft_id = ?`
	rows, err := s.db.Query(accidentsQuery, aircraft.ID)
	if err != nil {
		return nil, fmt.Errorf("error fetching accidents: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var accident models.AircraftAccident
		if err := rows.Scan(
			&accident.ID, &accident.Updated, &accident.EntryDate, &accident.EventLocalDate,
			&accident.EventLocalTime, &accident.LocationCityName, &accident.LocationStateName,
			&accident.LocationCountryName, &accident.RemarkText, &accident.EventTypeDescription,
			&accident.FSDODescription, &accident.FlightNumber, &accident.AircraftMissingFlag,
			&accident.AircraftDamageDescription, &accident.FlightActivity, &accident.FlightPhase,
			&accident.FARPart, &accident.MaxInjuryLevel, &accident.FatalFlag,
			&accident.FlightCrewInjuryNone, &accident.FlightCrewInjuryMinor,
			&accident.FlightCrewInjurySerious, &accident.FlightCrewInjuryFatal,
			&accident.FlightCrewInjuryUnknown, &accident.CabinCrewInjuryNone,
			&accident.CabinCrewInjuryMinor, &accident.CabinCrewInjurySerious,
			&accident.CabinCrewInjuryFatal, &accident.CabinCrewInjuryUnknown,
			&accident.PassengerInjuryNone, &accident.PassengerInjuryMinor,
			&accident.PassengerInjurySerious, &accident.PassengerInjuryFatal,
			&accident.PassengerInjuryUnknown, &accident.GroundInjuryNone,
			&accident.GroundInjuryMinor, &accident.GroundInjurySerious,
			&accident.GroundInjuryFatal, &accident.GroundInjuryUnknown,
			&accident.AircraftID,
		); err != nil {
			return nil, fmt.Errorf("error scanning accident row: %w", err)
		}
		aircraft.Accidents = append(aircraft.Accidents, &accident)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iteration error: %w", err)
	}

	return &aircraft, nil
}

// GetAllAircrafts fetches all aircraft from the database.
func (s *Store) GetAllAircrafts() ([]*models.Aircraft, error) {
	// Query to fetch all aircraft.
	query := `SELECT id, registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator FROM Aircrafts`

	// Perform the database query.
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching aircrafts: %w", err)
	}
	defer rows.Close()

	// Create a slice to store retrieved aircraft.
	var aircrafts []*models.Aircraft

	// Iterate over the rows and scan the values into aircraft struct.
	for rows.Next() {
		var aircraft models.Aircraft
		err := rows.Scan(
			&aircraft.ID,
			&aircraft.RegistrationNumber,
			&aircraft.AircraftMakeName,
			&aircraft.AircraftModelName,
			&aircraft.AircraftOperator,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning aircraft row: %w", err)
		}
		aircrafts = append(aircrafts, &aircraft)
	}

	// Check for any iteration errors.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iteration error: %w", err)
	}

	return aircrafts, nil
}

// GetAircraftById fetches an aircraft by its ID from the database.
func (s *Store) GetAircraftById(id int) (*models.Aircraft, error) {
	// Query to fetch the aircraft by ID.
	query := `SELECT id, registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator FROM Aircrafts WHERE id = ?`

	// Perform the database query.
	row := s.db.QueryRow(query, id)

	// Create a variable to hold the scanned aircraft.
	var aircraft models.Aircraft

	// Scan the values from the row into the aircraft struct.
	err := row.Scan(
		&aircraft.ID,
		&aircraft.RegistrationNumber,
		&aircraft.AircraftMakeName,
		&aircraft.AircraftModelName,
		&aircraft.AircraftOperator,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return nil for the aircraft if not found.
			return nil, nil
		}
		// Return the error if any other error occurred.
		return nil, fmt.Errorf("error scanning aircraft row: %w", err)
	}

	// Return the scanned aircraft.
	return &aircraft, nil
}

// GetAccidentById fetches an accident by its ID from the database.
func (s *Store) GetAccidentById(id int) (*models.AircraftAccident, error) {
	// Query to fetch the accident by ID.
	query := `
		SELECT 
			id, updated, entry_date, event_local_date, event_local_time,
			location_city_name, location_state_name, location_country_name, 
			remark_text, event_type_description, fsdo_description,
			flight_number, aircraft_missing_flag, aircraft_damage_description,
			flight_activity, flight_phase, far_part, max_injury_level, fatal_flag, aircraft_id
		FROM Accidents
		WHERE id = ?;
	`

	// Perform the database query.
	row := s.db.QueryRow(query, id)

	// Create a variable to hold the scanned accident.
	var accident models.AircraftAccident

	// Scan the values from the row into the accident struct.
	err := row.Scan(
		&accident.ID, &accident.Updated, &accident.EntryDate, &accident.EventLocalDate,
		&accident.EventLocalTime, &accident.LocationCityName, &accident.LocationStateName,
		&accident.LocationCountryName, &accident.RemarkText, &accident.EventTypeDescription,
		&accident.FSDODescription, &accident.FlightNumber, &accident.AircraftMissingFlag,
		&accident.AircraftDamageDescription, &accident.FlightActivity, &accident.FlightPhase,
		&accident.FARPart, &accident.MaxInjuryLevel, &accident.FatalFlag, &accident.AircraftID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return nil for the accident if not found.
			return nil, nil
		}
		// Return the error if any other error occurred.
		return nil, fmt.Errorf("error scanning accident row: %w", err)
	}

	// Return the scanned accident.
	return &accident, nil
}
