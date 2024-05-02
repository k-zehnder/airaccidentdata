package store

import (
	"database/sql"
	"fmt"

	"github.com/computers33333/airaccidentdata/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

// StoreInterface defines the methods that our store implementations must have
type StoreInterface interface {
	GetAccidents(page, limit int) ([]*models.Accident, int, error)
	GetAccidentsByRegistration(registrationNumber string) ([]*models.Accident, error)
	GetAircraftById(id int) ([]*models.Aircraft, int, error)
	GetAircrafts(page, limit int) ([]*models.Aircraft, int, error)
	GetAllImagesForAircraft(aircraftID int) ([]*models.AircraftImage, error)
	GetImageForAircraft(aircraftID, imageID int) (*models.AircraftImage, error)
}

// Store satisfies the StoreInterface
type Store struct {
	db *sql.DB
}

// NewStore establishes a new MYSQL database connection
func NewStore(dataSourceName string) (*Store, error) {
	// Attempt to open a database connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Ping the database to ensure the connection is active and the server is reachable
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Return a new Store instance with the established database connection
	return &Store{db: db}, nil
}

// GetAircrafts fetches a specific page of aircrafts from the database with pagination
func (s *Store) GetAircrafts(page, limit int) ([]*models.Aircraft, int, error) {
	// Calculate the offset based on the page number and limit
	offset := (page - 1) * limit

	// Query to fetch a specific page of aircrafts
	query := `
		SELECT id, registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator
		FROM Aircrafts
		LIMIT ? OFFSET ? 
	`

	// Perform the database query
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching aircrafts: %w", err)
	}
	defer rows.Close()

	// Create a slice to store retrieved aircraft
	var aircrafts []*models.Aircraft

	// Iterate over the rows and scan the values into aircraft struct
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
			return nil, 0, fmt.Errorf("error scanning aircraft row: %w", err)
		}
		aircrafts = append(aircrafts, &aircraft)
	}

	// Query to fetch the total count of aircrafts
	countQuery := `SELECT COUNT(*) FROM Aircrafts`
	var totalCount int
	err = s.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching total number of accidents: %w", err)
	}

	// Return the retrieved aircrafts and total count
	return aircrafts, totalCount, nil

}

// GetAircraftById fetches an aircraft by its ID from the database
func (s *Store) GetAircraftById(id int) (*models.Aircraft, error) {
	// Query to fetch the aircraft by ID
	query := `SELECT id, registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator FROM Aircrafts WHERE id = ?`

	// Perform the database query
	row := s.db.QueryRow(query, id)

	// Create a variable to hold the scanned aircraft
	var aircraft models.Aircraft

	// Scan the values from the row into the aircraft struct
	err := row.Scan(
		&aircraft.ID,
		&aircraft.RegistrationNumber,
		&aircraft.AircraftMakeName,
		&aircraft.AircraftModelName,
		&aircraft.AircraftOperator,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return nil for the aircraft if not found
			return nil, nil
		}
		// Return the error if any other error occurred
		return nil, fmt.Errorf("error scanning aircraft row: %w", err)
	}

	// Return the scanned aircraft
	return &aircraft, nil
}

// GetAccidents fetches a specific page of aircraft accidents from the database
func (s *Store) GetAccidents(page, limit int) ([]*models.Accident, int, error) {
	var accidents []*models.Accident
	offset := (page - 1) * limit
	query := `
		SELECT 
			id, updated, entry_date, event_local_date, event_local_time,
			remark_text, event_type_description, fsdo_description, flight_number, 
			aircraft_missing_flag, aircraft_damage_description, flight_activity, flight_phase, 
			far_part, fatal_flag, location_id, aircraft_id
		FROM Accidents
		ORDER BY id LIMIT ? OFFSET ?;
	`

	// Execute the query with pagination
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("query execution error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var accident models.Accident
		// Scan each column in the row into the corresponding field of the Accident struct
		if err := rows.Scan(
			&accident.ID, &accident.Updated, &accident.EntryDate, &accident.EventLocalDate,
			&accident.EventLocalTime, &accident.RemarkText,
			&accident.EventTypeDescription, &accident.FSDODescription, &accident.FlightNumber,
			&accident.AircraftMissingFlag, &accident.AircraftDamageDescription, &accident.FlightActivity,
			&accident.FlightPhase, &accident.FARPart, &accident.FatalFlag,
			&accident.LocationID, &accident.AircraftID,
		); err != nil {
			return nil, 0, err
		}

		accidents = append(accidents, &accident)
	}

	// Handle any iteration errors
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	// Fetch total count of accidents for pagination
	var totalCount int
	countQuery := "SELECT COUNT(*) FROM Accidents;"
	err = s.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("count query error: %w", err)
	}

	return accidents, totalCount, nil
}

// GetAccidentById fetches an accident by its ID from the database
func (s *Store) GetAccidentById(id int) (*models.Aircraft, error) {
	// Query to fetch the aircraft by ID
	query := `SELECT id, registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator FROM Aircrafts WHERE id = ?`

	// Perform the database query
	row := s.db.QueryRow(query, id)

	// Create a variable to hold the scanned aircraft
	var aircraft models.Aircraft

	// Scan the values from the row into the aircraft struct
	err := row.Scan(
		&aircraft.ID,
		&aircraft.RegistrationNumber,
		&aircraft.AircraftMakeName,
		&aircraft.AircraftModelName,
		&aircraft.AircraftOperator,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return nil for the aircraft if not found
			return nil, nil
		}
		// Return the error if any other error occurred
		return nil, fmt.Errorf("error scanning aircraft row: %w", err)
	}

	// Return the scanned aircraft
	return &aircraft, nil
}

// GetAllImagesForAircraft fetches all images associated with an aircraft by its ID
func (s *Store) GetAllImagesForAircraft(aircraftID int) ([]*models.AircraftImage, error) {
	// Query to fetch all images associated with the aircraft.
	query := `SELECT id, aircraft_id, image_url, s3_url FROM AircraftImages WHERE aircraft_id = ?`

	// Perform the database query
	rows, err := s.db.Query(query, aircraftID)
	if err != nil {
		return nil, fmt.Errorf("error fetching aircraft images: %w", err)
	}
	defer rows.Close()

	// Create a slice to store retrieved images
	var images []*models.AircraftImage

	// Iterate over the rows and scan the image details into the slice
	for rows.Next() {
		var image models.AircraftImage
		err := rows.Scan(&image.ID, &image.AircraftID, &image.ImageURL, &image.S3URL)
		if err != nil {
			return nil, fmt.Errorf("error scanning image details: %w", err)
		}
		images = append(images, &image)
	}

	// Check for any iteration errors.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iteration error: %w", err)
	}

	return images, nil
}

// GetImageForAircraft fetches a specific image associated with an aircraft by its ID
func (s *Store) GetImageForAircraft(aircraftID, imageID int) (*models.AircraftImage, error) {
	query := `SELECT id, aircraft_id, image_url, s3_url FROM AircraftImages WHERE aircraft_id = ? AND id = ?`

	row := s.db.QueryRow(query, aircraftID, imageID)

	var image models.AircraftImage
	err := row.Scan(&image.ID, &image.AircraftID, &image.ImageURL, &image.S3URL)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return nil for the image if not found
			return nil, nil
		}
		return nil, fmt.Errorf("error scanning image details: %w", err)
	}

	return &image, nil
}

// GetInjuriesByAccidentIdHandler fetches injuries about an accident by its ID
func (s *Store) GetInjuriesByAccidentIdHandler(aircraftId int) (*models.Injury, error) {
	query := `SELECT id, person_type, injury_severity, count, accident_id FROM Injuries WHERE accident_id = ?`

	row := s.db.QueryRow(query, aircraftId)

	var injury models.Injury
	err := row.Scan(&injury.ID, &injury.PersonType, &injury.InjurySeverity, &injury.Count, &injury.AccidentID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return nil for the injury if not found
			return nil, nil
		}
		return nil, fmt.Errorf("error scanning injury details: %w", err)

	}

	return &injury, nil
}
