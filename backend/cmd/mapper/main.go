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
	"sort"
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

	// Read all records into memory
	var records [][]string
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break // End of file reached
			}
			return fmt.Errorf("error reading record: %w", err)
		}
		records = append(records, record)
	}

	// Sort records by ENTRY_DATE in descending order
	sort.Slice(records, func(i, j int) bool {
		entryDate1 := parseDate(records[i][1])
		entryDate2 := parseDate(records[j][1])
		return entryDate1.After(entryDate2)
	})

	// Iteratively insert sorted records into the database
	for _, record := range records {
		// Attempt to parse the current CSV record into an Aircraft and AircraftAccident structures.
		aircraft, incident, err := parseRecordToIncident(record)
		if err != nil {
			// Log any errors encountered during parsing and skip to the next record.
			log.Printf("Error parsing record: %v", err)
			continue
		}

		// Check if the aircraft exists in the database
		aircraftID, err := getAircraftIDByRegistration(context.Background(), db, aircraft.RegistrationNumber)
		if err != nil {
			return fmt.Errorf("error checking aircraft existence: %w", err)
		}

		// If the aircraft does not exist, add it to the database
		if aircraftID == 0 {
			aircraftID, err = insertAircraft(context.Background(), db, aircraft.RegistrationNumber, aircraft.AircraftMakeName, aircraft.AircraftModelName, aircraft.AircraftOperator)
			if err != nil {
				return fmt.Errorf("error adding aircraft: %w", err)
			}
		}

		// Insert the accident associated with the aircraft
		if err := insertAccident(context.Background(), db, aircraftID, incident); err != nil {
			return fmt.Errorf("error inserting accident: %w", err)
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
func parseRecordToIncident(record []string) (*models.Aircraft, *models.AircraftAccident, error) {
	if len(record) < 42 {
		return nil, nil, fmt.Errorf("record does not have enough columns")
	}

	// Parse the fields for Aircraft struct
	aircraft := &models.Aircraft{
		RegistrationNumber: record[10], // REGIST_NBR
		AircraftMakeName:   record[13], // ACFT_MAKE_NAME
		AircraftModelName:  record[14], // ACFT_MODEL_NAME
		AircraftOperator:   record[12], // ACFT_OPRTR
	}

	// Parse the fields for AircraftAccident struct
	incident := &models.AircraftAccident{
		Updated:                   record[0],            // UPDATED
		EntryDate:                 parseDate(record[1]), // ENTRY_DATE
		EventLocalDate:            parseDate(record[2]), // EVENT_LCL_DATE
		EventLocalTime:            parseTime(record[3]), // EVENT_LCL_TIME
		LocationCityName:          record[4],            // LOC_CITY_NAME
		LocationStateName:         record[5],            // LOC_STATE_NAME
		LocationCountryName:       record[6],            // LOC_CNTRY_NAME
		RemarkText:                record[7],            // RMK_TEXT
		EventTypeDescription:      record[8],            // EVENT_TYPE_DESC
		FSDODescription:           record[9],            // FSDO_DESC
		FlightNumber:              record[11],           // FLT_NBR
		AircraftMissingFlag:       record[15],           // ACFT_MISSING_FLAG
		AircraftDamageDescription: record[16],           // ACFT_DMG_DESC
		FlightActivity:            record[17],           // FLT_ACTIVITY
		FlightPhase:               record[18],           // FLT_PHASE
		FARPart:                   record[19],           // FAR_PART
		MaxInjuryLevel:            record[20],           // MAX_INJ_LVL
		FatalFlag:                 record[21],           // FATAL_FLAG
		FlightCrewInjuryNone:      atoiSafe(record[22]), // FLT_CRW_INJ_NONE
		FlightCrewInjuryMinor:     atoiSafe(record[23]), // FLT_CRW_INJ_MINOR
		FlightCrewInjurySerious:   atoiSafe(record[24]), // FLT_CRW_INJ_SERIOUS
		FlightCrewInjuryFatal:     atoiSafe(record[25]), // FLT_CRW_INJ_FATAL
		FlightCrewInjuryUnknown:   atoiSafe(record[26]), // FLT_CRW_INJ_UNK
		CabinCrewInjuryNone:       atoiSafe(record[27]), // CBN_CRW_INJ_NONE
		CabinCrewInjuryMinor:      atoiSafe(record[28]), // CBN_CRW_INJ_MINOR
		CabinCrewInjurySerious:    atoiSafe(record[29]), // CBN_CRW_INJ_SERIOUS
		CabinCrewInjuryFatal:      atoiSafe(record[30]), // CBN_CRW_INJ_FATAL
		CabinCrewInjuryUnknown:    atoiSafe(record[31]), // CBN_CRW_INJ_UNK
		PassengerInjuryNone:       atoiSafe(record[32]), // PAX_INJ_NONE
		PassengerInjuryMinor:      atoiSafe(record[33]), // PAX_INJ_MINOR
		PassengerInjurySerious:    atoiSafe(record[34]), // PAX_INJ_SERIOUS
		PassengerInjuryFatal:      atoiSafe(record[35]), // PAX_INJ_FATAL
		PassengerInjuryUnknown:    atoiSafe(record[36]), // PAX_INJ_UNK
		GroundInjuryNone:          atoiSafe(record[37]), // GRND_INJ_NONE
		GroundInjuryMinor:         atoiSafe(record[38]), // GRND_INJ_MINOR
		GroundInjurySerious:       atoiSafe(record[39]), // GRND_INJ_SERIOUS
		GroundInjuryFatal:         atoiSafe(record[40]), // GRND_INJ_FATAL
		GroundInjuryUnknown:       atoiSafe(record[41]), // GRND_INJ_UNK
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
func insertAccident(ctx context.Context, db *sql.DB, aircraftID int, incident *models.AircraftAccident) error {
	// Check if the accident already exists
	var existingID int
	checkStmt := `
			SELECT id FROM Accidents
			WHERE aircraft_id = ? AND event_local_date = ? AND event_local_time = ?
		`
	err := db.QueryRowContext(ctx, checkStmt, aircraftID, incident.EventLocalDate.Format("2006-01-02"), incident.EventLocalTime).Scan(&existingID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error checking for existing accident: %w", err)
	}

	if existingID != 0 {
		// Accident already exists, handle it accordingly
		return nil // Or return an error if required
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
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
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
		// Return the error if any other error occurs.
		return 0, fmt.Errorf("error querying database: %w", err)
	}
	// Return the aircraft ID if found.
	return aircraftID, nil
}
