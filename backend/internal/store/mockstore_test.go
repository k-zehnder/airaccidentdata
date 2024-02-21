package store

import (
	"errors"
	"testing"

	"github.com/computers33333/airaccidentdata/internal/models"
)

// TestMockStore_GetAccidents tests the GetAccidents method of the MockStore.
// It checks whether the method correctly retrieves a set of accidents and handles pagination.
func TestMockStore_GetAccidents(t *testing.T) {
	// Setting up expected accidents for the test.
	expectedAccidents := []*models.AircraftAccident{
		{RegistrationNumber: "1234"},
	}
	// Creating a MockStore with the expected accidents.
	mockStore := NewMockStore(expectedAccidents, nil)

	// Retrieving accidents from the MockStore.
	accidents, total, err := mockStore.GetAccidents(1, 1)

	// Verifying that no error occurred.
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// Checking if the total count of accidents is as expected.
	if total != len(expectedAccidents) {
		t.Fatalf("Expected total count %d, got %d", len(expectedAccidents), total)
	}
	// Ensuring the correct number of accidents are returned for the requested page.
	if len(accidents) != 1 {
		t.Fatalf("Expected 1 accident, got %d", len(accidents))
	}
	// Comparing the actual accidents with the expected ones.
	for i, incident := range accidents {
		if incident.RegistrationNumber != expectedAccidents[i].RegistrationNumber {
			t.Errorf("Expected Registration Number %s, got %s", expectedAccidents[i].RegistrationNumber, incident.RegistrationNumber)
		}
	}
}

// TestMockStore_SaveAccidents tests the SaveAccidents method of the MockStore.
// It checks whether the method correctly stores a set of accidents.
func TestMockStore_SaveAccidents(t *testing.T) {
	// Initializing a MockStore without any pre-existing accidents.
	mockStore := NewMockStore(nil, nil)

	// Defining a set of accidents to be stored in the mock store.
	accidents := []*models.AircraftAccident{
		{RegistrationNumber: "1234A"},
		{RegistrationNumber: "1234B"},
	}

	// Attempting to store the defined accidents in the mock store.
	err := mockStore.SaveAccidents(accidents)

	// Ensuring no error occurred during the save operation.
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// Verifying the mock store contains the correct number of accidents.
	if len(mockStore.Accidents) != len(accidents) {
		t.Fatalf("Expected %d accidents, got %d", len(accidents), len(mockStore.Accidents))
	}
	// Checking if the stored accidents match the defined set.
	for i, accident := range accidents {
		if accident.RegistrationNumber != accidents[i].RegistrationNumber {
			t.Fatalf("Expected registration number %s, got %s", accidents[i].RegistrationNumber, accident.RegistrationNumber)
		}
	}
}

// TestMockStore_GetAccidents_Error tests the error handling in the GetAccidents method of the MockStore.
// It ensures the method correctly returns an error when a simulated error is present.
func TestMockStore_GetAccidents_Error(t *testing.T) {
	// Creating a simulated error to test the error handling of the MockStore.
	expectedError := errors.New("query error")

	// Initializing a MockStore with the simulated error.
	mockStore := NewMockStore(nil, expectedError)

	// Attempting to retrieve accidents from the MockStore, expecting an error due to the simulated condition.
	_, _, err := mockStore.GetAccidents(1, 1)

	// Verifying an error was returned and matches the expected error.
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

// TestMockStore_GetAccidents_EmptyPage tests the GetAccidents method when requesting a page number
// that is beyond the total number of accidents, expecting an empty result.
func TestMockStore_GetAccidents_EmptyPage(t *testing.T) {
	mockStore := NewMockStore([]*models.AircraftAccident{{RegistrationNumber: "1234"}}, nil)

	// Request a page number that is beyond the number of accidents.
	accidents, total, err := mockStore.GetAccidents(2, 10)

	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	if total != 1 {
		t.Fatalf("Expected total to reflect the actual number of accidents, got: %d", total)
	}
	if len(accidents) != 0 {
		t.Fatalf("Expected no accidents, got: %d", len(accidents))
	}
}

// TestMockStore_GetAccidents_EndExceeds tests the GetAccidents method when the 'limit'
// parameter causes the 'end' index to exceed the length of the accidents slice.
func TestMockStore_GetAccidents_EndExceeds(t *testing.T) {
	mockStore := NewMockStore([]*models.AircraftAccident{{RegistrationNumber: "1234"}}, nil)

	// Set a limit that exceeds the number of accidents.
	accidents, total, err := mockStore.GetAccidents(1, 2)

	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	if total != 1 {
		t.Fatalf("Expected total to be the actual number of accidents, got: %d", total)
	}
	if len(accidents) != 1 {
		t.Fatalf("Expected one accident, got: %d", len(accidents))
	}
	if accidents[0].RegistrationNumber != "1234" {
		t.Fatalf("Expected registration number '1234', got: '%s'", accidents[0].RegistrationNumber)
	}
}

// TestMockStore_SaveAccidents_Error tests that SaveAccidents returns an error when the MockStore is initialized with a QueryError.
func TestMockStore_SaveAccidents_Error(t *testing.T) {
	// Simulate a query error that would occur during the save operation.
	simulatedError := errors.New("simulated save error")

	// Initialize a MockStore with a simulated error.
	mockStore := NewMockStore(nil, simulatedError)

	// Attempt to save accidents in the mock store, which should result in the simulated error.
	err := mockStore.SaveAccidents([]*models.AircraftAccident{{RegistrationNumber: "ABCD"}})

	// Verify that the error returned by SaveAccidents matches the simulated error.
	if err != simulatedError {
		t.Errorf("Expected error '%v', got '%v'", simulatedError, err)
	}
}
