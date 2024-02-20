// This program processes CSV data on aircraft incidents, importing it into a MySQL database.
// It includes environment variable loading, database connection setup, CSV file parsing,
// and data insertion with comprehensive error handling and logging.

package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/computers33333/airaccidentdata/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Database setup
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Database setup failed: %v", err)
	}
	defer db.Close()

	// Open CSV file
	file, err := os.Open("downloaded_file.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Read and process CSV file
	if err := processCSV(file, db); err != nil {
		log.Fatalf("Failed to process CSV: %v", err)
	}

	log.Println("File processing completed successfully.")
}

// setupDatabase establishes a connection to the MySQL database.
func setupDatabase() (*sql.DB, error) {
	// Retrieve mysql environment variables.
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DATABASE")

	// Construct the Data Source Name (DSN) for the database connection.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, database)

	// Attempt to open a connection to the database with the specified DSN.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		// Return an error if the database connection cannot be established.
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	// Ping the database to verify the connection is established and reachable.
	if err := db.Ping(); err != nil {
		// Return an error if the database is not responding.
		return nil, fmt.Errorf("database is not reachable: %w", err)
	}

	// Return the database connection handle.
	return db, nil
}

// processCSV reads and processes the CSV file, inserting data into the database.
func processCSV(file *os.File, db *sql.DB) error {
	reader := csv.NewReader(file)

	// Skip header row
	if _, err := reader.Read(); err != nil {
		return fmt.Errorf("failed to read headers: %w", err)
	}

	// Iteratively read each record
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break // End of file reached
			}
			// Log any errors encountered during reading, then stop processing.
			return fmt.Errorf("error reading record: %w", err)
		}

		// Attempt to parse the current CSV record into an AircraftAccident structure.
		incident, err := parseRecordToIncident(record)
		if err != nil {
			// Log any errors encountered during parsing and skip to the next record.
			log.Printf("Error parsing record: %v", err)
			continue
		}

		// Attempt to insert the parsed incident into the database.
		if err := insertIncident(context.Background(), db, &incident); err != nil {
			// Log any errors encountered during database insertion.
			log.Printf("Error inserting incident: %v", err)
		}
	}

	// Successfully processed all records without critical errors.
	return nil
}

// loadEnv searches for the .env file starting in the current directory and moving up.
func loadEnv() error {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		// Check if .env exists in this directory.
		if _, err := os.Stat(filepath.Join(dir, ".env")); err == nil {
			// Load the .env file.
			return godotenv.Load(filepath.Join(dir, ".env"))
		}

		// Move up to the parent directory.
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// Root of the filesystem reached, .env not found
			return fmt.Errorf("root directory reached, .env file not found")
		}
		dir = parentDir
	}
}

// Converts string to int, returns 0 if conversion fails.
func atoiSafe(s string) int {
	// Attempt to convert the string to an integer.
	value, err := strconv.Atoi(s)
	if err != nil {
		// Log and handle any conversion error, returning 0 as a safe fallback.
		log.Printf("Error converting string to int: %v", err)
		return 0
	}
	// Return the successfully converted integer value.
	return value
}

// Helper function to parse a date string into time.Time, returns zero value on error.
func parseDate(dateStr string) time.Time {
	// Define the expected date format and attempt to parse the string.
	layout := "02-Jan-06"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		// If parsing fails, log the error and return the zero time value.
		log.Printf("Error parsing date: %v", err)
		return time.Time{}
	}
	// Return the parsed time value.
	return t
}

// Helper function to format a time string, returns empty string on error.
func parseTime(timeStr string) string {
	// Define the expected time format and attempt to parse the string.
	layout := "15:04:05Z"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		// Log any parsing errors and return an empty string as a fallback.
		log.Printf("Error parsing time: %v", err)
		return ""
	}
	// Format and return the time in a MySQL-compatible format.
	return t.Format("15:04:05")
}

// parseRecordToIncident converts a CSV record to an AircraftAccident struct.
func parseRecordToIncident(record []string) (models.AircraftAccident, error) {
	if len(record) < 42 {
		return models.AircraftAccident{}, fmt.Errorf("record does not have enough columns")
	}

	// Parse each field from the record.
	incident := models.AircraftAccident{
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
		RegistrationNumber:        record[10],
		FlightNumber:              record[11],
		AircraftOperator:          record[12],
		AircraftMakeName:          record[13],
		AircraftModelName:         record[14],
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

	return incident, nil
}

func insertIncident(ctx context.Context, db *sql.DB, incident *models.AircraftAccident) error {
	stmt := `INSERT INTO AircraftAccidents (
		updated, entry_date, event_local_date, event_local_time,
		location_city_name, location_state_name, location_country_name,
		remark_text, event_type_description, fsdo_description, registration_number,
		flight_number, aircraft_operator, aircraft_make_name, aircraft_model_name,
		aircraft_missing_flag, aircraft_damage_description, flight_activity, flight_phase,
		far_part, max_injury_level, fatal_flag, flight_crew_injury_none,
		flight_crew_injury_minor, flight_crew_injury_serious, flight_crew_injury_fatal,
		flight_crew_injury_unknown, cabin_crew_injury_none, cabin_crew_injury_minor,
		cabin_crew_injury_serious, cabin_crew_injury_fatal, cabin_crew_injury_unknown,
		passenger_injury_none, passenger_injury_minor, passenger_injury_serious,
		passenger_injury_fatal, passenger_injury_unknown, ground_injury_none,
		ground_injury_minor, ground_injury_serious, ground_injury_fatal, ground_injury_unknown
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE 
		updated = VALUES(updated),
		entry_date = VALUES(entry_date),
		event_local_date = VALUES(event_local_date),
		event_local_time = VALUES(event_local_time),
		location_city_name = VALUES(location_city_name),
		location_state_name = VALUES(location_state_name),
		location_country_name = VALUES(location_country_name),
		remark_text = VALUES(remark_text),
		event_type_description = VALUES(event_type_description),
		fsdo_description = VALUES(fsdo_description),
		flight_number = VALUES(flight_number),
		aircraft_operator = VALUES(aircraft_operator),
		aircraft_make_name = VALUES(aircraft_make_name),
		aircraft_model_name = VALUES(aircraft_model_name),
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
		ground_injury_unknown = VALUES(ground_injury_unknown);`

	// Debugging: Print the SQL statement and the values being inserted.
	fmt.Println("Executing SQL Statement: ", stmt)
	fmt.Printf("Inserting values: %+v\n", incident)

	_, err := db.ExecContext(ctx, stmt,
		incident.Updated, incident.EntryDate.Format("2006-01-02"), incident.EventLocalDate.Format("2006-01-02"), incident.EventLocalTime,
		incident.LocationCityName, incident.LocationStateName, incident.LocationCountryName,
		incident.RemarkText, incident.EventTypeDescription, incident.FSDODescription, incident.RegistrationNumber,
		incident.FlightNumber, incident.AircraftOperator, incident.AircraftMakeName, incident.AircraftModelName,
		incident.AircraftMissingFlag, incident.AircraftDamageDescription, incident.FlightActivity, incident.FlightPhase,
		incident.FARPart, incident.MaxInjuryLevel, incident.FatalFlag, incident.FlightCrewInjuryNone,
		incident.FlightCrewInjuryMinor, incident.FlightCrewInjurySerious, incident.FlightCrewInjuryFatal,
		incident.FlightCrewInjuryUnknown, incident.CabinCrewInjuryNone, incident.CabinCrewInjuryMinor,
		incident.CabinCrewInjurySerious, incident.CabinCrewInjuryFatal, incident.CabinCrewInjuryUnknown,
		incident.PassengerInjuryNone, incident.PassengerInjuryMinor, incident.PassengerInjurySerious,
		incident.PassengerInjuryFatal, incident.PassengerInjuryUnknown, incident.GroundInjuryNone,
		incident.GroundInjuryMinor, incident.GroundInjurySerious, incident.GroundInjuryFatal, incident.GroundInjuryUnknown,
	)
	if err != nil {
		fmt.Printf("Error executing SQL statement: %v\n", err)
		return fmt.Errorf("error inserting incident: %w", err)
	}
	return nil
}
