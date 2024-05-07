// Package store implements the primary data store interactions for accidents and aircraft.
package store

import (
	"database/sql"
	"fmt"

	"github.com/computers33333/airaccidentdata/internal/models"
	_ "github.com/go-sql-driver/mysql" // Blank identifier imports MySQL driver to initialize and register it.
)

// StoreInterface defines the methods that our store implementations must have.
type StoreInterface interface {
	GetAccidents(page, limit int) ([]*models.Accident, int, error)
	GetAccidentsByRegistration(registrationNumber string) ([]*models.Accident, error)
	GetAircraftById(id int) ([]*models.Aircraft, int, error)
	GetAircrafts(page, limit int) ([]*models.Aircraft, int, error)
	GetAllImagesForAircraft(aircraftID int) ([]*models.AircraftImage, error)
	GetImageForAircraft(aircraftID, imageID int) (*models.AircraftImage, error)
}

// Store satisfies the StoreInterface.
type Store struct {
	db *sql.DB
}

// NewStore establishes a new MYSQL database connection.
func NewStore(dataSourceName string) (*Store, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Store{db: db}, nil
}

// GetAircrafts fetches a specific page of aircrafts from the database with pagination.
func (s *Store) GetAircrafts(page, limit int) ([]*models.Aircraft, int, error) {
	offset := (page - 1) * limit

	query := `
		SELECT id, registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator
		FROM Aircrafts
		LIMIT ? OFFSET ? 
	`

	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching aircrafts: %w", err)
	}
	defer rows.Close()

	var aircrafts []*models.Aircraft
	for rows.Next() {
		var aircraft models.Aircraft
		if err := rows.Scan(
			&aircraft.ID,
			&aircraft.RegistrationNumber,
			&aircraft.AircraftMakeName,
			&aircraft.AircraftModelName,
			&aircraft.AircraftOperator,
		); err != nil {
			return nil, 0, fmt.Errorf("error scanning aircraft row: %w", err)
		}
		aircrafts = append(aircrafts, &aircraft)
	}

	countQuery := `SELECT COUNT(*) FROM Aircrafts`
	var totalCount int
	err = s.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching total number of accidents: %w", err)
	}

	return aircrafts, totalCount, nil
}

// GetAircraftById fetches an aircraft by its ID from the database.
func (s *Store) GetAircraftById(id int) (*models.Aircraft, error) {
	query := `SELECT id, registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator FROM Aircrafts WHERE id = ?`

	row := s.db.QueryRow(query, id)

	var aircraft models.Aircraft
	err := row.Scan(
		&aircraft.ID,
		&aircraft.RegistrationNumber,
		&aircraft.AircraftMakeName,
		&aircraft.AircraftModelName,
		&aircraft.AircraftOperator,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error scanning aircraft row: %w", err)
	}

	return &aircraft, nil
}

// GetAccidents fetches a specific page of aircraft accidents from the database.
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

	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("query execution error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var accident models.Accident
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

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	var totalCount int
	countQuery := "SELECT COUNT(*) FROM Accidents;"
	err = s.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("count query error: %w", err)
	}

	return accidents, totalCount, nil
}

// GetAccidentById fetches an accident by its ID from the database.
func (s *Store) GetAccidentById(id int) (*models.Accident, error) {
	query := `
		SELECT 
			id, updated, entry_date, event_local_date, event_local_time, remark_text, event_type_description, fsdo_description,
			flight_number, aircraft_missing_flag, aircraft_damage_description,
			flight_activity, flight_phase, far_part, fatal_flag, aircraft_id
		FROM Accidents
		WHERE id = ?;
	`

	row := s.db.QueryRow(query, id)

	var accident models.Accident
	err := row.Scan(
		&accident.ID, &accident.Updated, &accident.EntryDate, &accident.EventLocalDate,
		&accident.EventLocalTime, &accident.RemarkText, &accident.EventTypeDescription,
		&accident.FSDODescription, &accident.FlightNumber, &accident.AircraftMissingFlag,
		&accident.AircraftDamageDescription, &accident.FlightActivity, &accident.FlightPhase,
		&accident.FARPart, &accident.FatalFlag, &accident.AircraftID,
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

// GetLocationByAccidentId retrieves location details based on an accident's location ID.
func (s *Store) GetLocationByAccidentId(accidentId int) (*models.Location, error) {
	query := `
    SELECT Locations.id, Locations.city_name, Locations.state_name, Locations.country_name, Locations.latitude, Locations.longitude
    FROM Locations
    JOIN Accidents ON Locations.id = Accidents.location_id
    WHERE Accidents.id = ?;
    `
	row := s.db.QueryRow(query, accidentId)

	var location models.Location
	err := row.Scan(&location.ID, &location.CityName, &location.StateName, &location.CountryName, &location.Latitude, &location.Longitude)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching location: %w", err)
	}
	return &location, nil
}

// GetAllImagesForAircraft fetches all images associated with an aircraft by its ID.
func (s *Store) GetAllImagesForAircraft(aircraftID int) ([]*models.AircraftImage, error) {
	query := `SELECT id, aircraft_id, image_url, COALESCE(s3_url, '') AS s3_url FROM AircraftImages WHERE aircraft_id = ?`
	rows, err := s.db.Query(query, aircraftID)
	if err != nil {
		return nil, fmt.Errorf("error fetching aircraft images: %w", err)
	}
	defer rows.Close()

	var images []*models.AircraftImage
	for rows.Next() {
		var image models.AircraftImage
		err := rows.Scan(&image.ID, &image.AircraftID, &image.ImageURL, &image.S3URL)
		if err != nil {
			return nil, fmt.Errorf("error scanning image details: %w", err)
		}
		images = append(images, &image)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iteration error: %w", err)
	}

	return images, nil
}

// GetImageForAircraft fetches a specific image associated with an aircraft by its ID.
func (s *Store) GetImageForAircraft(aircraftID, imageID int) (*models.AircraftImage, error) {
	query := `SELECT id, aircraft_id, image_url, s3_url FROM AircraftImages WHERE aircraft_id = ? AND id = ?`

	row := s.db.QueryRow(query, aircraftID, imageID)

	var image models.AircraftImage
	err := row.Scan(&image.ID, &image.AircraftID, &image.ImageURL, &image.S3URL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error scanning image details: %w", err)
	}

	return &image, nil
}

// GetInjuriesByAccidentIdHandler fetches all injuries associated with a specific accident by its ID.
func (s *Store) GetInjuriesByAccidentIdHandler(accidentId int) ([]*models.Injury, error) {
	query := `SELECT id, person_type, injury_severity, count, accident_id FROM Injuries WHERE accident_id = ?`
	rows, err := s.db.Query(query, accidentId)
	if err != nil {
		return nil, fmt.Errorf("error querying injuries: %w", err)
	}
	defer rows.Close()

	var injuries []*models.Injury
	for rows.Next() {
		var injury models.Injury
		if err := rows.Scan(&injury.ID, &injury.PersonType, &injury.InjurySeverity, &injury.Count, &injury.AccidentID); err != nil {
			return nil, fmt.Errorf("error scanning injury details: %w", err)
		}
		injuries = append(injuries, &injury)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over injuries: %w", err)
	}

	return injuries, nil
}
