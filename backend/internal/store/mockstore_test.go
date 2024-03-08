package store

import (
	"errors"
	"fmt"
	"testing"

	"github.com/computers33333/airaccidentdata/internal/models"
)

// TestMockStore_GetAccidents tests the GetAccidents method of the MockStore.
// It verifies that the method returns the expected accidents and handles errors correctly.
func TestMockStore_GetAccidents(t *testing.T) {
	expectedAccidents := []*models.AircraftAccident{
		{ID: 1},
	}
	mockStore := NewMockStore(nil, expectedAccidents, nil)
	accidents, err := mockStore.GetAccidents()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(accidents) != len(expectedAccidents) {
		t.Fatalf("Expected %d accidents, got %d", len(expectedAccidents), len(accidents))
	}
	for i, accident := range accidents {
		if accident.ID != expectedAccidents[i].ID {
			t.Errorf("Expected ID %d, got %d", expectedAccidents[i].ID, accident.ID)
		}
	}
}

// TestMockStore_SaveAccidents tests the SaveAccidents method of the MockStore.
// It verifies that the method saves accidents correctly and handles errors properly.
func TestMockStore_SaveAccidents(t *testing.T) {
	mockStore := NewMockStore(nil, nil, nil)
	accidents := []*models.AircraftAccident{
		{ID: 2},
	}
	err := mockStore.SaveAccidents(accidents)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(mockStore.Accidents) != len(accidents) {
		t.Fatalf("Expected %d accidents, got %d", len(accidents), len(mockStore.Accidents))
	}
	for i, accident := range accidents {
		if accident.ID != accidents[i].ID {
			t.Fatalf("Expected ID %d, got %d", accidents[i].ID, accident.ID)
		}
	}
}

// TestMockStore_GetAircrafts tests the GetAircrafts method of the MockStore.
// It checks that the method returns the expected aircrafts and handles errors correctly.
func TestMockStore_GetAircrafts(t *testing.T) {
	expectedAircrafts := []*models.Aircraft{
		{RegistrationNumber: "ABC123"},
	}
	mockStore := NewMockStore(expectedAircrafts, nil, nil)
	aircrafts, err := mockStore.GetAircrafts()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(aircrafts) != len(expectedAircrafts) {
		t.Fatalf("Expected %d aircrafts, got %d", len(expectedAircrafts), len(aircrafts))
	}
	for i, aircraft := range aircrafts {
		if aircraft.RegistrationNumber != expectedAircrafts[i].RegistrationNumber {
			t.Errorf("Expected Registration Number %s, got %s", expectedAircrafts[i].RegistrationNumber, aircraft.RegistrationNumber)
		}
	}
}

// TestMockStore_SaveAircrafts tests the SaveAircrafts method of the MockStore.
// It ensures that the method saves aircrafts correctly and handles errors properly.
func TestMockStore_SaveAircrafts(t *testing.T) {
	mockStore := NewMockStore(nil, nil, nil)
	aircrafts := []*models.Aircraft{
		{RegistrationNumber: "XYZ789"},
	}
	err := mockStore.SaveAircrafts(aircrafts)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(mockStore.Aircrafts) != len(aircrafts) {
		t.Fatalf("Expected %d aircrafts, got %d", len(aircrafts), len(mockStore.Aircrafts))
	}
	for i, aircraft := range aircrafts {
		if aircraft.RegistrationNumber != aircrafts[i].RegistrationNumber {
			t.Fatalf("Expected registration number %s, got %s", aircrafts[i].RegistrationNumber, aircraft.RegistrationNumber)
		}
	}
}

// TestMockStore_GetAircrafts_Error tests error handling in the GetAircrafts method of the MockStore.
func TestMockStore_GetAircrafts_Error(t *testing.T) {
	expectedError := errors.New("query error")
	mockStore := NewMockStore(nil, nil, expectedError)
	_, err := mockStore.GetAircrafts()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

// TestMockStore_GetAccidents_Error tests error handling in the GetAccidents method of the MockStore.
func TestMockStore_GetAccidents_Error(t *testing.T) {
	expectedError := errors.New("query error")
	mockStore := NewMockStore(nil, nil, expectedError)
	_, err := mockStore.GetAccidents()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

// TestMockStore_SaveAircrafts_Error tests error handling in the SaveAircrafts method of the MockStore.
func TestMockStore_SaveAircrafts_Error(t *testing.T) {
	simulatedError := errors.New("simulated save error")
	mockStore := NewMockStore(nil, nil, simulatedError)
	err := mockStore.SaveAircrafts([]*models.Aircraft{{RegistrationNumber: "ABCD"}})
	if err != simulatedError {
		t.Errorf("Expected error '%v', got '%v'", simulatedError, err)
	}
}

// TestMockStore_SaveAccidents_Error tests error handling in the SaveAccidents method of the MockStore.
func TestMockStore_SaveAccidents_Error(t *testing.T) {
	simulatedError := errors.New("simulated save error")
	mockStore := NewMockStore(nil, nil, simulatedError)
	err := mockStore.SaveAccidents([]*models.AircraftAccident{{ID: 3}})
	if err != simulatedError {
		t.Errorf("Expected error '%v', got '%v'", simulatedError, err)
	}
}

// TestMockStore_GetAircraftWithAccidents tests the GetAircraftWithAccidents method of the MockStore.
// It verifies that the method returns the expected aircraft with associated accidents.
func TestMockStore_GetAircraftWithAccidents(t *testing.T) {
	expectedAircraft := &models.Aircraft{
		RegistrationNumber: "ABC123",
		ID:                 1,
	}
	expectedAccidents := []*models.AircraftAccident{
		{ID: 1, AircraftID: 1},
		{ID: 2, AircraftID: 1},
	}

	mockStore := NewMockStore([]*models.Aircraft{expectedAircraft}, expectedAccidents, nil)

	aircraftWithAccidents, err := mockStore.GetAircraftWithAccidents("ABC123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if aircraftWithAccidents == nil {
		t.Fatal("Expected non-nil aircraft, got nil")
	}

	if len(aircraftWithAccidents.Accidents) != len(expectedAccidents) {
		t.Fatalf("Expected %d accidents, got %d", len(expectedAccidents), len(aircraftWithAccidents.Accidents))
	}

	for i, accident := range aircraftWithAccidents.Accidents {
		if accident.ID != expectedAccidents[i].ID {
			t.Errorf("Expected accident ID %d, got %d", expectedAccidents[i].ID, accident.ID)
		}
		if accident.AircraftID != expectedAircraft.ID {
			t.Errorf("Expected accident AircraftID %d, got %d", expectedAircraft.ID, accident.AircraftID)
		}
	}
}

// TestMockStore_GetAircraftWithAccidents_NotFound tests the scenario when the aircraft is not found.
// It ensures that the method returns an error when the aircraft with the specified registration number is not found.
func TestMockStore_GetAircraftWithAccidents_NotFound(t *testing.T) {
	mockStore := NewMockStore(nil, nil, nil)

	// Specify a registration number that doesn't exist in the mock store
	registrationNumber := "NonExistent"

	aircraftWithAccidents, err := mockStore.GetAircraftWithAccidents(registrationNumber)
	if aircraftWithAccidents != nil {
		t.Fatalf("Expected nil aircraft, got %v", aircraftWithAccidents)
	}
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	expectedErrorMessage := fmt.Sprintf("aircraft with registration number %s not found", registrationNumber)
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error '%s', got '%v'", expectedErrorMessage, err)
	}
}
