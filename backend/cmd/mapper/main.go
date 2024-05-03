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
	"sort"
	"strconv"

	"github.com/computers33333/airaccidentdata/internal/models"
	"github.com/computers33333/airaccidentdata/internal/shared"
	_ "github.com/go-sql-driver/mysql" // Blank identifier imports MySQL driver to initialize and register it.
)

// main is the entry point of the application, responsible for processing CSV data and inserting it into a MySQL database.
func main() {
	if err := shared.LoadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := shared.SetupDatabase()
	if err != nil {
		log.Fatalf("Database setup failed: %v", err)
	}
	defer db.Close()

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

// processCSV reads and processes the CSV file, inserting data into the database.
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

	// Sort records by ENTRY_DATE in descending order
	sort.Slice(records, func(i, j int) bool {
		entryDate1, err1 := shared.ParseDate(records[i][1])
		entryDate2, err2 := shared.ParseDate(records[j][1])
		if err1 != nil || err2 != nil {
			return false
		}
		return entryDate1.After(entryDate2)
	})

	for _, record := range records {
		if err := processRecord(context.Background(), db, record); err != nil {
			log.Printf("Failed to process record: %v", err)
			continue
		}
	}

	return nil
}

// Read and parse each CSV row into a structured format.
func processRecord(ctx context.Context, db *sql.DB, record []string) error {
	aircraft, accident, location, err := parseRecordToIncident(record)
	if err != nil {
		return err
	}

	aircraftID, err := ensureAircraft(ctx, db, aircraft)
	if err != nil {
		return err
	}

	locationID, err := ensureLocation(ctx, db, location)
	if err != nil {
		return err
	}

	accidentID, err := insertAccident(ctx, db, aircraftID, locationID, accident)
	if err != nil {
		return err
	}

	injuries, err := extractInjuriesFromRecord(record, accidentID)
	if err != nil {
		return err
	}

	err = insertInjuries(ctx, db, accidentID, injuries)
	if err != nil {
		log.Printf("Failed to insert injuries: %v", err)
		return err
	}

	return nil
}

// parseRecordToIncident converts a CSV record to an Accident struct.
func parseRecordToIncident(record []string) (*models.Aircraft, *models.Accident, *models.Location, error) {
	if len(record) < 42 { // Ensure there are enough columns to parse
		return nil, nil, nil, fmt.Errorf("record does not have enough columns")
	}

	aircraft := &models.Aircraft{
		RegistrationNumber: record[10],
		AircraftMakeName:   record[13],
		AircraftModelName:  record[14],
		AircraftOperator:   record[12],
	}

	entryDate, err := shared.ParseDate(record[1])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error parsing entry date: %v", err)
	}
	eventLocalDate, err := shared.ParseDate(record[2])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error parsing event local date: %v", err)
	}
	eventLocalTime, err := shared.ParseTime(record[3])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error parsing event local time: %v", err)
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

	place := fmt.Sprintf("%s, %s, %s", record[4], record[5], record[6]) // Combine city, state, and country names
	latitude, longitude, err := getCoordinates(place)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error getting coordinates for %s: %v", place, err)
	}

	location := &models.Location{
		CityName:    record[4],
		StateName:   record[5],
		CountryName: record[6],
		Latitude:    latitude,
		Longitude:   longitude,
	}

	return aircraft, incident, location, nil
}

// insertAccident inserts or updates an accident associated with an aircraft in the database.
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

// Ensures the aircraft is in the Aircrafts table and returns the ID.
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

// Ensures the location is in the Locations table and returns the ID.
func ensureLocation(ctx context.Context, db *sql.DB, location *models.Location) (int, error) {
	_, err := db.ExecContext(ctx, "INSERT INTO Locations (city_name, state_name, country_name, latitude, longitude) VALUES (?, ?, ?, ?, ?)",
		location.CityName, location.StateName, location.CountryName, location.Latitude, location.Longitude)
	if err != nil {
		return 0, err
	}

	var locationID int
	err = db.QueryRowContext(ctx, "SELECT LAST_INSERT_ID()").Scan(&locationID)
	if err != nil {
		return 0, err
	}
	return locationID, nil
}

// extractInjuriesFromRecord leverages indexed patterns in CSV to categorize injury data by personnel type and severity.
func extractInjuriesFromRecord(record []string, accidentID int) ([]*models.Injury, error) {
	// Constants defining base indexes for each person type
	personTypeBaseIndex := map[string]int{
		"flight_crew": 22,
		"cabin_crew":  27,
		"passengers":  32,
		"ground":      37,
	}

	// Severity levels for injuries
	injurySeverities := []string{"none", "minor", "serious", "fatal", "unknown"}

	var injuries []*models.Injury

	// Iterates over indexed data to extract and categorize injuries based on personnel type and severity.
	for personType, baseIndex := range personTypeBaseIndex {
		for offset, severity := range injurySeverities {
			countIndex := baseIndex + offset
			if record[countIndex] == "" {
				continue
			}
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

// Function that takes injury objects and inserts them into the database using the accident_id to link them.
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

// getCoordinates retrieves the latitude and longitude coordinates of a given place using the Google Maps Geocoding API.
func getCoordinates(place string) (float64, float64, error) {
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	baseUrl := "https://maps.googleapis.com/maps/api/geocode/json"

	requestUrl := fmt.Sprintf("%s?address=%s&key=%s", baseUrl, url.QueryEscape(place), apiKey)

	resp, err := http.Get(requestUrl)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

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
