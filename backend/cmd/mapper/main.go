// Package main provides functionality to process CSV data and insert it into a MySQL database.
package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"github.com/computers33333/airaccidentdata/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// main is the entry point of the application.
func main() {
	// Load environment variables from .env file.
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Setup database connection.
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Database setup failed: %v", err)
	}
	defer db.Close()

	// Open and process CSV file.
	file, err := os.Open("downloaded_file.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	if err := processCSV(file, db); err != nil {
		log.Fatalf("Failed to process CSV: %v", err)
	}

	log.Println("File processing completed successfully.")
}

// setupDatabase establishes a connection to the MySQL database.
func setupDatabase() (*sql.DB, error) {
	// Retrieve MySQL environment variables.
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DATABASE")

	// Construct DSN for database connection.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, database)

	// Open a connection to the database.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	// Ping the database to verify connectivity.
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database is not reachable: %w", err)
	}

	return db, nil
}

// processCSV reads and processes the CSV file, inserting data into the database.
func processCSV(file *os.File, db *sql.DB) error {
	reader := csv.NewReader(file)

	// Skip header row.
	if _, err := reader.Read(); err != nil {
		return fmt.Errorf("failed to read headers: %w", err)
	}

	// Read all records into memory.
	var records [][]string
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading record: %w", err)
		}
		records = append(records, record)
	}

	// Sort records by ENTRY_DATE in descending order.
	sort.Slice(records, func(i, j int) bool {
		entryDate1 := parseDate(records[i][1])
		entryDate2 := parseDate(records[j][1])
		return entryDate1.After(entryDate2)
	})

	// Insert sorted records into the database.
	for _, record := range records {
		aircraft, incident, err := parseRecordToIncident(record)
		if err != nil {
			log.Printf("Error parsing record: %v", err)
			continue
		}

		// Fetch coordinates
		location := fmt.Sprintf("%s, %s, %s", incident.LocationCityName, incident.LocationStateName, incident.LocationCountryName)
		lat, lng, err := getCoordinates(location)
		if err != nil {
			log.Printf("Failed to get coordinates for %s: %v", location, err)
			continue
		}

		aircraftID, err := getAircraftIDByRegistration(context.Background(), db, aircraft.RegistrationNumber)
		if err != nil {
			return fmt.Errorf("error checking aircraft existence: %w", err)
		}

		if aircraftID == 0 {
			aircraftID, err = insertAircraft(context.Background(), db, aircraft.RegistrationNumber, aircraft.AircraftMakeName, aircraft.AircraftModelName, aircraft.AircraftOperator)
			if err != nil {
				return fmt.Errorf("error adding aircraft: %w", err)
			}
		}

		if err := insertAccident(context.Background(), db, aircraftID, incident, lat, lng); err != nil {
			return fmt.Errorf("error inserting accident: %w", err)
		}
	}

	return nil
}

// loadEnv searches for the .env file starting in the current directory and moving up.
func loadEnv() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".env")); err == nil {
			return godotenv.Load(filepath.Join(dir, ".env"))
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return fmt.Errorf("root directory reached, .env file not found")
		}
		dir = parentDir
	}
}

// atoiSafe converts string to int, returns 0 if conversion fails or the string is empty.
func atoiSafe(s string) int {
	if s == "" {
		return 0
	}
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		return 0
	}
	return value
}

// Helper function to parse a date string into time.Time, returns zero value on error.
func parseDate(dateStr string) time.Time {
	layout := "02-Jan-06"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		log.Printf("Error parsing date: %v", err)
		return time.Time{}
	}
	return t
}

// Helper function to format a time string, returns empty string on error.
func parseTime(timeStr string) string {
	layout := "15:04:05Z"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		log.Printf("Error parsing time: %v", err)
		return ""
	}
	return t.Format("15:04:05")
}

// parseRecordToIncident converts a CSV record to an AircraftAccident struct.
func parseRecordToIncident(record []string) (*models.Aircraft, *models.AircraftAccident, error) {
	if len(record) < 42 {
		return nil, nil, fmt.Errorf("record does not have enough columns")
	}

	// Parse the fields for Aircraft struct
	aircraft := &models.Aircraft{
		RegistrationNumber: record[10],
		AircraftMakeName:   record[13],
		AircraftModelName:  record[14],
		AircraftOperator:   record[12],
	}

	// Parse the fields for AircraftAccident struct
	incident := &models.AircraftAccident{
		Updated:                   record[0],
		EntryDate:                 parseDate(record[1]),
		EventLocalDate:            parseDate(record[2]),
		EventLocalTime:            parseTime(record[3]),
		LocationCityName:          record[4],
		LocationStateName:         record[5],
		LocationCountryName:       record[6],
		RemarkText:                record[7],
		EventTypeDescription:      record[8],
		FSDODescription:           record[9],
		FlightNumber:              record[11],
		AircraftMissingFlag:       record[15],
		AircraftDamageDescription: record[16],
		FlightActivity:            record[17],
		FlightPhase:               record[18],
		FARPart:                   record[19],
		MaxInjuryLevel:            record[20],
		FatalFlag:                 record[21],
		FlightCrewInjuryNone:      atoiSafe(record[22]),
		FlightCrewInjuryMinor:     atoiSafe(record[23]),
		FlightCrewInjurySerious:   atoiSafe(record[24]),
		FlightCrewInjuryFatal:     atoiSafe(record[25]),
		FlightCrewInjuryUnknown:   atoiSafe(record[26]),
		CabinCrewInjuryNone:       atoiSafe(record[27]),
		CabinCrewInjuryMinor:      atoiSafe(record[28]),
		CabinCrewInjurySerious:    atoiSafe(record[29]),
		CabinCrewInjuryFatal:      atoiSafe(record[30]),
		CabinCrewInjuryUnknown:    atoiSafe(record[31]),
		PassengerInjuryNone:       atoiSafe(record[32]),
		PassengerInjuryMinor:      atoiSafe(record[33]),
		PassengerInjurySerious:    atoiSafe(record[34]),
		PassengerInjuryFatal:      atoiSafe(record[35]),
		PassengerInjuryUnknown:    atoiSafe(record[36]),
		GroundInjuryNone:          atoiSafe(record[37]),
		GroundInjuryMinor:         atoiSafe(record[38]),
		GroundInjurySerious:       atoiSafe(record[39]),
		GroundInjuryFatal:         atoiSafe(record[40]),
		GroundInjuryUnknown:       atoiSafe(record[41]),
	}

	return aircraft, incident, nil
}

// insertAircraft inserts a new aircraft into the database or updates it if it already exists.
func insertAircraft(ctx context.Context, db *sql.DB, registrationNumber, aircraftMake, aircraftModel, aircraftOperator string) (int, error) {
	stmt := `
		INSERT INTO Aircrafts (registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator) 
		VALUES (?, ?, ?, ?) 
		ON DUPLICATE KEY UPDATE 
			aircraft_make_name = VALUES(aircraft_make_name), 
			aircraft_model_name = VALUES(aircraft_model_name), 
			aircraft_operator = VALUES(aircraft_operator)
	`

	// Execute the SQL statement
	_, err := db.ExecContext(ctx, stmt, registrationNumber, aircraftMake, aircraftModel, aircraftOperator)
	if err != nil {
		return 0, fmt.Errorf("error inserting or updating aircraft: %w", err)
	}

	// Retrieve the ID of the inserted or updated aircraft
	var aircraftID int
	err = db.QueryRowContext(ctx, "SELECT id FROM Aircrafts WHERE registration_number = ?", registrationNumber).Scan(&aircraftID)
	if err != nil {
		return 0, fmt.Errorf("error retrieving aircraft ID: %w", err)
	}

	return aircraftID, nil
}

// insertAccident inserts or updates an accident associated with an aircraft in the database.
func insertAccident(ctx context.Context, db *sql.DB, aircraftID int, incident *models.AircraftAccident, lat float64, lng float64) error {
	// Check if the accident already exists based on unique constraints
	var existingAccidentID int
	checkStmt := `
        SELECT id FROM Accidents
        WHERE aircraft_id = ? AND event_local_date = ? AND event_local_time = ?
    `
	err := db.QueryRowContext(ctx, checkStmt, aircraftID, incident.EventLocalDate, incident.EventLocalTime).Scan(&existingAccidentID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error checking for existing accident: %w", err)
	}

	if existingAccidentID != 0 {
		// Accident already exists, update the existing entry
		updateStmt := `
		UPDATE Accidents
		SET updated = ?, 
			entry_date = ?, 
			event_local_date = ?, 
			event_local_time = ?, 
			location_city_name = ?, 
			location_state_name = ?, 
			location_country_name = ?, 
			latitude = ?, 
			longitude = ?, 
			remark_text = ?, 
			event_type_description = ?, 
			fsdo_description = ?, 
			flight_number = ?, 
			aircraft_missing_flag = ?, 
			aircraft_damage_description = ?, 
			flight_activity = ?, 
			flight_phase = ?, 
			far_part = ?, 
			max_injury_level = ?, 
			fatal_flag = ?, 
			flight_crew_injury_none = ?, 
			flight_crew_injury_minor = ?, 
			flight_crew_injury_serious = ?, 
			flight_crew_injury_fatal = ?, 
			flight_crew_injury_unknown = ?, 
			cabin_crew_injury_none = ?, 
			cabin_crew_injury_minor = ?, 
			cabin_crew_injury_serious = ?, 
			cabin_crew_injury_fatal = ?, 
			cabin_crew_injury_unknown = ?, 
			passenger_injury_none = ?, 
			passenger_injury_minor = ?, 
			passenger_injury_serious = ?, 
			passenger_injury_fatal = ?, 
			passenger_injury_unknown = ?, 
			ground_injury_none = ?, 
			ground_injury_minor = ?, 
			ground_injury_serious = ?, 
			ground_injury_fatal = ?, 
			ground_injury_unknown = ?
		WHERE id = ?
	`

		_, err := db.ExecContext(ctx, updateStmt, incident.Updated, incident.EntryDate, incident.EventLocalDate, incident.EventLocalTime, incident.LocationCityName, incident.LocationStateName, incident.LocationCountryName, lat, lng, incident.RemarkText, incident.EventTypeDescription, incident.FSDODescription, incident.FlightNumber, incident.AircraftMissingFlag, incident.AircraftDamageDescription, incident.FlightActivity, incident.FlightPhase, incident.FARPart, incident.MaxInjuryLevel, incident.FatalFlag, incident.FlightCrewInjuryNone, incident.FlightCrewInjuryMinor, incident.FlightCrewInjurySerious, incident.FlightCrewInjuryFatal, incident.FlightCrewInjuryUnknown, incident.CabinCrewInjuryNone, incident.CabinCrewInjuryMinor, incident.CabinCrewInjurySerious, incident.CabinCrewInjuryFatal, incident.CabinCrewInjuryUnknown, incident.PassengerInjuryNone, incident.PassengerInjuryMinor, incident.PassengerInjurySerious, incident.PassengerInjuryFatal, incident.PassengerInjuryUnknown, incident.GroundInjuryNone, incident.GroundInjuryMinor, incident.GroundInjurySerious, incident.GroundInjuryFatal, incident.GroundInjuryUnknown, existingAccidentID)
		if err != nil {
			return fmt.Errorf("error updating existing accident: %w", err)
		}

		return nil
	}

	stmt := `
    INSERT INTO Accidents (
        updated, 
        entry_date, 
        event_local_date, 
        event_local_time,
        location_city_name,
        location_state_name,
        location_country_name,
		latitude,
		longitude,
        remark_text,
        event_type_description,
        fsdo_description,
        flight_number,
        aircraft_missing_flag,
        aircraft_damage_description,
        flight_activity,
        flight_phase,
        far_part,
        max_injury_level,
        fatal_flag,
        flight_crew_injury_none,
        flight_crew_injury_minor,
        flight_crew_injury_serious,
        flight_crew_injury_fatal,
        flight_crew_injury_unknown,
        cabin_crew_injury_none,
        cabin_crew_injury_minor,
        cabin_crew_injury_serious,
        cabin_crew_injury_fatal,
        cabin_crew_injury_unknown,
        passenger_injury_none,
        passenger_injury_minor,
        passenger_injury_serious,
        passenger_injury_fatal,
        passenger_injury_unknown,
        ground_injury_none,
        ground_injury_minor,
        ground_injury_serious,
        ground_injury_fatal,
        ground_injury_unknown,
        aircraft_id
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    ON DUPLICATE KEY UPDATE
        updated = VALUES(updated),
        entry_date = VALUES(entry_date),
        event_local_date = VALUES(event_local_date),
        event_local_time = VALUES(event_local_time),
        location_city_name = VALUES(location_city_name),
        location_state_name = VALUES(location_state_name),
        location_country_name = VALUES(location_country_name),
		latitude = VALUES(latitude),
        longitude = VALUES(longitude),
        remark_text = VALUES(remark_text),
        event_type_description = VALUES(event_type_description),
        fsdo_description = VALUES(fsdo_description),
        flight_number = VALUES(flight_number),
        aircraft_missing_flag = VALUES(aircraft_missing_flag),
        aircraft_damage_description = VALUES(aircraft_damage_description),
        flight_activity = VALUES(flight_activity),
        flight_phase = VALUES(flight_phase),
        far_part = VALUES(far_part),
        max_injury_level = VALUES(max_injury_level),
        fatal_flag = VALUES(fatal_flag),
        flight_crew_injury_none = VALUES(flight_crew_injury_none),
        flight_crew_injury_minor = VALUES(flight_crew_injury_minor),
        flight_crew_injury_serious = VALUES(flight_crew_injury_serious),
        flight_crew_injury_fatal = VALUES(flight_crew_injury_fatal),
        flight_crew_injury_unknown = VALUES(flight_crew_injury_unknown),
        cabin_crew_injury_none = VALUES(cabin_crew_injury_none),
        cabin_crew_injury_minor = VALUES(cabin_crew_injury_minor),
        cabin_crew_injury_serious = VALUES(cabin_crew_injury_serious),
        cabin_crew_injury_fatal = VALUES(cabin_crew_injury_fatal),
        cabin_crew_injury_unknown = VALUES(cabin_crew_injury_unknown),
        passenger_injury_none = VALUES(passenger_injury_none),
        passenger_injury_minor = VALUES(passenger_injury_minor),
        passenger_injury_serious = VALUES(passenger_injury_serious),
        passenger_injury_fatal = VALUES(passenger_injury_fatal),
        passenger_injury_unknown = VALUES(passenger_injury_unknown),
        ground_injury_none = VALUES(ground_injury_none),
        ground_injury_minor = VALUES(ground_injury_minor),
        ground_injury_serious = VALUES(ground_injury_serious),
        ground_injury_fatal = VALUES(ground_injury_fatal),
        ground_injury_unknown = VALUES(ground_injury_unknown),
        aircraft_id = VALUES(aircraft_id)
	`
	_, execErr := db.ExecContext(ctx, stmt,
		incident.Updated,
		incident.EntryDate.Format("2006-01-02"),
		incident.EventLocalDate.Format("2006-01-02"),
		incident.EventLocalTime,
		incident.LocationCityName,
		incident.LocationStateName,
		incident.LocationCountryName,
		lat,
		lng,
		incident.RemarkText,
		incident.EventTypeDescription,
		incident.FSDODescription,
		incident.FlightNumber,
		incident.AircraftMissingFlag,
		incident.AircraftDamageDescription,
		incident.FlightActivity,
		incident.FlightPhase,
		incident.FARPart,
		incident.MaxInjuryLevel,
		incident.FatalFlag,
		incident.FlightCrewInjuryNone,
		incident.FlightCrewInjuryMinor,
		incident.FlightCrewInjurySerious,
		incident.FlightCrewInjuryFatal,
		incident.FlightCrewInjuryUnknown,
		incident.CabinCrewInjuryNone,
		incident.CabinCrewInjuryMinor,
		incident.CabinCrewInjurySerious,
		incident.CabinCrewInjuryFatal,
		incident.CabinCrewInjuryUnknown,
		incident.PassengerInjuryNone,
		incident.PassengerInjuryMinor,
		incident.PassengerInjurySerious,
		incident.PassengerInjuryFatal,
		incident.PassengerInjuryUnknown,
		incident.GroundInjuryNone,
		incident.GroundInjuryMinor,
		incident.GroundInjurySerious,
		incident.GroundInjuryFatal,
		incident.GroundInjuryUnknown,
		aircraftID,
	)
	if execErr != nil {
		return fmt.Errorf("error inserting or updating accident: %w", execErr)
	}
	return nil
}

// getAircraftIDByRegistration retrieves the aircraft ID from the database based on the registration number.
func getAircraftIDByRegistration(ctx context.Context, db *sql.DB, registrationNumber string) (int, error) {
	var aircraftID int
	err := db.QueryRowContext(ctx, "SELECT id FROM Aircrafts WHERE registration_number = ?", registrationNumber).Scan(&aircraftID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Aircraft does not exist in the database.
			return 0, nil
		}
		return 0, fmt.Errorf("error querying database: %w", err)
	}
	return aircraftID, nil
}

func getCoordinates(place string) (float64, float64, error) {
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	baseUrl := "https://maps.googleapis.com/maps/api/geocode/json"

	// Construct request URL
	requestUrl := fmt.Sprintf("%s?address=%s&key=%s", baseUrl, url.QueryEscape(place), apiKey)

	// Make the HTTP request
	resp, err := http.Get(requestUrl)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var geoResp models.GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return 0, 0, err
	}

	if len(geoResp.Results) == 0 {
		return 0, 0, fmt.Errorf("no results found for %s", place)
	}

	lat := geoResp.Results[0].Geometry.Location.Lat
	lng := geoResp.Results[0].Geometry.Location.Lng

	return lat, lng, nil
}
