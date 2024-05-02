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

var personTypeBaseIndex = map[string]int{
	"flight_crew": 22, // Starting index for flight crew injuries
	"cabin_crew":  27, // Starting index for cabin crew injuries
	"passengers":  32, // Starting index for passenger injuries
	"ground":      37, // Starting index for ground personnel injuries
}

var injurySeverities = []string{"none", "minor", "serious", "fatal", "unknown"}

// main is the entry point of the application
func main() {
	// Load environment variables from .env file
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

// setupDatabase establishes a connection to the MySQL database
func setupDatabase() (*sql.DB, error) {
	// Retrieve MySQL environment variables
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DATABASE")

	// Construct DSN for database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, database)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	// Ping the database to verify connectivity
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database is not reachable: %w", err)
	}

	return db, nil
}

// processCSV reads and processes the CSV file, inserting data into the database
func processCSV(file *os.File, db *sql.DB) error {
	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil { // Skip header
		return err
	}

	var records [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		records = append(records, record)
	}

	// Sort records by ENTRY_DATE in descending order.
	sort.Slice(records, func(i, j int) bool {
		entryDate1, err1 := parseDate(records[i][1])
		entryDate2, err2 := parseDate(records[j][1])
		if err1 != nil || err2 != nil {
			return false
		}
		return entryDate1.After(entryDate2)
	})

	for i, record := range records {
		if i < 3 {
			fmt.Println(record)

			if err := processRecord(context.Background(), db, record); err != nil {
				log.Printf("Failed to process record: %v", err)
				continue
			}
		}
	}

	return nil
}

// Read and parse each CSV row into a structured format
func processRecord(ctx context.Context, db *sql.DB, record []string) error {
	// Parse the record to get aircraft and accident data
	aircraft, accident, err := parseRecordToIncident(record)
	if err != nil {
		return err
	}

	// Ensure aircraft is in the database and get its ID
	aircraftID, err := ensureAircraft(ctx, db, aircraft)
	if err != nil {
		return err
	}

	// Ensure location is in the database and get its ID
	locationID, err := ensureLocation(ctx, db, nil)
	if err != nil {
		return err
	}

	// Insert the accident with references to aircraft_id and location_id, and get accident ID
	accidentID, err := insertAccident(ctx, db, aircraftID, locationID, accident)
	if err != nil {
		return err
	}

	// Extract and insert injuries associated with the accident ID
	injuries, err := extractInjuriesFromRecord(record, accidentID)
	if err != nil {
		return err
	}

	// Proceed to insert injuries
	err = insertInjuries(ctx, db, accidentID, injuries)
	if err != nil {
		log.Printf("Failed to insert injuries: %v", err)
		return err
	}

	// If everything executed correctly, return nil
	return nil
}

// loadEnv searches for the .env file starting in the current directory and moving up
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

// atoiSafe converts string to int, returns 0 if conversion fails or the string is empty
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

// Helper function to parse a date string into time.Time, returns time.Time and error
func parseDate(dateStr string) (time.Time, error) {
	layout := "02-Jan-06"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing date '%s': %v", dateStr, err)
	}
	return t, nil
}

// Helper function to format a time string into time.Time, returns time.Time and error
func parseTime(timeStr string) (string, error) {
	layout := "15:04:05Z"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return "", fmt.Errorf("error parsing time '%s': %v", timeStr, err)
	}
	return t.Format("15:04:05"), nil
}

// parseRecordToIncident converts a CSV record to an Accident struct
func parseRecordToIncident(record []string) (*models.Aircraft, *models.Accident, error) {
	if len(record) < 42 { // Ensure there are enough columns to parse
		return nil, nil, fmt.Errorf("record does not have enough columns")
	}

	// Parse the fields for Aircraft struct
	aircraft := &models.Aircraft{
		RegistrationNumber: record[10],
		AircraftMakeName:   record[13],
		AircraftModelName:  record[14],
		AircraftOperator:   record[12],
	}

	// Parse the fields for Accident struct
	entryDate, err := parseDate(record[1])
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing entry date: %v", err)
	}
	eventLocalDate, err := parseDate(record[2])
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing event local date: %v", err)
	}
	eventLocalTime, err := parseTime(record[3])
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing event local time: %v", err)
	}

	incident := &models.Accident{
		Updated:                   record[0],
		EntryDate:                 entryDate,
		EventLocalDate:            eventLocalDate,
		EventLocalTime:            eventLocalTime,
		RemarkText:                record[7],
		EventTypeDescription:      record[8],
		FSDODescription:           record[9],
		FlightNumber:              record[11],
		AircraftMissingFlag:       record[15],
		AircraftDamageDescription: record[16],
		FlightActivity:            record[17],
		FlightPhase:               record[18],
		FARPart:                   record[19],
		FatalFlag:                 record[21],
	}

	return aircraft, incident, nil
}

// insertAircraft inserts a new aircraft into the database or updates it if it already exists
func insertAircraft(ctx context.Context, db *sql.DB, registrationNumber, aircraftMake, aircraftModel, aircraftOperator string) (int, error) {
	stmt := `
		INSERT INTO Aircrafts (registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator) 
		VALUES (?, ?, ?, ?);
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

// insertAccident inserts or updates an accident associated with an aircraft in the database
func insertAccident(ctx context.Context, db *sql.DB, aircraftID, locationID int, accident *models.Accident) (int, error) {
	stmt := `
    INSERT INTO Accidents (updated, entry_date, event_local_date, event_local_time, remark_text, event_type_description, fsdo_description, flight_number, aircraft_missing_flag, aircraft_damage_description, flight_activity, flight_phase, far_part, fatal_flag, location_id, aircraft_id)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	_, err := db.ExecContext(ctx, stmt, accident.Updated, accident.EntryDate, accident.EventLocalDate, accident.EventLocalTime, accident.RemarkText, accident.EventTypeDescription, accident.FSDODescription, accident.FlightNumber, accident.AircraftMissingFlag, accident.AircraftDamageDescription, accident.FlightActivity, accident.FlightPhase, accident.FARPart, accident.FatalFlag, locationID, aircraftID)
	if err != nil {
		return 0, err
	}

	var accidentID int
	err = db.QueryRowContext(ctx, "SELECT LAST_INSERT_ID()").Scan(&accidentID)
	if err != nil {
		return 0, err
	}

	return accidentID, nil
}

// Ensures the aircraft is in the Aircrafts table and returns the ID
func ensureAircraft(ctx context.Context, db *sql.DB, aircraft *models.Aircraft) (int, error) {
	stmt := `
    INSERT INTO Aircrafts (registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator)
    VALUES (?, ?, ?, ?);
    `

	_, err := db.ExecContext(ctx, stmt, aircraft.RegistrationNumber, aircraft.AircraftMakeName, aircraft.AircraftModelName, aircraft.AircraftOperator)
	if err != nil {
		return 0, err
	}

	var aircraftID int
	err = db.QueryRowContext(ctx, "SELECT id FROM Aircrafts WHERE registration_number = ?", aircraft.RegistrationNumber).Scan(&aircraftID)
	if err != nil {
		return 0, err
	}

	return aircraftID, nil
}

// Ensures the location is in the Locations table and returns the ID
func ensureLocation(ctx context.Context, db *sql.DB, location *models.Location) (int, error) {
	// If the location is nil, insert a default location
	if location == nil {
		// Inserting a dummy location
		_, err := db.ExecContext(ctx, "INSERT INTO Locations (city_name, state_name, country_name, latitude, longitude) VALUES (?, ?, ?, ?, ?)",
			"Default City", "Default State", "Default Country", 0.0, 0.0)
		if err != nil {
			return 0, err
		}

		// Retrieve the ID of the inserted location
		var locationID int
		err = db.QueryRowContext(ctx, "SELECT LAST_INSERT_ID()").Scan(&locationID)
		if err != nil {
			return 0, err
		}
		return locationID, nil
	}

	// Insert the provided location
	_, err := db.ExecContext(ctx, "INSERT INTO Locations (city_name, state_name, country_name, latitude, longitude) VALUES (?, ?, ?, ?, ?)",
		location.CityName, location.StateName, location.CountryName, location.Latitude, location.Longitude)
	if err != nil {
		return 0, err
	}

	// Retrieve the ID of the inserted location
	var locationID int
	err = db.QueryRowContext(ctx, "SELECT LAST_INSERT_ID()").Scan(&locationID)
	if err != nil {
		return 0, err
	}
	return locationID, nil
}

// Function to extract injuries from a record
func extractInjuriesFromRecord(record []string, accidentID int) ([]*models.Injury, error) {
	var injuries []*models.Injury

	for personType, baseIndex := range personTypeBaseIndex {
		for offset, severity := range injurySeverities {
			countIndex := baseIndex + offset
			count, err := strconv.Atoi(record[countIndex])
			if err != nil {
				log.Printf("Error converting string to int for %s %s: %v", personType, severity, err)
				continue
			}
			injuries = append(injuries, &models.Injury{
				PersonType:     personType,
				InjurySeverity: severity,
				Count:          count,
				AccidentID:     accidentID,
			})
		}
	}

	return injuries, nil
}

// Function that takes injury objects and inserts them into the database using the accident_id to link them
func insertInjuries(ctx context.Context, db *sql.DB, accidentID int, injuries []*models.Injury) error {
	stmt := `INSERT INTO Injuries (accident_id, person_type, injury_severity, count) VALUES (?, ?, ?, ?)`
	for _, injury := range injuries {
		_, err := db.ExecContext(ctx, stmt, accidentID, injury.PersonType, injury.InjurySeverity, injury.Count)
		if err != nil {
			return fmt.Errorf("error inserting injury data: %w", err)
		}
	}
	return nil
}

// getAircraftIDByRegistration retrieves the aircraft ID from the database based on the registration number
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
